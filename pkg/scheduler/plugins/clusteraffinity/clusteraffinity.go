package clusteraffinity

import (
	"k8s.io/klog"
	"volcano.sh/volcano/pkg/scheduler/api"
	"volcano.sh/volcano/pkg/scheduler/framework"
)

const PluginName = "cluster-affinity"

type clusterAffinity struct {
	pluginArguments framework.Arguments
}

func New(arguments framework.Arguments) framework.Plugin {
	return &clusterAffinity{
		pluginArguments: arguments,
	}
}

func (cl *clusterAffinity) Name() string {
	return PluginName
}

func (cl *clusterAffinity) OnSessionOpen(ssn *framework.Session) {
	klog.V(3).Infof("Enter ClusterAffinity Open...")
	defer klog.V(3).Infof("Leaving ClusterAffinity Close...")

	ssn.AddClusterPredicateFn(cl.Name(), func(*api.ClusterTaskInfo, *api.ClusterDetailInfo, *api.PlacementInfo) error {
		klog.V(3).Infof("enter cluster affinity predicates")
		return nil
	})

	ssn.AddBatchClusterOrderFn(cl.Name(), func(task *api.ClusterTaskInfo, clusters []*api.ClusterDetailInfo, placement *api.PlacementInfo) (map[string]float64, error) {
		klog.V(3).Infof("enter cluster affinity batch order")
		res := make(map[string]float64, len(clusters))
		for _, cluster := range clusters {
			res[cluster.Name] = 100
		}
		return res, nil
	})
}

func (cl *clusterAffinity) OnSessionClose(ssn *framework.Session) {

}
