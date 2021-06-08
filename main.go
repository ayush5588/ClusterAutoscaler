package main

import(
    "fmt"
    //"os"
    "log"
    //list "github.com/ayush5588/ClusterAutoscaler/pkg/podNodeList"
    GetMetrics "github.com/ayush5588/ClusterAutoscaler/pkg/objectMetrics"
    )


func main() {

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
         var arr []GetMetrics.NodeUsage
         arr, err := GetMetrics.GetNodeMetrics()
         if err != nil {
            log.Fatal(err)
         }
         for _, p := range arr {
            fmt.Printf("Name: %s\nCPU Usage: %dn\nMemory Usage: %dki\n\n",p.NodeName,p.NodeCpuUsage,p.NodeMemUsage)
         }

}
