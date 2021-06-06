package objectMetrics

import (
    "context"
    "encoding/json"
    //"fmt"
    "strconv"
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

type PodUsage struct {
    PodName string
    PodCpuUsage int
    PodMemUsage int
}

func getMetrics(clientset *kubernetes.Clientset, pods *PodMetricsList) error {
    data, err:= clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").Do(context.TODO()).Raw()
    if err != nil {
        return err
    }
    //fmt.Printf("Type of data: %T\n\n",data)
    err = json.Unmarshal(data, &pods)
    return err
    }

func GetObjectMetrics() ([]PodUsage, error) {
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
    var pods PodMetricsList
    err = getMetrics(clientset, &pods)
    if err != nil {
        panic(err.Error())
    }

    var PodMetricsArr []PodUsage

    for _, m := range pods.Items {
        //fmt.Println(m.Metadata.Name, m.Metadata.Namespace)
        var PodMetrics PodUsage
        PodMetrics.PodName = m.Metadata.Name
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

        PodMetrics.PodCpuUsage = cUsage
        PodMetrics.PodMemUsage = mUsage
        PodMetricsArr = append(PodMetricsArr,PodMetrics)
    }
    
    return PodMetricsArr,nil
}
