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
	"context"
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch"
	kworkv1alpha1 "github.com/karmada-io/karmada/pkg/apis/work/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
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
		// TODO: task maybe update
		if len(task.ResourceBinding.Spec.Clusters) != 0 {
			continue
		}
		candidates := []*api.Cluster{}
		placement := alloc.getPlacement(ssn)
		for _, cluster := range ssn.Clusters {
			if err := ssn.ClusterPredicateFn(task, cluster, placement); err == nil {
				candidates = append(candidates, cluster)
			}
		}
		klog.V(3).Infof("candidates: %v", candidates)

		scores, err := ssn.BatchClusterOrderFn(task, candidates, placement)
		if err != nil {
			klog.Errorf("calcate scores error for task %s: %v", task.Name, err)
		}
		klog.V(3).Infof("scores: %v", scores)

		// select clusters by score, only 0 and 100 for now
		// TODO: sort by socre
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
		klog.V(3).Infof("seeds: %v", seeds)

		// TODO: assign replicas to each clusters by weight
		replicas := []kworkv1alpha1.TargetCluster{}
		for _, cluster := range seeds {
			replicas = append(replicas, kworkv1alpha1.TargetCluster{
				Name:     cluster.Name,
				Replicas: task.ResourceBinding.Spec.Resource.Replicas,
			})
		}
		klog.V(3).Infof("target replicas: %v", replicas)

		// TODO: patch task with scheduler's decision
		newRb := task.ResourceBinding.DeepCopy()
		newRb.Spec.Clusters = replicas

		oldData, err := json.Marshal(task.ResourceBinding)
		if err != nil {
			klog.Errorf("failed to marshal the existing resource binding(%s/%s): %v", task.ResourceBinding.Namespace, task.Name, err)
		}
		newData, err := json.Marshal(newRb)
		if err != nil {
			klog.Errorf("failed to marshal the new resource binding(%s/%s): %v", newRb.Namespace, newRb.Name, err)
		}
		patchBytes, err := jsonpatch.CreateMergePatch(oldData, newData)
		if err != nil {
			klog.Errorf("failed to create a merge patch: %v", err)
		}
		klog.V(3).Infof("patch: %v", patchBytes)

		_, err = ssn.KarmadaClient().WorkV1alpha2().
			ResourceBindings(task.ResourceBinding.Namespace).
			Patch(context.TODO(), newRb.Name, types.MergePatchType, patchBytes, metav1.PatchOptions{})
		if err != nil {
			klog.Errorf("patch failed: %v", err)
		}
	}

}

func (alloc *Action) UnInitialize() {
	klog.V(3).Infof("Enter...")
	defer klog.V(3).Infof("Leaving...")
}

func (alloc *Action) getPlacement(ssn *framework.Session) *api.PlacementInfo {
	return nil
}
