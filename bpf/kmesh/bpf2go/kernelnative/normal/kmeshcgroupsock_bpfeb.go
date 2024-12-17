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

type KmeshCgroupSockBuf struct{ Data [40]int8 }

type KmeshCgroupSockClusterSockData struct{ ClusterId uint32 }

type KmeshCgroupSockKmeshConfig struct {
	BpfLogLevel      uint32
	NodeIp           [4]uint32
	PodGateway       [4]uint32
	AuthzOffload     uint32
	EnableMonitoring uint32
}

type KmeshCgroupSockManagerKey struct {
	NetnsCookie uint64
	_           [8]byte
}

type KmeshCgroupSockRatelimitKey struct {
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

type KmeshCgroupSockRatelimitValue struct {
	LastTopup uint64
	Tokens    uint64
}

type KmeshCgroupSockSockStorageData struct {
	ConnectNs      uint64
	Direction      uint8
	ConnectSuccess uint8
	_              [6]byte
}

// LoadKmeshCgroupSock returns the embedded CollectionSpec for KmeshCgroupSock.
func LoadKmeshCgroupSock() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshCgroupSockBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshCgroupSock: %w", err)
	}

	return spec, err
}

// LoadKmeshCgroupSockObjects loads KmeshCgroupSock and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshCgroupSockObjects
//	*KmeshCgroupSockPrograms
//	*KmeshCgroupSockMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshCgroupSockObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshCgroupSock()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshCgroupSockSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockSpecs struct {
	KmeshCgroupSockProgramSpecs
	KmeshCgroupSockMapSpecs
}

// KmeshCgroupSockSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockProgramSpecs struct {
	CgroupConnect4Prog *ebpf.ProgramSpec `ebpf:"cgroup_connect4_prog"`
	ClusterManager     *ebpf.ProgramSpec `ebpf:"cluster_manager"`
	FilterChainManager *ebpf.ProgramSpec `ebpf:"filter_chain_manager"`
	FilterManager      *ebpf.ProgramSpec `ebpf:"filter_manager"`
}

// KmeshCgroupSockMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockMapSpecs struct {
	KmCgrptailcall *ebpf.MapSpec `ebpf:"km_cgrptailcall"`
	KmCluster      *ebpf.MapSpec `ebpf:"km_cluster"`
	KmClusterEps   *ebpf.MapSpec `ebpf:"km_cluster_eps"`
	KmClusterSock  *ebpf.MapSpec `ebpf:"km_cluster_sock"`
	KmClusterstats *ebpf.MapSpec `ebpf:"km_clusterstats"`
	KmConfigmap    *ebpf.MapSpec `ebpf:"km_configmap"`
	KmEpsData      *ebpf.MapSpec `ebpf:"km_eps_data"`
	KmListener     *ebpf.MapSpec `ebpf:"km_listener"`
	KmLogEvent     *ebpf.MapSpec `ebpf:"km_log_event"`
	KmMaglevOuter  *ebpf.MapSpec `ebpf:"km_maglev_outer"`
	KmManage       *ebpf.MapSpec `ebpf:"km_manage"`
	KmRatelimit    *ebpf.MapSpec `ebpf:"km_ratelimit"`
	KmSockstorage  *ebpf.MapSpec `ebpf:"km_sockstorage"`
	KmTailcallCtx  *ebpf.MapSpec `ebpf:"km_tailcall_ctx"`
	KmTmpbuf       *ebpf.MapSpec `ebpf:"km_tmpbuf"`
	KmeshMap1600   *ebpf.MapSpec `ebpf:"kmesh_map1600"`
	KmeshMap192    *ebpf.MapSpec `ebpf:"kmesh_map192"`
	KmeshMap296    *ebpf.MapSpec `ebpf:"kmesh_map296"`
	KmeshMap64     *ebpf.MapSpec `ebpf:"kmesh_map64"`
}

// KmeshCgroupSockObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockObjects struct {
	KmeshCgroupSockPrograms
	KmeshCgroupSockMaps
}

func (o *KmeshCgroupSockObjects) Close() error {
	return _KmeshCgroupSockClose(
		&o.KmeshCgroupSockPrograms,
		&o.KmeshCgroupSockMaps,
	)
}

// KmeshCgroupSockMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockMaps struct {
	KmCgrptailcall *ebpf.Map `ebpf:"km_cgrptailcall"`
	KmCluster      *ebpf.Map `ebpf:"km_cluster"`
	KmClusterEps   *ebpf.Map `ebpf:"km_cluster_eps"`
	KmClusterSock  *ebpf.Map `ebpf:"km_cluster_sock"`
	KmClusterstats *ebpf.Map `ebpf:"km_clusterstats"`
	KmConfigmap    *ebpf.Map `ebpf:"km_configmap"`
	KmEpsData      *ebpf.Map `ebpf:"km_eps_data"`
	KmListener     *ebpf.Map `ebpf:"km_listener"`
	KmLogEvent     *ebpf.Map `ebpf:"km_log_event"`
	KmMaglevOuter  *ebpf.Map `ebpf:"km_maglev_outer"`
	KmManage       *ebpf.Map `ebpf:"km_manage"`
	KmRatelimit    *ebpf.Map `ebpf:"km_ratelimit"`
	KmSockstorage  *ebpf.Map `ebpf:"km_sockstorage"`
	KmTailcallCtx  *ebpf.Map `ebpf:"km_tailcall_ctx"`
	KmTmpbuf       *ebpf.Map `ebpf:"km_tmpbuf"`
	KmeshMap1600   *ebpf.Map `ebpf:"kmesh_map1600"`
	KmeshMap192    *ebpf.Map `ebpf:"kmesh_map192"`
	KmeshMap296    *ebpf.Map `ebpf:"kmesh_map296"`
	KmeshMap64     *ebpf.Map `ebpf:"kmesh_map64"`
}

func (m *KmeshCgroupSockMaps) Close() error {
	return _KmeshCgroupSockClose(
		m.KmCgrptailcall,
		m.KmCluster,
		m.KmClusterEps,
		m.KmClusterSock,
		m.KmClusterstats,
		m.KmConfigmap,
		m.KmEpsData,
		m.KmListener,
		m.KmLogEvent,
		m.KmMaglevOuter,
		m.KmManage,
		m.KmRatelimit,
		m.KmSockstorage,
		m.KmTailcallCtx,
		m.KmTmpbuf,
		m.KmeshMap1600,
		m.KmeshMap192,
		m.KmeshMap296,
		m.KmeshMap64,
	)
}

// KmeshCgroupSockPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockPrograms struct {
	CgroupConnect4Prog *ebpf.Program `ebpf:"cgroup_connect4_prog"`
	ClusterManager     *ebpf.Program `ebpf:"cluster_manager"`
	FilterChainManager *ebpf.Program `ebpf:"filter_chain_manager"`
	FilterManager      *ebpf.Program `ebpf:"filter_manager"`
}

func (p *KmeshCgroupSockPrograms) Close() error {
	return _KmeshCgroupSockClose(
		p.CgroupConnect4Prog,
		p.ClusterManager,
		p.FilterChainManager,
		p.FilterManager,
	)
}

func _KmeshCgroupSockClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshcgroupsock_bpfeb.o
var _KmeshCgroupSockBytes []byte
