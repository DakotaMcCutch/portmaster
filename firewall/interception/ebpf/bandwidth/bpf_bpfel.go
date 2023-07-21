// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package ebpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfSkInfo struct {
	Rx       uint64
	Tx       uint64
	Reported uint64
}

type bpfSkKey struct {
	SrcIp    [4]uint32
	DstIp    [4]uint32
	SrcPort  uint16
	DstPort  uint16
	Protocol uint8
	Ipv6     uint8
	_        [2]byte
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	SocketOperations *ebpf.ProgramSpec `ebpf:"socket_operations"`
	UdpRecvmsg       *ebpf.ProgramSpec `ebpf:"udp_recvmsg"`
	UdpSendmsg       *ebpf.ProgramSpec `ebpf:"udp_sendmsg"`
	Udpv6Recvmsg     *ebpf.ProgramSpec `ebpf:"udpv6_recvmsg"`
	Udpv6Sendmsg     *ebpf.ProgramSpec `ebpf:"udpv6_sendmsg"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	PmBandwidthMap *ebpf.MapSpec `ebpf:"pm_bandwidth_map"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	PmBandwidthMap *ebpf.Map `ebpf:"pm_bandwidth_map"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.PmBandwidthMap,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	SocketOperations *ebpf.Program `ebpf:"socket_operations"`
	UdpRecvmsg       *ebpf.Program `ebpf:"udp_recvmsg"`
	UdpSendmsg       *ebpf.Program `ebpf:"udp_sendmsg"`
	Udpv6Recvmsg     *ebpf.Program `ebpf:"udpv6_recvmsg"`
	Udpv6Sendmsg     *ebpf.Program `ebpf:"udpv6_sendmsg"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.SocketOperations,
		p.UdpRecvmsg,
		p.UdpSendmsg,
		p.Udpv6Recvmsg,
		p.Udpv6Sendmsg,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel.o
var _BpfBytes []byte
