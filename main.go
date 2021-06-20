package main

import(
    "fmt"
    "log"

    //GetMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics"
    promMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics/prometheusMetrics"
    )

var kubeConfig string = "/home/ayush5588/go/src/github.com/ClusterAutoscaler/realKubeConfig.conf"

// Just for the dev env. WIll be replaced with env variables at the end
var promServerIP string = "http://10.101.202.25:80/api/v1/query?query="


func main() {
    
        // Get Pod Status
        var podStatusPhaseArr []promMetrics.TempPodStatusStruct
        podStatusPhaseArr, err1 := promMetrics.PodStatusPhase(promServerIP)
        if err1 != nil {
            log.Fatal(err1)
        }
        for _, p := range podStatusPhaseArr {
            fmt.Printf("PodName: %s\nPodPhase: %s\nPhaseValue: %s\n\n", p.PodName, p.PodPhase, p.PhaseValue)
        }

        fmt.Println("\n\n--------------------------- NODE STATUS CONDITION INFO ----------------------------\n\n")
        // Get Node Status
        var nodeStatusPhaseArr []promMetrics.TempNodeStatusStruct
        nodeStatusPhaseArr, err2 := promMetrics.NodeStatusPhase(promServerIP)
        if err2 != nil {
            log.Fatal(err2)
        }
        for _, n := range nodeStatusPhaseArr {
            fmt.Printf("NodeName: %s\nCondition: %s\nConditionStatus: %s\nConditionStatusValue: %s\n\n", n.NodeName, n.Condition, n.ConditionStatus, n.ConditionStatusValue)
        }
    
         /*
         // UNCOMMENT to get the Pod Metrics
         var arr []GetMetrics.PodUsage
         arr, err := GetMetrics.GetPodMetrics(kubeConfig)
         if err != nil {
            log.Fatal(err)
         }
         for _, p := range arr {
            fmt.Printf("Name: %s\nCPU Usage: %dn\nMemory Usage: %dki\n\n",p.PodName,p.PodCpuUsage,p.PodMemUsage)
         }
         */

        
        /*
         // Get Node Metrics
         var nodeArr []GetMetrics.NodeUsage
         nodeArr, err := GetMetrics.GetNodeMetrics(kubeConfig)
         if err != nil {
            log.Fatal(err)
         }
         for _, n := range nodeArr {
            fmt.Printf("Name: %s\nCPU Usage: %dn\nMemory Usage: %dki\n\n",n.NodeName,n.NodeCpuUsage,n.NodeMemUsage)
         }
         */
    

}
