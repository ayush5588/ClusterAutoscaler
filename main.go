// Code to get the pods in the cluster
package main

import(
    "fmt"
    //"os"
    "log"
    //list "github.com/ayush5588/ClusterAutoscaler/pkg/podNodeList"
    GetMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/objectMetrics"
    )


func main() {
         // check whether the list is of Pods or Nodes

         // taking the command line argument for the object type i.e. pod or node
         /*
         itemType := os.Args[1]
         arr := []string{}
         arr, err := list.GetItems(itemType)
         if err != nil {
            log.Fatal(err)
            }
         fmt.Printf("size: %d\n",len(arr))
         fmt.Printf("Resource type: %s\n", itemType)
         for _,element := range arr {
            fmt.Printf("%s name = %s\n",itemType,element)
         }*/
         var arr []GetMetrics.PodUsage
         arr, err := GetMetrics.GetObjectMetrics()
         if err != nil {
            log.Fatal(err)
         }
         //fmt.Println(arr)
         for _, p := range arr {
            fmt.Printf("Name: %s\nCPU Usage: %dn\nMemory Usage: %dki\n\n",p.PodName,p.PodCpuUsage,p.PodMemUsage)
            //fmt.Printf("Type: %T\n",p)
            //fmt.Println(p.Name)
         }

}
