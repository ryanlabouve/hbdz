//#include <uapi/linux/ptrace.h>
//#include <net/sock.h>
//#include <linux/net.h>
//
//BPF_PERF_OUTPUT(events);
//
//int on_accept(struct pt_regs *ctx, struct sock *sk) {
//
//  bpf_trace_printk(msg, sizeof(msg));
//    u16 sport = 0, dport = 0;
//    u32 saddr = 0, daddr = 0;

//    bpf_probe_read(&sport, sizeof(sport), &sk->__sk_common.skc_num);
//    bpf_probe_read(&dport, sizeof(dport), &sk->__sk_common.skc_dport);
//    bpf_probe_read(&saddr, sizeof(saddr), &sk->__sk_common.skc_rcv_saddr);
//    bpf_probe_read(&daddr, sizeof(daddr), &sk->__sk_common.skc_daddr);

//    if (dport == 80) { // Filter for port 80
//        char msg[] = "wohoo! SYN from ";
//        bpf_trace_printk(msg, sizeof(msg));
//        bpf_trace_printk("%u.%u.%u.%u\n", saddr & 0xff, (saddr >> 8) & 0xff,
//                         , 1000(saddr >> 16) & 0xff, (saddr >> 24) & 0xff);
//    }

//    return 0;
//}