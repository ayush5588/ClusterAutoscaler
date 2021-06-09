package main

import(
    "fmt"
    "log"

    GetMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics"
    )

var kubeConfig string = "/home/ayush5588/go/src/github.com/ClusterAutoscaler/realKubeConfig.conf"

func main() {

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


        /* Get Node Metrics */
         var nodeArr []GetMetrics.NodeUsage
         nodeArr, err := GetMetrics.GetNodeMetrics(kubeConfig)
         if err != nil {
            log.Fatal(err)
         }
         for _, n := range nodeArr {
            fmt.Printf("Name: %s\nCPU Usage: %dn\nMemory Usage: %dki\n\n",n.NodeName,n.NodeCpuUsage,n.NodeMemUsage)
         }

}
