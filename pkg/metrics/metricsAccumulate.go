package metricsAccumulate

import (

    "log"

    metricsStruct "github.com/ayush5588/ClusterAutoscaler/metricsStruct"
    metrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics/prometheusMetrics"
)

var promServerIP string = "http://10.101.202.25:80/api/v1/query?query="


func accumulate() (metricsStruct.Metrics, error) {
    var FinalMetricsStruct metricsStruct.Metrics
    
    FinalMetricsStruct.UnscheduledPodsCNT = 0
    FinalMetricsStruct.CpuUtilizationCNT = 0
    FinalMetricsStruct.MemUtilizationCNT = 0
    FinalMetricsStruct.PIDPressureStatus = false
    FinalMetricsStruct.DISKPressureStatus = false
    FinalMetricsStruct.MEMPressureStatus = false
    FinalMetricsStruct.StructSET = false

    // Checking for the Node status 
    var tempNodeStatusArr []metrics.TempNodeStatusStruct
    tempNodeStatusArr, err := metrics.NodeStatusPhase(promServerIP)
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
                FinalMetricsStruct.DISKPressure = true
                FinalMetricsStruct.StructSET = true
            }
        }
    }

    // Checking for the Unscheduled pods 
    var tempPodsNotScheduledArr []metrics.TempPodsNotScheduledStruct
    tempPodsNotScheduledArr, err = metrics.PodsNotScheduled(promServerIP)
    if err != nil {
        log.Fatal(err)
        //FinalMetricsStruct.StructSET = false
        //return FinalMetricsStruct, err
    }
    if len(tempPodsNotScheduledArr) > 0 {
        FinalMetricsStruct.UnscheduledPodsCNT = 1;
        FinalMetricsStruct.StructSET = true
    }


    //  Calculate the current utilization for the CPU and MEM

    /*  Get the metrics from the Metrics-Server and the resources existing to be allocated from Prometheus 
        (will have to add logic in the promMetrics.go) and calculate the percentage to get the Current Utilization 
        of the resources. 
    */



}


