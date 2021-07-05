package metrics

import (
   "io/ioutil"
   "encoding/json"
   //"log"
   "net/http"
   //"fmt"

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

type TempNodeResourceStruct struct {
    NodeName string
    Resource string
    ResourceUnit string
    ResourceAvailable string
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
    //fmt.Println(promServerIP+query)
    resp, err := http.Get(promServerIP+query)
    if err != nil {
        return PodStatusPhaseArr, err
        //log.Fatal(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return PodStatusPhaseArr, err
        //log.Fatal(err)
    }

    var pod metricsStruct.PodStatusPhaseStruct
    err = json.Unmarshal(body,&pod)
    if err != nil {
        return PodStatusPhaseArr, err
        //log.Fatal(err)
    }


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




// []TempNodeStatusStruct, error
func NodeStatusPhase (promServerIP string) ([]TempNodeStatusStruct, error) {
    
    var NodeStatusPhaseArr []TempNodeStatusStruct

    query := "kube_node_status_condition"
    resp, err := http.Get(promServerIP+query)
    if err != nil {
        return NodeStatusPhaseArr, err
        //log.Fatal(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return NodeStatusPhaseArr, err
        //log.Fatal(err)
    }

    
    var node metricsStruct.NodeStatusStruct
    err = json.Unmarshal(body,&node)
    if err != nil {
        return NodeStatusPhaseArr, err
        //log.Fatal(err)
        
    }


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
        //log.Fatal(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return PodsNotScheduledArr, err
        //log.Fatal(err)
    }

    var pod metricsStruct.PodStatusUnschedulableStruct
    err = json.Unmarshal(body,&pod)
    if err != nil {
        return PodsNotScheduledArr, err
        //log.Fatal(err)
    }


    if pod.Data.Result == nil {
        return PodsNotScheduledArr, nil
        //log.Fatal(err)
    }

    for _, p:= range pod.Data.Result {
       var tempPod TempPodsNotScheduledStruct
       tempPod.Namespace = p.Metric.Namespace
       tempPod.PodName = p.Metric.Pod
       PodsNotScheduledArr = append(PodsNotScheduledArr, tempPod)
   }
    
   /*
   var tempPod TempPodsNotScheduledStruct
   str1, str2 := checkType2(pod.Data.Result)
   tempPod.Namespace = str1
   tempPod.PodName = str2
   PodsNotScheduledArr = append(PodsNotScheduledArr, tempPod)
   */

   return PodsNotScheduledArr, nil
}



func NodeAllocatableResources (promServerIP string) ([]TempNodeResourceStruct, error){

    var tempNodeResArr []TempNodeResourceStruct

    query := "kube_node_status_allocatable"
    resp, err := http.Get(promServerIP+query)
    if err != nil {
        return tempNodeResArr, err
        //log.Fatal(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return tempNodeResArr, err
        //log.Fatal(err)
    }


    var node metricsStruct.NodeResAllocatableStruct
    err = json.Unmarshal(body, &node)
    if err != nil {
        return tempNodeResArr, err
        //log.Fatal(err)
    }

    for _, n := range node.Data.Result {
        var tempNode TempNodeResourceStruct
        tempNode.NodeName = n.Metric.Node
        tempNode.Resource = n.Metric.Resource
        tempNode.ResourceUnit = n.Metric.Unit
        str := checkType(n.Value)
        tempNode.ResourceAvailable = str
        tempNodeResArr = append(tempNodeResArr, tempNode)
    }

    return tempNodeResArr, nil
}


/*
func PodInNodes(promServerIP string) (map[string][]string, error) {

    nodesMap := make(map[string][]string)
    query := "kube_pod_info"
    resp, err := http.Get(promServerIP+query)
    if err != nil {
        //log.Fatal(err)
        return nodesMap, err
    }
    body, err := ioutil.ReadAll(resp.Body)

    var nodePod metricsStruct.PodInfoStruct

    err = json.Unmarshal(body, &nodePod)

    if err != nil {
        return nodesMap, err

    }

    for _, n := range nodePod.Data.Result {
        nodesMap[n.Metric.Node] = append(nodesMap[n.Metric.Node], n.Metric.Pod)
    }

    return nodesMap, nil

}

*/



/*
func AllFunc(promServerIP string) (*bigcache.BigCache) {

    //PodStatusPhase (promServerIP)
    NodeStatusPhase (promServerIP)
    PodsNotScheduled (promServerIP)
    NodeAllocatableResources (promServerIP)

    return Cache
}
*/
