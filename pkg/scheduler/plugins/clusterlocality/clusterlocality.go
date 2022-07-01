package clusterlocality

import (
	"k8s.io/klog"
	"volcano.sh/volcano/pkg/scheduler/api"
	"volcano.sh/volcano/pkg/scheduler/framework"
)

const PluginName = "cluster-locality"

type clusterLocality struct {
	pluginArguments framework.Arguments
}

func New(arguments framework.Arguments) framework.Plugin {
	return &clusterLocality{
		pluginArguments: arguments,
	}
}

func (cl *clusterLocality) Name() string {
	return PluginName
}

func (cl *clusterLocality) OnSessionOpen(ssn *framework.Session) {
	klog.V(3).Infof("Enter ClusterLocality Open...")
	defer klog.V(3).Infof("Leaving ClusterLocality Close...")

	ssn.AddClusterPredicateFn(cl.Name(), func(*api.ClusterTaskInfo, *api.ClusterDetailInfo, *api.PlacementInfo) error {
		klog.V(3).Infof("enter cluster locality predicates")
		return nil
	})

	ssn.AddBatchClusterOrderFn(cl.Name(), func(task *api.ClusterTaskInfo, clusters []*api.ClusterDetailInfo, placement *api.PlacementInfo) (map[string]float64, error) {
		klog.V(3).Infof("enter cluster locality batch order")
		res := make(map[string]float64, len(clusters))
		for _, cluster := range clusters {
			res[cluster.Name] = 0
			for _, target := range task.ResourceBinding.Spec().Clusters {
				if cluster.Name == target.Name {
					res[cluster.Name] = 100
				}
			}
		}
		return res, nil
	})
}

func (cl *clusterLocality) OnSessionClose(ssn *framework.Session) {

}
