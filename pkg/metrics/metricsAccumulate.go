/* UnderDevelopment file*/

package metrics

import (
    "fmt"
    "log"
    "strconv"

    getMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics/metricsServerMetrics"
    promMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics/prometheusMetrics"

)

var promServerIP string = "http://10.101.202.25:80/api/v1/query?query="
var kubeConfig string = "/home/ayush5588/go/src/github.com/ClusterAutoscaler/realKubeConfig.conf"


func Start() {
    
    // Check the Node status for PIDPressure, MemoryPressure, DiskPressure
    var tempNodeStatusArr []promMetrics.TempNodeStatusStruct
    tempNodeStatusArr, err := promMetrics.NodeStatusPhase(promServerIP)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, n := range tempNodeStatusArr {
        if n.ConditionStatus == "true" && n.ConditionStatusValue == "1" {
            if n.Condition == "PIDPressure" {
                // UPSCALE
                fmt.Printf("UPSCALE due to node -> %s\n", n.NodeName)
            }else if n.Condition == "MemoryPressure" {
                // UPSCALE
                fmt.Printf("UPSCALE due to node -> %s\n", n.NodeName)
            }else if n.Condition == "DiskPressure" {
                // UPSCALE
                fmt.Printf("UPSCALE due to node -> %s\n", n.NodeName)
            }else{
                fmt.Printf("No PIDPressure OR MemoryPressure OR DiskPressure in node -> %s\n", n.NodeName)
            }
        }
            //fmt.Println("NO Issues with PIDPressure OR MemoryPressure OR DiskPressure!!\n\n")
    }

    var tempPodsNotScheduledArr []promMetrics.TempPodsNotScheduledStruct
    tempPodsNotScheduledArr, err = promMetrics.PodsNotScheduled(promServerIP)
    if err != nil {
        log.Fatal(err)
    }

    var cntUnscheduledPods int = 0
    for _, p := range tempPodsNotScheduledArr {
        if p.PodName != "" {
            cntUnscheduledPods += 1
            break
        }
    }

    if cntUnscheduledPods >= 1 {
        // UPSCALE
        fmt.Println("\n\nUnscheduled Pods in the Cluster -> UPSCALE\n\n")
    }else {
        fmt.Println("\n\nNo unscheduled pods\n\n")
    }
    


    //  Calculate the current utilization for the CPU and MEM

    /*  Get the metrics from the Metrics-Server and the resources existing to be allocated from Prometheus 
        and calculate the percentage to get the Current Utilization of the resources. 
    */


    // Get allocatable resource metrics from Prometheus
    var tempNodeResArr []promMetrics.TempNodeResourceStruct
    tempNodeResArr, err = promMetrics.NodeAllocatableResources(promServerIP)
    if err != nil {
        log.Fatal(err)
    }

    mp := make(map[string]float64)
    for _, node := range tempNodeResArr {
        nodeName := node.NodeName
        nodeRes := node.Resource
        _, exist := mp[nodeName+"-"+nodeRes]
        if !exist {
            str1 := node.ResourceAvailable
            allocatableResource, err := strconv.ParseFloat(str1,64)
            if err != nil {
                log.Fatal(err)
            }
            mp[nodeName+"-"+nodeRes] = allocatableResource
        }
    }



    // Get current usage of resource metrics from Metrics-Server
    var nodeArr []getMetrics.NodeUsage
    nodeArr, err = getMetrics.GetNodeMetrics(kubeConfig)
    if err != nil {
        log.Fatal(err)
    }



    // Check the utilization of the resources to decide whether to UPSCALE or DOWNSCALE
    for _, n := range nodeArr{
        fmt.Printf("NODE -> %s\n", n.NodeName)

        var upscale bool = false
        var downscale bool = false

        // Memory Utilization
        memAllocatable, exist :=  mp[n.NodeName+"-"+"memory"]
        if exist {
            memUsage := float64(n.NodeMemUsage * 1024)  // Converting to bytes (from Kibibyte)

            memoryUtilization := ( memUsage / memAllocatable) * 100
            fmt.Printf("Memory Utilization = %0.2f\n", memoryUtilization)
            if memoryUtilization >= 50.00 {
                upscale = true
            }else if memoryUtilization <= 20.00 {
                // Pre-Checks such as PodDisruptionBudget (PDB) to be done before deciding to DOWNSCALE
                downscale = true
            }else {
                //fmt.Printf("Memory Utilization is NETURAL: %0.2f\n", memoryUtilization)
            }
        }

        // CPU Utilization
        cpuAllocatableCores, exist := mp[n.NodeName+"-"+"cpu"]
        if exist {
            // convert CpuUsage UNIT from nanocores to cores by dividing it by 1 billion
            cpuUsageCores := float64(n.NodeCpuUsage) / 1000000000

            cpuUtilization := (cpuUsageCores / cpuAllocatableCores) * 100
            fmt.Printf("CPU Utilization = %0.2f\n", cpuUtilization)
            if cpuUtilization >= 50.00 {
                upscale = true
            }else if cpuUtilization <= 20.00 {
                // Pre-Checks such as PodDisruptionBudget (PDB) to be done before deciding to DOWNSCALE
                downscale = true
            }else {
                //fmt.Printf("CPU Utilization is NEUTRAL: %0.2f\n\n", cpuUtilization)
            }
        }

        if upscale {
            fmt.Println("UPSCALE\n\n")
        }else {
            if downscale {
                fmt.Println("DOWNSCALE\n\n")
            }else {
                fmt.Println("NEUTRAL\n\n")
            }
        }
    }

}

