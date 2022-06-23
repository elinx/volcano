package api

import (
	kclusterv1alpha1 "github.com/karmada-io/api/cluster/v1alpha1"
	kpolicyv1alpha1 "github.com/karmada-io/api/policy/v1alpha1"
	kworkv1alpha1 "github.com/karmada-io/api/work/v1alpha1"
)

type Cluster struct {
	Name    string
	Cluster kclusterv1alpha1.Cluster
}

type ClusterTaskInfo struct {
	Name            string
	ResourceBinding kworkv1alpha1.ResourceBinding
}

type PlacementInfo struct {
	Name      string
	Placement kpolicyv1alpha1.Placement
}
