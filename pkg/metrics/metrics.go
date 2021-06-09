/* metrics pkg is to get the resource metrics usage of pods and nodes */

package metrics

import (
    "context"
    "encoding/json"
    "strconv"

    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/kubernetes"
    metricsStruct "github.com/ayush5588/ClusterAutoscaler/pkg/metrics/metricsStruct"
)

type PodUsage struct {
    PodName string
    PodCpuUsage int
    PodMemUsage int
}


type NodeUsage struct {
    NodeName string
    NodeCpuUsage int
    NodeMemUsage int
}


/* Gets the Pod metrics from the specified API */
func getPodMetrics(clientset *kubernetes.Clientset, pods *metricsStruct.PodMetricsStruct, apiPath string) error {
    data, err:= clientset.RESTClient().Get().AbsPath(apiPath).Do(context.TODO()).Raw()
    if err != nil {
        return err
    }
    err = json.Unmarshal(data, &pods)
    return err
    }


/* Gets the Node metrics from the specified API */
func getNodeMetrics(clientset *kubernetes.Clientset, nodes *metricsStruct.NodeMetricsStruct, apiPath string) error {
    data, err := clientset.RESTClient().Get().AbsPath(apiPath).Do(context.TODO()).Raw()
    if err != nil {
        return err
    }
    err = json.Unmarshal(data, &nodes)
    return err
}

// kubeConfig file path
//var kubeconfig string = "/home/ayush5588/go/src/github.com/ClusterAutoscaler/realKubeConfig.conf"

/* 
    1. Creates a config object from the kubeconfig file provided
    2. Creates a clientset (groups of client) from the config object
    3. Calls the appropriate function to get the metrics from the given API 
    4. Creates a slice of struct elements having resource usage and returns it back
*/
func GetNodeMetrics(kubeConfig string) ([]NodeUsage, error) {
    config, err := clientcmd.BuildConfigFromFlags("",kubeConfig)
    if err != nil {
        return nil, err
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, err
    }
    var nodes metricsStruct.NodeMetricsStruct
    err = getNodeMetrics(clientset, &nodes, "apis/metrics.k8s.io/v1beta1/nodes")

    var NodeMetricsArr []NodeUsage
    for _, m := range nodes.Items {
        var NodeMetrics NodeUsage
        
        NodeMetrics.NodeName = m.Metadata.Name
        
        fmtNodeCPU := m.Usage.CPU[:len(m.Usage.CPU)-1]
        if len(fmtNodeCPU) > 0 {
            ci, err := strconv.Atoi(fmtNodeCPU)
            if err != nil {
                return nil, err
            }
            NodeMetrics.NodeCpuUsage = ci
        }
        
        fmtNodeMem := m.Usage.Memory[:len(m.Usage.Memory)-2]
        if len(fmtNodeMem) > 0 {
            cm, err := strconv.Atoi(fmtNodeMem)
            if err != nil {
                return nil, err
            }
            NodeMetrics.NodeMemUsage = cm
        }

        NodeMetricsArr = append(NodeMetricsArr,NodeMetrics)
    }
    return NodeMetricsArr, nil
}


func GetPodMetrics(kubeConfig string) ([]PodUsage, error) {
    config, err := clientcmd.BuildConfigFromFlags("",kubeConfig)
    if err != nil {
        panic(err.Error())
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }
    var pods metricsStruct.PodMetricsStruct
    err = getPodMetrics(clientset, &pods, "apis/metrics.k8s.io/v1beta1/pods")
    if err != nil {
        return nil, err
    }

    var PodMetricsArr []PodUsage

    for _, m := range pods.Items {
    var PodMetrics PodUsage
    PodMetrics.PodName = m.Metadata.Name
    var cUsage int = 0
    var mUsage int = 0
    /* Goes through all the containers in the pod to gather the metrics usage */
    for _,c := range m.Containers {
        /* removes the last character (or last 2 in case of memory) which denotes the unit of 
           measurement (such as n in 34334n i.e. CPU usage) of the usage in order to sum it up
        */

        fmtPodCPU := c.Usage.CPU[:len(c.Usage.CPU)-1]
        if len(fmtPodCPU) > 0 {
            ci, err := strconv.Atoi(fmtPodCPU)
            if err != nil {
                return nil, err
            }
            cUsage = cUsage + ci
        }
        
        fmtPodMem := c.Usage.Memory[:len(c.Usage.Memory)-2]
        if len(fmtPodMem) > 0 {
            mi, err := strconv.Atoi(fmtPodMem)
            if err != nil {
                return nil, err
            }
            mUsage = mUsage + mi
            }
        }

        PodMetrics.PodCpuUsage = cUsage
        PodMetrics.PodMemUsage = mUsage
        PodMetricsArr = append(PodMetricsArr,PodMetrics)
    }

    return PodMetricsArr,nil
}


