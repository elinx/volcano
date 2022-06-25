package api

import (
	// kclusterv1alpha1 "github.com/karmada-io/api/cluster/v1alpha1"
	// kpolicyv1alpha1 "github.com/karmada-io/api/policy/v1alpha1"
	// kworkv1alpha1 "github.com/karmada-io/api/work/v1alpha1"

	kclusterv1alpha1 "github.com/karmada-io/karmada/pkg/apis/cluster/v1alpha1"
	kpolicyv1alpha1 "github.com/karmada-io/karmada/pkg/apis/policy/v1alpha1"
	kworkv1alpha1 "github.com/karmada-io/karmada/pkg/apis/work/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
)

type ClusterTaskID types.UID
type PlacementID types.UID
type Cluster struct {
	Name      string
	Namespace string
	Cluster   *kclusterv1alpha1.Cluster
}

func NewCluster(cluster *kclusterv1alpha1.Cluster) *Cluster {
	return &Cluster{
		Name:      cluster.Name,
		Namespace: cluster.Namespace,
		Cluster:   cluster,
	}
}

type ClusterTaskInfo struct {
	Name            string
	ResourceBinding *kworkv1alpha1.ResourceBinding
}

func NewClusterTaskInfo(rb *kworkv1alpha1.ResourceBinding) *ClusterTaskInfo {
	return &ClusterTaskInfo{
		ResourceBinding: rb,
	}
}

type PlacementType string

const (
	ClusterPlacement   PlacementType = "Cluster"
	NamespacePlacement PlacementType = "Namespace"
)

type PlacementInfo struct {
	Name      string
	Type      PlacementType
	Placement *kpolicyv1alpha1.Placement
}

func NewPlacementInfo(p *kpolicyv1alpha1.Placement, t PlacementType) *PlacementInfo {
	return &PlacementInfo{
		Type:      t,
		Placement: p,
	}
}
