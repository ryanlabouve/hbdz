package main

import (
	"fmt"
	bpf "github.com/iovisor/gobpf/bcc"
	"os"
)

const source = `
#include <uapi/linux/ptrace.h>
#include <net/sock.h>
#include <linux/net.h>
#include <bcc/proto.h>

BPF_PERF_OUTPUT(events);

int on_accept(struct pt_regs *ctx, struct sock *sk) {
    u16 sport = 0, dport = 0;
    u32 saddr = 0, daddr = 0;

    bpf_probe_read(&sport, sizeof(sport), &sk->__sk_common.skc_num);
    bpf_probe_read(&dport, sizeof(dport), &sk->__sk_common.skc_dport);
    bpf_probe_read(&saddr, sizeof(saddr), &sk->__sk_common.skc_rcv_saddr);
    bpf_probe_read(&daddr, sizeof(daddr), &sk->__sk_common.skc_daddr);

    if (dport == 80) { // Filter for port 80
        char msg[] = "wohoo! SYN from ";
        bpf_trace_printk(msg, sizeof(msg));
        bpf_trace_printk("%u.%u.%u.%u\n", saddr & 0xff, (saddr >> 8) & 0xff,
                         (saddr >> 16) & 0xff, (saddr >> 24) & 0xff);
    }

    return 0;
}
`

func main() {
	m := bpf.NewModule(source, []string{})
	defer m.Close()

	fn, err := m.Load("on_accept", bpf.Kprobe("inet_csk_accept"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load program: %v\n", err)
		os.Exit(1)
	}

	if err := m.AttachKprobe("inet_csk_accept", fn); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to attach program: %v\n", err)
		os.Exit(1)
	}

	channel := make(chan []byte)

	m.PerfMap("events", channel)

	fmt.Println("Listening for SYN packets...")

	for {
		data := <-channel
		fmt.Print(string(data))
	}
}
