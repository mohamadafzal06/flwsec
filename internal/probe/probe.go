package probe

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang  probe ../../bpf/probe.c -- -O3 -Wall -Werror -Wno-address-of-packed-member

func load() (probeObjects, error) {
	var objs probeObjects
	err := loadProbeObjects(&objs, nil)
	if err != nil {
		return probeObjects{}, nil
	}
	return objs, nil

}

// func must(e error) error {
// 	if e != nil {
// 		return e
// 	}
// 	return nil
// }
