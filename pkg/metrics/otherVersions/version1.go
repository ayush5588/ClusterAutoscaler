/* This is a test file */

package main

import (
    "context"
    "encoding/json"
    "fmt"
    //"strconv"
    "time"

    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/kubernetes"
    //"k8s.io/client-go/rest"
)

// PodMetricsList : PodMetricsList
type PodMetricsList struct {
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

type NodeStatusStruct struct {
	Status string `json:"status"`
	Data   struct {
		Resulttype string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name                     string `json:"__name__"`
				AppKubernetesIoInstance  string `json:"app_kubernetes_io_instance"`
				AppKubernetesIoManagedBy string `json:"app_kubernetes_io_managed_by"`
				AppKubernetesIoName      string `json:"app_kubernetes_io_name"`
				ContainerRuntimeVersion  string `json:"container_runtime_version"`
				HelmShChart              string `json:"helm_sh_chart"`
				Instance                 string `json:"instance"`
				InternalIP               string `json:"internal_ip"`
				Job                      string `json:"job"`
				KernelVersion            string `json:"kernel_version"`
				KubeletVersion           string `json:"kubelet_version"`
				KubeproxyVersion         string `json:"kubeproxy_version"`
				KubernetesName           string `json:"kubernetes_name"`
				KubernetesNamespace      string `json:"kubernetes_namespace"`
				KubernetesNode           string `json:"kubernetes_node"`
				Node                     string `json:"node"`
				OsImage                  string `json:"os_image"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type PodUsage struct {
    Name string
    cpuUsage int
    memUsage int
}

func getMetrics(clientset *kubernetes.Clientset, pods *NodeStatusStruct) error {
    // Using Prometheus-Server service ClusterIP address (kubectl get svc)
    data, err:= clientset.RESTClient().Get().AbsPath("http://10.103.151.144:80/api/v1/query?query=kube_node_info").Do(context.TODO()).Raw()
    if err != nil {
        return err
    }
    //fmt.Printf("Type of data: %T\n\n",data)
    err = json.Unmarshal(data, &pods)
    return err
    }

// (PodUsage[], error)
func main()  {
    kubeconfig := "/home/ayush5588/go/src/github.com/ClusterAutoscaler/realKubeConfig.conf"
    config, err := clientcmd.BuildConfigFromFlags("",kubeconfig)
    if err != nil {
        panic(err.Error())
    }
    // creates the clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }
    var pods NodeStatusStruct
    err = getMetrics(clientset, &pods)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(pods)
    /*
    var PodMetricsArr []PodUsage

    for _, m := range pods.Items {
        //fmt.Println(m.Metadata.Name, m.Metadata.Namespace)
        var PodMetrics PodUsage
        PodMetrics.Name = m.Metadata.Name
        var cUsage int = 0
        var mUsage int = 0
        for _,c := range m.Containers {
            //fmt.Printf("Name: %s\n",c.Name)
            //fmt.Printf("CPU Usage: %s\n",c.Usage.CPU)
            //fmt.Printf("Memory Usage: %s\n\n",c.Usage.Memory)
            fmtCPU := c.Usage.CPU[:len(c.Usage.CPU)-1]
            if len(fmtCPU) > 0 {
                ci, err := strconv.Atoi(fmtCPU)
                if err != nil {
                    //panic(err.Error())
                    return nil, err
                }
                cUsage = cUsage + ci
            }
    
            fmtMem := c.Usage.Memory[:len(c.Usage.Memory)-2]
            if len(fmtMem) > 0 {
                mi, err := strconv.Atoi(fmtMem)
                if err != nil {
                    //panic(err.Error())
                    return nil, err
                }
                mUsage = mUsage + mi
            }
        }

        PodMetrics.cpuUsage = cUsage
        PodMetrics.memUsage = mUsage
        PodMetricsArr = append(PodMetricsArr,PodMetrics)
    }
    
    return PodMetricsArr,nil
    */
}
