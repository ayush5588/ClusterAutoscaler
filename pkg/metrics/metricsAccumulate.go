/* UnderDevelopment file*/

package metrics

import (
    "fmt"
    "log"
    "strconv"
    
//    metricsStruct "github.com/ayush5588/ClusterAutoscaler/metricsStruct"
    getMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics/metricsServerMetrics"
    promMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics/prometheusMetrics"
)

var promServerIP string = "http://10.101.202.25:80/api/v1/query?query="
var kubeConfig string = "/home/ayush5588/go/src/github.com/ClusterAutoscaler/realKubeConfig.conf"



func Start() {
    
    // Checking for the Node status 
    var tempNodeStatusArr []promMetrics.TempNodeStatusStruct
    tempNodeStatusArr, err := promMetrics.NodeStatusPhase(promServerIP)
    if err != nil {
        log.Fatal(err)
    }
    for _, n := range tempNodeStatusArr {
        if n.ConditionStatus == "true" && n.ConditionStatusValue == "1" {
            if n.Condition == "PIDPressure" {
                // UPSCALE
            }else if n.Condition == "MemoryPressure" {
                // UPSCALE
            }else if n.Condition == "DiskPressure" {
                // UPSCALE
            }else{
                fmt.Printf("No PIDPressure OR MemoryPressure OR DiskPressure in node -> %s\n", n.NodeName)
            }
        }
            //fmt.Println("NO Issues with PIDPressure OR MemoryPressure OR DiskPressure!!\n\n")
    }

    fmt.Println("\n\n")

    // Checking for the Unscheduled pods 
    var tempPodsNotScheduledArr []promMetrics.TempPodsNotScheduledStruct
    tempPodsNotScheduledArr, err = promMetrics.PodsNotScheduled(promServerIP)
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Println(tempPodsNotScheduledArr)
    var cntUnscheduledPods int = 0
    for _, p := range tempPodsNotScheduledArr {
        if p.Namespace != "" {
            cntUnscheduledPods += 1
            break
        }
    }
   if cntUnscheduledPods >= 1 {
        // UPSCALE
   }else {
        fmt.Println("No unscheduled pods\n\n")
   }


    //  Calculate the current utilization for the CPU and MEM

    /*  Get the metrics from the Metrics-Server and the resources existing to be allocated from Prometheus 
        (will have to add logic in the promMetrics.go) and calculate the percentage to get the Current Utilization 
        of the resources. 
    */

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
    
    var nodeArr []getMetrics.NodeUsage
    nodeArr, err = getMetrics.GetNodeMetrics(kubeConfig)
    if err != nil {
        log.Fatal(err)
    }
    
    //fmt.Println(len(tempNodeResArr))


    /*
    nodeMap := make(map[string][]string)
    nodeMap, err = promMetrics.PodInNodes(promServerIP)
    if err != nil {
        log.Fatal(err)
    }

    for nodeName, podsArr := range nodeMap {
        fmt.Println(nodeName)
        for _, p := range podsArr {
            fmt.Printf("%s\n", p)
        }
        fmt.Println("\n\n")
    }

    fmt.Println("\n\n\n")
    */
    

    
    for _, n := range nodeArr{
        fmt.Printf("NODE -> %s\n", n.NodeName)
        // Memory utilization
        memAllocatable, exist := mp[n.NodeName+"-"+"memory"]
        if exist {
            memUsage := float64(n.NodeMemUsage * 1024)  // Converting to bytes (from Kibibyte)

            memoryUtilization := ( memUsage / memAllocatable) * 100

            if memoryUtilization >= 75.00 {
                // UPSCALE
            }else if memoryUtilization <= 20.00 {
                // Pre-Checks such as PodDisruptionBudget (PDB) to be done before deciding to DOWNSCALE
                fmt.Printf("Memory Utilization is BELOW threshold: %0.2f\n", memoryUtilization)
            }else {
                fmt.Printf("Memory Utilization is NETURAL: %0.2f\n", memoryUtilization)
            }
        }
        
        // CPU Utilization
        cpuAllocatableCores, exist := mp[n.NodeName+"-"+"cpu"]
        if exist {
            // convert CpuUsage UNIT from nanocores to cores by dividing it by 1 billion
            cpuUsageCores := float64(n.NodeCpuUsage) / 1000000000

            cpuUtilization := (cpuUsageCores / cpuAllocatableCores) * 100

            if cpuUtilization >= 75.00 {
                // UPSCALE
            }else if cpuUtilization <= 20.00 {
                // Pre-Checks such as PodDisruptionBudget (PDB) to be done before deciding to DOWNSCALE
                fmt.Printf("CPU Utilization is BELOW threshold: %0.2f\n\n", cpuUtilization)
            }else {
                fmt.Printf("CPU Utilization is NEUTRAL: %0.2f\n\n", cpuUtilization)
            }
        }
    }
    

    /*
    for _, n1 := range tempNodeResArr {
        for _, n2 := range nodeArr {
            if n1.NodeName == n2.NodeName {
                if n1.Resource == "memory" {
                    //fmt.Printf("Node Name: %s\n", n1.NodeName)
                    str1 := n1.ResourceAvailable
                    memAllocatable, err := strconv.ParseFloat(str1,64)
                    if err != nil {
                        log.Fatal(err)
                    }

                    memUsage := float64(n2.NodeMemUsage * 1024)  // Converting to bytes (from Kibibyte)

                    memoryUtilization := ( memUsage / memAllocatable) * 100
                    if memoryUtilization >= 75.00 {
                        // UPSCALE
                    }else if memoryUtilization <= 20.00 {
                        // Pre-Checks such as PodDisruptionBudget (PDB) to be done before deciding to DOWNSCALE
                        //fmt.Printf("Memory Utilization is BELOW threshold: %0.2f\n\n", memoryUtilization)
                    }else {
                        //fmt.Printf("Memory Utilization is NETURAL: %0.2f\n\n", memoryUtilization)
                    }

                }else if n1.Resource == "cpu" {
                    //fmt.Printf("Node Name: %s\n", n1.NodeName)
                    cpuAllocatableCores, err := strconv.ParseFloat(n1.ResourceAvailable, 64)
                    if err != nil {
                        log.Fatal(err)
                    }
                    // convert CpuUsage UNIT from nanocores to cores by dividing it by 1 billion
                    cpuUsageCores := float64(n2.NodeCpuUsage) / 1000000000

                    cpuUtilization := (cpuUsageCores / cpuAllocatableCores) * 100

                    if cpuUtilization >= 75.00 {
                        // UPSCALE
                    }else if cpuUtilization <= 20.00 {
                        // Pre-Checks such as PodDisruptionBudget (PDB) to be done before deciding to DOWNSCALE
                        //fmt.Printf("CPU Utilization is BELOW threshold: %0.2f\n", cpuUtilization)
                    }else {
                        //fmt.Printf("CPU Utilization is NEUTRAL: %0.2f\n", cpuUtilization)
                    }

                }
            }
        }
    }
    */

}


