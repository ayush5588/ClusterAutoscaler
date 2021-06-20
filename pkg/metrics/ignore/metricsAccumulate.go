package metrics


import (
    "fmt"
    "log"
    
    metricsStruct "github.com/ayush5588/ClusterAutoscaler/metricsStruct"
    getMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics/metricsServerMetrics"
    promMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics/prometheusMetrics"
)

var promServerIP string = "http://10.101.202.25:80/api/v1/query?query="
var kubeConfig string = "/home/ayush5588/go/src/github.com/ClusterAutoscaler/realKubeConfig.conf"

// metricsStruct.Metrics, error
func main() {
    var FinalMetricsStruct metricsStruct.Metrics
    
    FinalMetricsStruct.UnscheduledPodsCNT = 0
    FinalMetricsStruct.CpuUtilizationCNT = 0
    FinalMetricsStruct.MemUtilizationCNT = 0
    FinalMetricsStruct.PIDPressureStatus = false
    FinalMetricsStruct.DISKPressureStatus = false
    FinalMetricsStruct.MEMPressureStatus = false
    FinalMetricsStruct.StructSET = false

    // Checking for the Node status 
    var tempNodeStatusArr []promMetrics.TempNodeStatusStruct
    tempNodeStatusArr, err := promMetrics.NodeStatusPhase(promServerIP)
    if err != nil {
        log.Fatal(err)
        //return FinalMetricsStruct, err
    }
    for _, n := range tempNodeStatusArr {
        if n.ConditionStatus == "true" && n.ConditionStatusValue == "1" {
            if n.Condition == "PIDPressure" {
                FinalMetricsStruct.PIDPressureStatus = true
                FinalMetricsStruct.StructSET = true
            }else if n.Condition == "MemoryPressure" {
                FinalMetricsStruct.MEMPressureStatus = true
                FinalMetricsStruct.StructSET = true
            }else if n.Condition == "DiskPressure" {
                FinalMetricsStruct.DISKPressureStatus = true
                FinalMetricsStruct.StructSET = true
            }
    }else {
        fmt.Println("NO Issues!!")
    }

    // Checking for the Unscheduled pods 
    var tempPodsNotScheduledArr []promMetrics.TempPodsNotScheduledStruct
    tempPodsNotScheduledArr, err = promMetrics.PodsNotScheduled(promServerIP)
    if err != nil {
        log.Fatal(err)
        //FinalMetricsStruct.StructSET = false
        //return FinalMetricsStruct, err
    }
    if len(tempPodsNotScheduledArr) > 0 {
        FinalMetricsStruct.UnscheduledPodsCNT = 1;
        FinalMetricsStruct.StructSET = true
    }else {
        fmt.Println("NO unscheduled Pods !!")
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
        //return FinalMetricsStruct, err
    }

    var nodeArr []getMetrics.NodeUsage
    nodeArr, err = getMetrics.GetNodeMetrics(kubeConfig)
    if err != nil {
        log.Fatal(err)
    }

    for _, n1 := range tempNodeResArr {
        for _, n2 := range nodeArr {
            if n1.NodeName == n2.NodeName {
                if n1.Resource == "memory" {
                    str1 := n1.ResourceAvailable[:len(n1.ResourceAvailable)-4]
                    memAllocatable, err := strconv.ParseFloat(str1,64)
                    if err != nil {
                        //fmt.Println(err)
                        log.Fatal(err)
                    }
                    memAllocatable = memAllocatable * 1000000000
                    memUsage := n2.NodeMemUsage * 1024  // Converting to bytes (from Kibibyte)
                    memoryUtilization := ( memUsage / int(memAllocatable)) * 100
                    if memoryUtilization >= 75 {
                        // UPSCALE
                    }else if memoryUtilization <= 20 {
                        // DOWNSCALE
                        fmt.Println("Memory Utilization is BELOW threshold")
                    }else {
                        fmt.Println("Memory Utilization is NETURAL")
                    }

                } else if n1.Resource == "cpu" {
                    cpuAllocatableCores, err := strconv.Atoi(n1.ResourceAvailable)
                    if err != nil {
                        //fmt.Println(err)
                        log.Fatal(err)
                    }
                    // convert CpuUsage UNIT from nanocores to cores by dividing it by 1 billion
                    cpuUsageCores := int(n2.NodeCpuUsage / 1000000000)

                    cpuUtilization := (cpuUsageCores / cpuAllocatableCores) * 100

                    if cpuUtilization >= 75 {
                        // UPSCALE
                    }else if cpuUtilization <= 20 {
                        // DOWNSCALE
                        fmt.Println("CPU Utilization is BELOW threshold")
                    }else {
                        fmt.Println("CPU Utilization is NEUTRAL")
                    }

                }
            }
        }
    }


}


