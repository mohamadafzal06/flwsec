#include <stdlib.h>
#include <stdint.h>

#include <linux/bpf.h>
#include <linux/bpf_common.h>
#include <bpf/bpf_helpers.h>

SEC("classifier/probe")
int probe(struct __sk_buff *skb){
    bpf_printk("Hello, world. I have a packet\n");
    return 0;
}

char _license[] SEC("license") = "Dual MIT/GPL";
