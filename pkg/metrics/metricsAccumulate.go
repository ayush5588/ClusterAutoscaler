/* UnderDevelopment file*/

package metrics

import (
    "fmt"
    "log"
    "strconv"

    getMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics/metricsServerMetrics"
    promMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics/prometheusMetrics"
    getNodeList "github.com/ayush5588/ClusterAutoscaler/pkg/podNodeList"
)

var promServerIP string = "http://10.101.202.25:80/api/v1/query?query="
var kubeConfig string = "/home/ayush5588/go/src/github.com/ClusterAutoscaler/realKubeConfig.conf"


func shiftToAnotherNode (nodeName string, allocatableMap map[string]float64) (string, error) {
    var finalNode string = ""

    NodeList, err := getNodeList.GetItems("node", kubeConfig)
    if err != nil {
        return "",err
    }
    requestMap, err := promMetrics.NodeResourceRequest(promServerIP)
    if err != nil {
        return "", err
    }
    
    requiredMEM := requestMap[nodeName+"-"+"memory"]
    requiredCPU := requestMap[nodeName+"-"+"cpu"]

    for _, node := range NodeList {
        if node != nodeName && node != "master"{
            memLeft := allocatableMap[node+"-"+"memory"] - requestMap[node+"-"+"memory"]
            cpuLeft := allocatableMap[node+"-"+"cpu"] - requestMap[node+"-"+"cpu"]

            if requiredMEM <= memLeft && requiredCPU <= cpuLeft {
                finalNode = node
                break
            }
        }
    }

    return finalNode, nil
}

func Start() {
    
    // Check the Node status for PIDPressure, MemoryPressure, DiskPressure
    var tempNodeStatusArr []promMetrics.TempNodeStatusStruct
    tempNodeStatusArr, err := promMetrics.NodeStatusPhase(promServerIP)
    if err != nil {
        log.Fatal(err)
    }
    
    PIDStatusMap := make(map[string]bool)
    MEMStatusMap := make(map[string]bool)
    DISKStatusMap := make(map[string]bool)

    for _, n := range tempNodeStatusArr {
        if n.ConditionStatus == "true" && n.ConditionStatusValue == "1" {
            if n.Condition == "PIDPressure" {
                // UPSCALE
                //fmt.Printf("UPSCALE due to node -> %s\n", n.NodeName)
                PIDStatusMap[n.NodeName] = true
            }else if n.Condition == "MemoryPressure" {
                // UPSCALE
                //fmt.Printf("UPSCALE due to node -> %s\n", n.NodeName)
                MEMStatusMap[n.NodeName] = true
            }else if n.Condition == "DiskPressure" {
                // UPSCALE
                //fmt.Printf("UPSCALE due to node -> %s\n", n.NodeName)
                DISKStatusMap[n.NodeName] = true
            }else{
                //fmt.Printf("No PIDPressure OR MemoryPressure OR DiskPressure in node -> %s\n", n.NodeName)
                PIDStatusMap[n.NodeName] = false
                MEMStatusMap[n.NodeName] = false
                DISKStatusMap[n.NodeName] = false
            }
        }
            //fmt.Println("NO Issues with PIDPressure OR MemoryPressure OR DiskPressure!!\n\n")
    }

    fmt.Println("NODE\t\tPIDPressure\tMEMPressure\tDISKPressure\tStatus\n")
    for nodeName, _ := range PIDStatusMap {
        if nodeName == "master" {
            fmt.Printf("%s\t\t", nodeName)
        }else {
            fmt.Printf("%s\t", nodeName)
        }

        fmt.Printf("%v\t\t%v\t\t%v\t\t", PIDStatusMap[nodeName], MEMStatusMap[nodeName], DISKStatusMap[nodeName])
        if PIDStatusMap[nodeName] || MEMStatusMap[nodeName] || DISKStatusMap[nodeName] {
            fmt.Println("UPSCALE")
        }else {
            fmt.Println("NEUTRAL")
        }
    }

    fmt.Println("\n\n")

    var tempPodsNotScheduledArr []promMetrics.TempPodsNotScheduledStruct
    tempPodsNotScheduledArr, err = promMetrics.PodsNotScheduled(promServerIP)
    if err != nil {
        log.Fatal(err)
    }

    var  UnscheduledPodsArr []string
    var cntUnscheduledPods int = 0
    for _, p := range tempPodsNotScheduledArr {
        if p.PodName != "" {
            cntUnscheduledPods += 1
            UnscheduledPodsArr = append(UnscheduledPodsArr, p.PodName)
           //fmt.Printf("\nUnscheduled Pod Name: %s", p.PodName)
            //break
        }
    }

    if cntUnscheduledPods >= 1 {
        // UPSCALE
        /*
        if cntUnscheduledPods == 1 {
            fmt.Printf("\n\n%d Unscheduled Pods in the Cluster -> UPSCALE\n\n", cntUnscheduledPods)
        }*/
        fmt.Println("S.No\tUnscheduled POD\n")
        for idx, podName := range UnscheduledPodsArr {
            fmt.Printf("%d\t%s\n", idx+1, podName)
        }
        fmt.Println("\n\nSTATUS: UPSCALE\n\n")
    }else {
        fmt.Println("No unscheduled pods\n")
        fmt.Println("\nSTATUS: NEUTRAL \n\n")
    }

    
    fmt.Println("\n")

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
    
    nodeAction := make(map[string]string)
    nodeRemarks := make(map[string]string)

    fmt.Println("\nNODE\t\tMEMORY%\t\tCPU%\n")

    // Check the utilization of the resources to decide whether to UPSCALE or DOWNSCALE
    for _, n := range nodeArr{
        //fmt.Printf("\nNODE -> %s\t", n.NodeName)
        if n.NodeName == "master" {
            fmt.Printf("%s\t\t", n.NodeName)
        } else {
            fmt.Printf("%s\t", n.NodeName)
        }
        //fmt.Printf("%s\t", n.NodeName)

        var upscale bool = false
        var downscale bool = false
        

        // Memory Utilization
        memAllocatable, exist :=  mp[n.NodeName+"-"+"memory"]
        if exist {
            memUsage := float64(n.NodeMemUsage * 1024)  // Converting to bytes (from Kibibyte)

            memoryUtilization := ( memUsage / memAllocatable) * 100
            //fmt.Printf("Memory Utilization = %0.2f\t", memoryUtilization)
            fmt.Printf("%0.2f\t", memoryUtilization)

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
            //cpuNodeUsage = cpuUsageCores

            cpuUtilization := (cpuUsageCores / cpuAllocatableCores) * 100
            //fmt.Printf("CPU utilization : %0.2f\t\n", cpuUtilization)
            fmt.Printf("\t%0.2f\n", cpuUtilization)

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
            //fmt.Println("\nUPSCALE\n\n")
        }else {
            if downscale {
                if n.NodeName != "master" {
                    destinationNode, err := shiftToAnotherNode(n.NodeName, mp)

                    if err != nil {
                        log.Fatal(err)
                    }

                    if destinationNode != "" {
                        nodeAction[n.NodeName] = "DOWNSCALE"
                        nodeRemarks[n.NodeName] = "Can move pods to " + destinationNode
                        //fmt.Println("\nDOWNSCALE")
                        //fmt.Printf("Can move pods to node: %s\n\n\n", destinationNode)
                    } else {
                        nodeAction[n.NodeName] = "CANNOT DOWNSCALE"
                        nodeRemarks[n.NodeName] = "Cannot downscale as no Nodes fulfill the resource requirement"
                        //fmt.Println("\nUnder utilized node but cannot DOWNSCALE as pods cannot be moved to other nodes\n\n")
                    }
                    //fmt.Println("DOWNSCALE\n\n")

                } else {
                    //fmt.Println("\nUnder utilized node but cannot DOWNSCALE as pods of MASTER node cannot be moved to other nodes\n\n")
                }
            }else {
                nodeAction[n.NodeName] = "NEUTRAL"
                nodeRemarks[n.NodeName] = "\tNo action needed"
                //fmt.Println("\nNEUTRAL\n\n")
            }
        }
    }

    fmt.Println("\n\n")

    fmt.Println("\nNODE\t\t STATUS\t\t\tREMARKS\n")
    for val, _ := range nodeAction {
        fmt.Printf("%s\t", val)
        fmt.Printf("%s\t", nodeAction[val])
        fmt.Printf("%s\n", nodeRemarks[val])
    }

    fmt.Println("\n\n")

}

