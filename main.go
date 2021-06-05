// Code to get the pods in the cluster
package main

import(
    "fmt"
    "os"
    "log"
    list "github.com/ayush5588/ClusterAutoscaler/pkg/podNodeList"
    )


func main() {
         // check whether the list is of Pods or Nodes

         // taking the command line argument for the object type i.e. pod or node
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
         }
}
