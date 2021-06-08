/* objectMetrics pkg is to get the resource metrics usage of pods and nodes */

package objectMetrics

import (
    "context"
    "encoding/json"
    "strconv"
    "time"

    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/kubernetes"
)


type PodMetricsStruct struct {
    Kind       string `json:"kind"`
    APIVersion string `json:"apiVersion"`
    Metadata   struct {
        SelfLink string `json:"selfLink"`
    } `json:"metadata"`
    Items []struct {
        Metadata struct {
            Name              string    `json:"name"`
            Namespace         string    `json:"namespace"`
            SelfLink          string    `json:"selfLink"`
            CreationTimestamp time.Time `json:"creationTimestamp"`
        } `json:"metadata"`
        Timestamp  time.Time `json:"timestamp"`
        Window     string    `json:"window"`
        Containers []struct {
            Name  string `json:"name"`
            Usage struct {
                CPU    string `json:"cpu"`
                Memory string `json:"memory"`
            } `json:"usage"`
        } `json:"containers"`
    } `json:"items"`
}


type NodeMetricsStruct struct {
	Kind       string `json:"kind"`
	Apiversion string `json:"apiVersion"`
	Metadata   struct {
		Selflink string `json:"selfLink"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Selflink          string    `json:"selfLink"`
			Creationtimestamp time.Time `json:"creationTimestamp"`
		} `json:"metadata"`
		Timestamp time.Time `json:"timestamp"`
		Window    string    `json:"window"`
		Usage     struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"usage"`
	} `json:"items"`
}


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
func getPodMetrics(clientset *kubernetes.Clientset, pods *PodMetricsStruct, apiPath string) error {
    data, err:= clientset.RESTClient().Get().AbsPath(apiPath).Do(context.TODO()).Raw()
    if err != nil {
        return err
    }
    err = json.Unmarshal(data, &pods)
    return err
    }


/* Gets the Node metrics from the specified API */
func getNodeMetrics(clientset *kubernetes.Clientset, nodes *NodeMetricsStruct, apiPath string) error {
    data, err := clientset.RESTClient().Get().AbsPath(apiPath).Do(context.TODO()).Raw()
    if err != nil {
        return err
    }
    err = json.Unmarshal(data, &nodes)
    return err
}

// kubeConfig file path
var kubeconfig string = "/home/ayush5588/go/src/github.com/ClusterAutoscaler/kubeConfig.conf"

/* 
    1. Creates a config object from the kubeconfig file provided
    2. Creates a clientset (groups of client) from the config object
    3. Calls the appropriate function to get the metrics from the given API 
    4. Creates a slice of struct elements having resource usage and returns it back
*/
func GetNodeMetrics() ([]NodeUsage, error) {
    config, err := clientcmd.BuildConfigFromFlags("",kubeconfig)
    if err != nil {
        return nil, err
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, err
    }
    var nodes NodeMetricsStruct
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


func GetPodMetrics() ([]PodUsage, error) {
    config, err := clientcmd.BuildConfigFromFlags("",kubeconfig)
    if err != nil {
        panic(err.Error())
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }
    var pods PodMetricsStruct
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


