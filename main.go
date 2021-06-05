// Code to get the pods in the cluster
package main

import(
    "fmt"
    "log"
    list "github.com/ayush5588/ClusterAutoscaler/podNodeList"
    )


func main() {
         // check whether the list is of Pods or Nodes
         var arr []string
         err := list.getItem("Pod",&arr)
         if err != nil {
            log.Fatal(err)
         }
         listType := "Pods"
         fmt.Printf("Resource type: %s", listType)
         for _,element := range arr.items {
            fmt.Printf("%s name = %s\n",listType,element)
         }
}
