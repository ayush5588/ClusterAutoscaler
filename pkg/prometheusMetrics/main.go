/* This is under development prometheus metrics collection file */

package main

import (
   "io/ioutil"
   "encoding/json"
   "log"
   "net/http"
//   "fmt"

   metricsStruct "github.com/ayush5588/ClusterAutoscaler/metricsStruct"
)


type TempPodStatusStruct struct {
    PodName string
    PodPhase string
    PodPhaseValue string
}

type TempNodeStatusStruct struct {
    NodeName string
    Condition string
    NodeStatus string
    NodeStatusValue string
}

type TempPodsNotScheduledStruct struct {
    Namespace string
    PodName string
}



func PodStatusPhase (promServerIP string) ([]TempPodStatusStruct, error) {

    query := "kube_pod_status"
    resp, err := http.Get(promServerIP+query)
    if err != nil {
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var pod metricsStruct.PodStatusPhaseStruct
    err = json.Unmarshal(body,&pod)
    if err != nil {
        return nil, err
    }

    var PodStatusPhaseArr []TempPodStatusStruct

    for _, p := range pod.Data.Result {
        var tempPod TempPodStatusStruct
        tempPod.PodName = p.Metric.Pod
        tempPod.PodPhase = p.Metric.Phase
        tempPod.PodPhaseValue = p.Value[1]
        PodStatusPhaseArr = append(PodStatusPhaseArr,tempPod)
    }

    return PodStatusPhaseArr, nil
}



func NodeStatusPhase (promServerIP string) ([]TempNodeStatusStruct, error) {

    query := "kube_node_status_condition"
    resp, err := http.get(promServerIP+query)
    if err != nil {
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.body)
    if err != nil {
        return nil, err
    }

    var node metricsStruct.NodeStatusStruct
    err = json.Unmarshal(body,&node)
    if err != nil {
        return nil, err
    }

    var NodeStatusPhaseArr []TempNodeStatusStruct

    for _, n := range node.Data.Result {
        var tempNode TempNodeStatusStruct
        tempNode.NodeName = n.Metric.Node
        tempNode.Condition = n.Metric.Condition
        tempNode.NodeStatus = n.Metric.Status
        tempNode.NodeStatusValue = n.Value[1]
        NodeStatusPhaseArr = append(NodeStatusPhaseArr, tempNode)
    }

    return NodeStatusPhaseArr, nil
}


func PodsNotScheduled (promServerIP string) ([]TempPodsNotScheduledStruct, error) {

    query := "kube_pod_status_unschedulable"
    resp, err := http.Get(promServerIP+query)
    if err != nil {
        return nil, err
    }

    body, err := ioutil.ReadAll(resp)
    if err != nil {
        return nil, err
    }

    var pod metricsStruct.PodStatusUnschedulableStruct
    err := json.Unmarshal(body,&pod)
    if err != nil {
        return nil, err
    }

    var PodsNotScheduledArr []TempPodsNotScheduledStruct
    if pod.Data.Result == nil {
        return nil, nil
    }
   // Will add the part for the case where there are more than 1 unscheduled pods
   var tempPod TempPodsNotScheduledStruct
   tempPod.Namespace = pod[0]
   tempPod.PodName = pod[1]
   PodsNotScheduledArr = append(PodsNotScheduledArr, tempPod)

   return PodsNotScheduledArr, nil
}

/*
func main() {
   //10.103.151.144 is the ClusterIP address of the prometheus-server service 
   resp, err := http.Get("http://10.101.202.25:80/api/v1/query?query=kube_node_status_condition")
   // metrics.NodeInfoStruct is a structure for kube_node_info which is defined in the metricsStruct.go of metrics package
   //var node metricsStruct.NodeInfoStruct
   var node metricsStruct.NodeStatusStruct
   if err != nil {
      log.Fatalln(err)
   }
//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
   err = json.Unmarshal(body,&node)
   //fmt.Println(pod)
   if err != nil {
        log.Fatalln(err)
   }
//   fmt.Println(node.Data.Result.Value)
   for _, d := range node.Data.Result {
       fmt.Printf("Node: %s\nCondition: %s\nStatus: %s\nValue: %s",d.Metric.Node,d.Metric.Condition,d.Metric.Status,d.Value[1])
       fmt.Println("\n")
   }

}
*/
