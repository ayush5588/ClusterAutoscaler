/* This is under development prometheus metrics collection file */

package promMetrics
//package main

import (
   "io/ioutil"
   "encoding/json"
   //"log"
   "net/http"
   "fmt"

   metricsStruct "github.com/ayush5588/ClusterAutoscaler/metricsStruct"
)


type TempPodStatusStruct struct {
    PodName string
    PodPhase string
    PhaseValue string
}

type TempNodeStatusStruct struct {
    NodeName string
    Condition string
    ConditionStatus string
    ConditionStatusValue string
}

type TempPodsNotScheduledStruct struct {
    Namespace string
    PodName string
}

func checkType (value []interface{}) (string) {
    flag := false
    var ans string
    for _, v := range value {
        if v == nil {
            ans = ""
        } else {
            switch v.(type) {
            case int:
                //fmt.Println("it is a int")
            case string:
                //fmt.Println("it is a string")
                flag = true
            default:
                //fmt.Println("i don't know it")
            }
        }
        if flag == true {
            ans = v.(string)
            break
        }
    }
    if flag == false {
        ans = ""
    }
    return ans
}

func checkType2(value []interface{}) (string, string) {
    flag := false
    cnt := 1
    var str1, str2 string
    for _, v := range value {
        switch v.(type) {
            case string:
                flag = true
            default:
                flag = false
        }
        if flag == true{
            if cnt == 1 {
                str1 = v.(string)
                cnt = cnt + 1
            }else {
                str2 = v.(string)
                break
            }
        }else {
            str1 = ""
            str2 = ""
        }
    }
    return str1, str2
}


func PodStatusPhase (promServerIP string) ([]TempPodStatusStruct, error) {

    var PodStatusPhaseArr []TempPodStatusStruct

    query := "kube_pod_status_phase"
    fmt.Println(promServerIP+query)
    resp, err := http.Get(promServerIP+query)
    if err != nil {
        return PodStatusPhaseArr, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return PodStatusPhaseArr, err
    }

    var pod metricsStruct.PodStatusPhaseStruct
    err = json.Unmarshal(body,&pod)
    if err != nil {
        return PodStatusPhaseArr, err
    }

    //var PodStatusPhaseArr []TempPodStatusStruct

    for _, p := range pod.Data.Result {
        var tempPod TempPodStatusStruct
        tempPod.PodName = p.Metric.Pod
        tempPod.PodPhase = p.Metric.Phase        
        str := checkType(p.Value)
        tempPod.PhaseValue = str

        PodStatusPhaseArr = append(PodStatusPhaseArr,tempPod)
    }

    return PodStatusPhaseArr, nil
}





func NodeStatusPhase (promServerIP string) ([]TempNodeStatusStruct, error) {
    
    var NodeStatusPhaseArr []TempNodeStatusStruct

    query := "kube_node_status_condition"
    resp, err := http.Get(promServerIP+query)
    if err != nil {
        return NodeStatusPhaseArr, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return NodeStatusPhaseArr, err
    }

    var node metricsStruct.NodeStatusStruct
    err = json.Unmarshal(body,&node)
    if err != nil {
        return NodeStatusPhaseArr, err
    }

    //var NodeStatusPhaseArr []TempNodeStatusStruct

    for _, n := range node.Data.Result {
        var tempNode TempNodeStatusStruct
        tempNode.NodeName = n.Metric.Node
        tempNode.Condition = n.Metric.Condition
        tempNode.ConditionStatus = n.Metric.Status
        str := checkType(n.Value)
        tempNode.ConditionStatusValue = str
        NodeStatusPhaseArr = append(NodeStatusPhaseArr, tempNode)
    }

    return NodeStatusPhaseArr, nil
}




func PodsNotScheduled (promServerIP string) ([]TempPodsNotScheduledStruct, error) {

    var PodsNotScheduledArr []TempPodsNotScheduledStruct

    query := "kube_pod_status_unschedulable"
    resp, err := http.Get(promServerIP+query)
    if err != nil {
        return PodsNotScheduledArr, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return PodsNotScheduledArr, err
    }

    var pod metricsStruct.PodStatusUnschedulableStruct
    err = json.Unmarshal(body,&pod)
    if err != nil {
        return PodsNotScheduledArr, err
    }

    //var PodsNotScheduledArr []TempPodsNotScheduledStruct
    if pod.Data.Result == nil {
        return PodsNotScheduledArr, nil
    }
   // Will add the part for the case where there are more than 1 unscheduled pods
   var tempPod TempPodsNotScheduledStruct
   str1, str2 := checkType2(pod.Data.Result)
   tempPod.Namespace = str1
   tempPod.PodName = str2
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
       fmt.Printf("%T: ", d.Value[0])
//       fmt.Printf("Node: %s\nCondition: %s\nStatus: %s\nValue: %s\nType: %T",d.Metric.Node,d.Metric.Condition,d.Metric.Status,d.Value[1],d.Value[1])
       fmt.Println("\n")
   }

}
*/
