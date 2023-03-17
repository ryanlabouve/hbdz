package main

import (
	"fmt"
	"github.com/iovisor/gobpf/bcc"
	"os"
	"os/signal"
)

const source string = `
#include <uapi/linux/ptrace.h>
#include <linux/sched.h>

struct data_t {
    u32 pid;
};

BPF_PERF_OUTPUT(events);

int trace_clone(struct pt_regs *ctx) {
	struct data_t data = {};

	struct task_struct *task;
	task = (struct task_struct *)bpf_get_current_task();

	data.pid = task->pid;

	events.perf_submit(ctx, &data, sizeof(data));

	return 0;
}
`

func main() {
	m := bcc.NewModule(source, []string{})
	defer m.Close()

	kprobe, err := m.LoadKprobe("trace_clone")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load trace_clone: %s\n", err)
		os.Exit(1)
	}

	err = m.AttachKprobe("__x64_sys_clone", kprobe, -1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to attach clone: %s\n", err)
		os.Exit(1)
	}

	table := bcc.NewTable(m.TableId("events"), m)

	channel := make(chan []byte)
	lostChannel := make(chan uint64)

	perfMap, err := bcc.InitPerfMap(table, channel, lostChannel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init perf map: %s\n", err)
		os.Exit(1)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	fmt.Println("Listening for new processes...")

	go func() {
		var data struct {
			Pid uint32
		}
		for {
			select {
			case <-sig:
				return
			case payload := <-channel:
				data.Pid = bcc.GetHostByteOrder().Uint32(payload)
				fmt.Printf("New process created: PID %d\n", data.Pid)

			}
		}
	}()

	perfMap.Start()
	<-sig
	perfMap.Stop()
}
