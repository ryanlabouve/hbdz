package main

import (
	"fmt"
	"github.com/cilium/ebpf"
	"os"
)

func main() {
	f, err := os.Open("collection-spec.c")
	if err != nil {
		fmt.Println("Error loading C program")
	}
	defer f.Close()

	spec, err := ebpf.LoadCollectionSpecFromReader(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load collection spec: %v\n", err)
		os.Exit(1)
	}

	coll, err := ebpf.NewCollectionWithOptions(spec, ebpf.CollectionOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create collection: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("%v", coll)
	//prog := coll.Programs["on_accept"]
	//if prog == nil {
	//	fmt.Fprintln(os.Stderr, "Failed to find program")
	//	os.Exit(1)
	//}
	//
	//kprobe, err := ebpf.NewKprobe("inet_csk_accept", asm.Asm(prog.FD()))
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Failed to create kprobe: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//if err := kprobe.Attach(); err != nil {
	//	fmt.Fprintf(os.Stderr, "Failed to attach kprobe: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//channel := make(chan []byte)
	//
	//perfMap, err := ebpf.NewPerfMapWithOptions(coll, "events", ebpf.PerfBufferOptions{
	//	OnLost: func(count uint64) {
	//		fmt.Fprintf(os.Stderr, "Lost %d events\n", count)
	//	},
	//})
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Failed to create perf map: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//perfMap.Start(channel)
	//defer perfMap.Stop()
	//
	//fmt.Println("Listening for SYN packets...")
	//
	//for {
	//	data := <-channel
	//	fmt.Print(string(data))
	//}
}
