// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package enhanced

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadKmeshTracePointCompat returns the embedded CollectionSpec for KmeshTracePointCompat.
func LoadKmeshTracePointCompat() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshTracePointCompatBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshTracePointCompat: %w", err)
	}

	return spec, err
}

// LoadKmeshTracePointCompatObjects loads KmeshTracePointCompat and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshTracePointCompatObjects
//	*KmeshTracePointCompatPrograms
//	*KmeshTracePointCompatMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshTracePointCompatObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshTracePointCompat()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshTracePointCompatSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshTracePointCompatSpecs struct {
	KmeshTracePointCompatProgramSpecs
	KmeshTracePointCompatMapSpecs
}

// KmeshTracePointCompatSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshTracePointCompatProgramSpecs struct {
	ConnectRet *ebpf.ProgramSpec `ebpf:"connect_ret"`
}

// KmeshTracePointCompatMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshTracePointCompatMapSpecs struct {
}

// KmeshTracePointCompatObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshTracePointCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshTracePointCompatObjects struct {
	KmeshTracePointCompatPrograms
	KmeshTracePointCompatMaps
}

func (o *KmeshTracePointCompatObjects) Close() error {
	return _KmeshTracePointCompatClose(
		&o.KmeshTracePointCompatPrograms,
		&o.KmeshTracePointCompatMaps,
	)
}

// KmeshTracePointCompatMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshTracePointCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshTracePointCompatMaps struct {
}

func (m *KmeshTracePointCompatMaps) Close() error {
	return _KmeshTracePointCompatClose()
}

// KmeshTracePointCompatPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshTracePointCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshTracePointCompatPrograms struct {
	ConnectRet *ebpf.Program `ebpf:"connect_ret"`
}

func (p *KmeshTracePointCompatPrograms) Close() error {
	return _KmeshTracePointCompatClose(
		p.ConnectRet,
	)
}

func _KmeshTracePointCompatClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshtracepointcompat_bpfel.o
var _KmeshTracePointCompatBytes []byte
