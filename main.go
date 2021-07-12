package main

import (
    "fmt"
    "time"
    "flag"
    "log"

    ClusterAutoscaler "github.com/ayush5588/ClusterAutoscaler/pkg/metrics"

)
func main() {

    pathPtr := flag.String("path", "", "a string")
    promIpPtr := flag.String("promIP", "", "a string")

    flag.Parse()

    if *pathPtr == "" || *promIpPtr == "" {
        err := "path or promIP empty!"
        log.Fatal(err)
    }


    for i:=0; i<10; i++ {
        fmt.Printf("-------------------CHECK %d-----------------\n\n",i+1)
        ClusterAutoscaler.Start(*promIpPtr, *pathPtr)
        time.Sleep(20*time.Second)
    }
}
