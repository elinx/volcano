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
	"k8s.io/klog"

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
	klog.V(3).Infof("Enter...")
	defer klog.V(3).Infof("Leaving...")

	klog.V(3).Infof("clusters: %v", ssn.Clusters)
	klog.V(3).Infof("cluster tasks: %v", ssn.ClusterTasks)
	klog.V(3).Infof("placement: %v", ssn.Placements)
}

func (alloc *Action) UnInitialize() {
	klog.V(3).Infof("Enter...")
	defer klog.V(3).Infof("Leaving...")
}
