/*
 Copyright 2021 The Volcano Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package clusterallocate

import (
	"github.com/karmada-io/karmada/pkg/apis/cluster/v1alpha1"
	"k8s.io/klog"

	"volcano.sh/volcano/pkg/scheduler/api"
	"volcano.sh/volcano/pkg/scheduler/framework"
)

type Action struct{}

func New() *Action {
	return &Action{}
}

func (alloc *Action) Name() string {
	return "cluster-allocate"
}

func (alloc *Action) Initialize() {
	klog.V(3).Infof("Enter...")
	defer klog.V(3).Infof("Leaving...")
}

func (alloc *Action) Execute(ssn *framework.Session) {
	klog.V(3).Infof("Enter ClusterAllocate...")
	defer klog.V(3).Infof("Leaving ClusterAllocate...")

	klog.V(3).Infof("clusters: %v", ssn.Clusters)
	klog.V(3).Infof("cluster tasks: %v", ssn.ClusterTasks)
	klog.V(3).Infof("placement: %v", ssn.Placements)

	for _, task := range ssn.ClusterTasks {
		// TODO: if task don't need to be scheduled, continue
		candidates := []*api.Cluster{}
		placement := alloc.getPlacement(ssn)
		for _, cluster := range ssn.Clusters {
			if err := ssn.ClusterPredicateFn(task, cluster, placement); err != nil {
				candidates = append(candidates, cluster)
			}
		}
		scores, err := ssn.BatchClusterOrderFn(task, candidates, placement)
		if err != nil {
			klog.Errorf("calcate scores error for task %s: %v", task.Name, err)
		}
		// select clusters by score
		// TODO: sort by socre, but only 0 and 100 for now
		seeds := []*api.Cluster{}
		for _, cluster := range candidates {
			if score, ok := scores[cluster.Name]; !ok {
				continue
			} else {
				if score == 100 {
					seeds = append(seeds, cluster)
				}
			}
		}
		// TODO: assign replicas to each clusters by weight
		replicas := []v1alpha1.TargetCluster{}
		for _, cluster := range seeds {
			replicas = append(replicas, []v1alpha1.TargetCluster{
				Name:     cluster.Name,
				Replicas: float64(task.ResourceBinding.Spec.Resource.Replicas),
			})
		}
		// TODO: patch task with scheduler's decision
		newRb := task.ResourceBinding.DeepCopy()
		newRb.Spec.Clusters =
			ssn.karmadaClient.WorkV1alpha2().ResourceBindings("xx").Patch()
	}

}

func (alloc *Action) UnInitialize() {
	klog.V(3).Infof("Enter...")
	defer klog.V(3).Infof("Leaving...")
}

func (alloc *Action) getPlacement(ssn *framework.Session) *api.PlacementInfo {
	return nil
}
