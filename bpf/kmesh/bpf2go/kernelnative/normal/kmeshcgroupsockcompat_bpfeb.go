// Code generated by bpf2go; DO NOT EDIT.
//go:build mips || mips64 || ppc64 || s390x

package normal

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type KmeshCgroupSockCompatBuf struct{ Data [40]int8 }

type KmeshCgroupSockCompatClusterSockData struct{ ClusterId uint32 }

type KmeshCgroupSockCompatLogEvent struct {
	Ret uint32
	Msg [255]int8
	_   [1]byte
}

type KmeshCgroupSockCompatManagerKey struct {
	NetnsCookie uint64
	_           [8]byte
}

type KmeshCgroupSockCompatRatelimitKey struct {
	Key struct {
		SkSkb struct {
			Netns  uint64
			Ipv4   uint32
			Port   uint32
			Family uint32
			_      [4]byte
		}
	}
}

type KmeshCgroupSockCompatRatelimitValue struct {
	LastTopup uint64
	Tokens    uint64
}

type KmeshCgroupSockCompatSockStorageData struct {
	ConnectNs      uint64
	Direction      uint8
	ConnectSuccess uint8
	_              [6]byte
}

// LoadKmeshCgroupSockCompat returns the embedded CollectionSpec for KmeshCgroupSockCompat.
func LoadKmeshCgroupSockCompat() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshCgroupSockCompatBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshCgroupSockCompat: %w", err)
	}

	return spec, err
}

// LoadKmeshCgroupSockCompatObjects loads KmeshCgroupSockCompat and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshCgroupSockCompatObjects
//	*KmeshCgroupSockCompatPrograms
//	*KmeshCgroupSockCompatMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshCgroupSockCompatObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshCgroupSockCompat()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshCgroupSockCompatSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockCompatSpecs struct {
	KmeshCgroupSockCompatProgramSpecs
	KmeshCgroupSockCompatMapSpecs
}

// KmeshCgroupSockCompatSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockCompatProgramSpecs struct {
	CgroupConnect4Prog *ebpf.ProgramSpec `ebpf:"cgroup_connect4_prog"`
	ClusterManager     *ebpf.ProgramSpec `ebpf:"cluster_manager"`
	FilterChainManager *ebpf.ProgramSpec `ebpf:"filter_chain_manager"`
	FilterManager      *ebpf.ProgramSpec `ebpf:"filter_manager"`
}

// KmeshCgroupSockCompatMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockCompatMapSpecs struct {
	InnerMap            *ebpf.MapSpec `ebpf:"inner_map"`
	KmeshCluster        *ebpf.MapSpec `ebpf:"kmesh_cluster"`
	KmeshClusterStats   *ebpf.MapSpec `ebpf:"kmesh_cluster_stats"`
	KmeshConfigMap      *ebpf.MapSpec `ebpf:"kmesh_config_map"`
	KmeshEvents         *ebpf.MapSpec `ebpf:"kmesh_events"`
	KmeshListener       *ebpf.MapSpec `ebpf:"kmesh_listener"`
	KmeshManage         *ebpf.MapSpec `ebpf:"kmesh_manage"`
	KmeshRatelimit      *ebpf.MapSpec `ebpf:"kmesh_ratelimit"`
	KmeshTailCallCtx    *ebpf.MapSpec `ebpf:"kmesh_tail_call_ctx"`
	KmeshTailCallProg   *ebpf.MapSpec `ebpf:"kmesh_tail_call_prog"`
	MapOfClusterEps     *ebpf.MapSpec `ebpf:"map_of_cluster_eps"`
	MapOfClusterEpsData *ebpf.MapSpec `ebpf:"map_of_cluster_eps_data"`
	MapOfClusterSock    *ebpf.MapSpec `ebpf:"map_of_cluster_sock"`
	MapOfSockStorage    *ebpf.MapSpec `ebpf:"map_of_sock_storage"`
	OuterMap            *ebpf.MapSpec `ebpf:"outer_map"`
	OuterOfMaglev       *ebpf.MapSpec `ebpf:"outer_of_maglev"`
	TmpBuf              *ebpf.MapSpec `ebpf:"tmp_buf"`
	TmpLogBuf           *ebpf.MapSpec `ebpf:"tmp_log_buf"`
}

// KmeshCgroupSockCompatObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockCompatObjects struct {
	KmeshCgroupSockCompatPrograms
	KmeshCgroupSockCompatMaps
}

func (o *KmeshCgroupSockCompatObjects) Close() error {
	return _KmeshCgroupSockCompatClose(
		&o.KmeshCgroupSockCompatPrograms,
		&o.KmeshCgroupSockCompatMaps,
	)
}

// KmeshCgroupSockCompatMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockCompatMaps struct {
	InnerMap            *ebpf.Map `ebpf:"inner_map"`
	KmeshCluster        *ebpf.Map `ebpf:"kmesh_cluster"`
	KmeshClusterStats   *ebpf.Map `ebpf:"kmesh_cluster_stats"`
	KmeshConfigMap      *ebpf.Map `ebpf:"kmesh_config_map"`
	KmeshEvents         *ebpf.Map `ebpf:"kmesh_events"`
	KmeshListener       *ebpf.Map `ebpf:"kmesh_listener"`
	KmeshManage         *ebpf.Map `ebpf:"kmesh_manage"`
	KmeshRatelimit      *ebpf.Map `ebpf:"kmesh_ratelimit"`
	KmeshTailCallCtx    *ebpf.Map `ebpf:"kmesh_tail_call_ctx"`
	KmeshTailCallProg   *ebpf.Map `ebpf:"kmesh_tail_call_prog"`
	MapOfClusterEps     *ebpf.Map `ebpf:"map_of_cluster_eps"`
	MapOfClusterEpsData *ebpf.Map `ebpf:"map_of_cluster_eps_data"`
	MapOfClusterSock    *ebpf.Map `ebpf:"map_of_cluster_sock"`
	MapOfSockStorage    *ebpf.Map `ebpf:"map_of_sock_storage"`
	OuterMap            *ebpf.Map `ebpf:"outer_map"`
	OuterOfMaglev       *ebpf.Map `ebpf:"outer_of_maglev"`
	TmpBuf              *ebpf.Map `ebpf:"tmp_buf"`
	TmpLogBuf           *ebpf.Map `ebpf:"tmp_log_buf"`
}

func (m *KmeshCgroupSockCompatMaps) Close() error {
	return _KmeshCgroupSockCompatClose(
		m.InnerMap,
		m.KmeshCluster,
		m.KmeshClusterStats,
		m.KmeshConfigMap,
		m.KmeshEvents,
		m.KmeshListener,
		m.KmeshManage,
		m.KmeshRatelimit,
		m.KmeshTailCallCtx,
		m.KmeshTailCallProg,
		m.MapOfClusterEps,
		m.MapOfClusterEpsData,
		m.MapOfClusterSock,
		m.MapOfSockStorage,
		m.OuterMap,
		m.OuterOfMaglev,
		m.TmpBuf,
		m.TmpLogBuf,
	)
}

// KmeshCgroupSockCompatPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockCompatObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockCompatPrograms struct {
	CgroupConnect4Prog *ebpf.Program `ebpf:"cgroup_connect4_prog"`
	ClusterManager     *ebpf.Program `ebpf:"cluster_manager"`
	FilterChainManager *ebpf.Program `ebpf:"filter_chain_manager"`
	FilterManager      *ebpf.Program `ebpf:"filter_manager"`
}

func (p *KmeshCgroupSockCompatPrograms) Close() error {
	return _KmeshCgroupSockCompatClose(
		p.CgroupConnect4Prog,
		p.ClusterManager,
		p.FilterChainManager,
		p.FilterManager,
	)
}

func _KmeshCgroupSockCompatClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshcgroupsockcompat_bpfeb.o
var _KmeshCgroupSockCompatBytes []byte
