package api

import (
	// kclusterv1alpha1 "github.com/karmada-io/api/cluster/v1alpha1"
	// kpolicyv1alpha1 "github.com/karmada-io/api/policy/v1alpha1"
	// kworkv1alpha1 "github.com/karmada-io/api/work/v1alpha1"

	kclusterv1alpha1 "github.com/karmada-io/karmada/pkg/apis/cluster/v1alpha1"
	kpolicyv1alpha1 "github.com/karmada-io/karmada/pkg/apis/policy/v1alpha1"

	kworkv1alpha1 "github.com/karmada-io/karmada/pkg/apis/work/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
)

type ClusterTaskID types.UID
type PlacementID types.UID
type ClusterDetailInfo struct {
	Name              string
	Score             int64
	AvailableReplicas int64
	Cluster           *kclusterv1alpha1.Cluster
}

// ProviderInfo indicate the provider information
type ProviderInfo struct {
	Name              string
	Score             int64 // the highest score in all clusters of the provider
	AvailableReplicas int64

	// Regions under this provider
	Regions map[string]struct{}
	// Zones under this provider
	Zones map[string]struct{}
	// Clusters under this provider, sorted by cluster.Score descending.
	Clusters []ClusterDetailInfo
}

// RegionInfo indicate the region information
type RegionInfo struct {
	Name              string
	Score             int64 // the highest score in all clusters of the region
	AvailableReplicas int64

	// Zones under this provider
	Zones map[string]struct{}
	// Clusters under this region, sorted by cluster.Score descending.
	Clusters []ClusterDetailInfo
}

// ZoneInfo indicate the zone information
type ZoneInfo struct {
	Name              string
	Score             int64 // the highest score in all clusters of the zone
	AvailableReplicas int64

	// Clusters under this zone, sorted by cluster.Score descending.
	Clusters []ClusterDetailInfo
}

// ClustersTopology indicate the cluster global view
type ClustersTopology struct {
	Providers map[string]ProviderInfo
	Regions   map[string]RegionInfo
	Zones     map[string]ZoneInfo

	// Clusters from global view, sorted by cluster.Score descending.
	Clusters []ClusterDetailInfo
}

func NewCluster(cluster *kclusterv1alpha1.Cluster) *ClusterDetailInfo {
	return &ClusterDetailInfo{
		Name:    cluster.Name,
		Cluster: cluster,
	}
}

type ClusterTaskStatus string

const (
	ClusterTaskUnscheduled ClusterTaskStatus = "Unscheduled"
	ClusterTaskScheduled   ClusterTaskStatus = "Scheduled"
	ClusterTaskOutofDate   ClusterTaskStatus = "OutOfDate"
)

type ClusterTaskType string

const (
	ClusterTask   ClusterTaskType = "Cluster"
	NamespaceTask ClusterTaskType = "Namespace"
)

type IResourceBinding interface {
	TypeMeta() *metav1.TypeMeta
	ObjectMeta() *metav1.ObjectMeta
	Spec() *kworkv1alpha1.ResourceBindingSpec
	Status() *kworkv1alpha1.ResourceBindingStatus
	DeepCopy() IResourceBinding
}

type ResourceBinding struct {
	*kworkv1alpha1.ResourceBinding
}

func (rb *ResourceBinding) TypeMeta() *metav1.TypeMeta {
	return &rb.ResourceBinding.TypeMeta
}

func (rb *ResourceBinding) ObjectMeta() *metav1.ObjectMeta {
	return &rb.ResourceBinding.ObjectMeta
}

func (rb *ResourceBinding) Spec() *kworkv1alpha1.ResourceBindingSpec {
	return &rb.ResourceBinding.Spec
}

func (rb *ResourceBinding) Status() *kworkv1alpha1.ResourceBindingStatus {
	return &rb.ResourceBinding.Status
}

func (rb *ResourceBinding) DeepCopy() IResourceBinding {
	return &ResourceBinding{
		ResourceBinding: rb.ResourceBinding.DeepCopy(),
	}
}

type ClusterResourceBinding struct {
	*kworkv1alpha1.ClusterResourceBinding
}

func (rb *ClusterResourceBinding) TypeMeta() *metav1.TypeMeta {
	return &rb.ClusterResourceBinding.TypeMeta
}

func (rb *ClusterResourceBinding) ObjectMeta() *metav1.ObjectMeta {
	return &rb.ClusterResourceBinding.ObjectMeta
}

func (rb *ClusterResourceBinding) Spec() *kworkv1alpha1.ResourceBindingSpec {
	return &rb.ClusterResourceBinding.Spec
}

func (rb *ClusterResourceBinding) Status() *kworkv1alpha1.ResourceBindingStatus {
	return &rb.ClusterResourceBinding.Status
}

func (rb *ClusterResourceBinding) DeepCopy() IResourceBinding {
	return &ClusterResourceBinding{
		ClusterResourceBinding: rb.ClusterResourceBinding.DeepCopy(),
	}
}

type ClusterTaskInfo struct {
	UID                    ClusterTaskID
	Name                   string
	State                  ClusterTaskStatus
	Type                   ClusterTaskType
	CreationTimestamp      metav1.Time
	ScheduleStartTimestamp metav1.Time
	Replicas               int

	ResourceBinding IResourceBinding
}

func NewClusterTaskInfo(typ ClusterTaskType, rb IResourceBinding) *ClusterTaskInfo {
	return &ClusterTaskInfo{
		Name:            rb.ObjectMeta().Name,
		State:           ClusterTaskUnscheduled,
		Type:            NamespaceTask,
		ResourceBinding: rb,
	}
}

func (t *ClusterTaskInfo) OnDeleteCluster(clusterName string) {
	if t.State != ClusterTaskScheduled {
		return
	}
	target := t.ResourceBinding.Spec().Clusters
	// TODO: delete cluster from resource binding target or not?
	for _, tc := range target {
		if tc.Name == clusterName {
			t.State = ClusterTaskOutofDate
			return
		}
	}
}

func (t *ClusterTaskInfo) OnPlacementChanged(selector labels.Selector) {
	if t.State != ClusterTaskScheduled {
		return
	}
	rbLabels := t.ResourceBinding.ObjectMeta().Labels
	if selector.Matches(labels.Set(rbLabels)) {
		t.State = ClusterTaskOutofDate
	}
}

type PlacementType string

const (
	ClusterPlacement   PlacementType = "Cluster"
	NamespacePlacement PlacementType = "Namespace"
)

type IPropagationPolicy interface {
	TypeMeta() *metav1.TypeMeta
	ObjectMeta() *metav1.ObjectMeta
	Spec() *kpolicyv1alpha1.PropagationSpec
	DeepCopy() IPropagationPolicy
}

type PropagationPolicy struct {
	*kpolicyv1alpha1.PropagationPolicy
}

func (pp *PropagationPolicy) TypeMeta() *metav1.TypeMeta {
	return &pp.PropagationPolicy.TypeMeta
}

func (pp *PropagationPolicy) ObjectMeta() *metav1.ObjectMeta {
	return &pp.PropagationPolicy.ObjectMeta
}

func (pp *PropagationPolicy) Spec() *kpolicyv1alpha1.PropagationSpec {
	return &pp.PropagationPolicy.Spec
}

func (pp *PropagationPolicy) DeepCopy() IPropagationPolicy {
	return &PropagationPolicy{
		PropagationPolicy: pp.PropagationPolicy.DeepCopy(),
	}
}

type ClusterPropagationPolicy struct {
	*kpolicyv1alpha1.ClusterPropagationPolicy
}

func (pp *ClusterPropagationPolicy) TypeMeta() *metav1.TypeMeta {
	return &pp.ClusterPropagationPolicy.TypeMeta
}

func (pp *ClusterPropagationPolicy) ObjectMeta() *metav1.ObjectMeta {
	return &pp.ClusterPropagationPolicy.ObjectMeta
}

func (pp *ClusterPropagationPolicy) Spec() *kpolicyv1alpha1.PropagationSpec {
	return &pp.ClusterPropagationPolicy.Spec
}

func (pp *ClusterPropagationPolicy) DeepCopy() IPropagationPolicy {
	return &ClusterPropagationPolicy{
		ClusterPropagationPolicy: pp.ClusterPropagationPolicy.DeepCopy(),
	}
}

type PlacementInfo struct {
	UID    PlacementID
	Name   string
	Type   PlacementType
	Policy IPropagationPolicy
}

func NewPlacementInfo(typ PlacementType, policy IPropagationPolicy) *PlacementInfo {
	return &PlacementInfo{
		Type:   typ,
		Policy: policy,
	}
}
