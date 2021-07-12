package main

import (
    "fmt"
    "time"
    ClusterAutoscaler "github.com/ayush5588/ClusterAutoscaler/pkg/metrics"

)
func main() {

    for i:=0; i<10; i++ {
        fmt.Printf("-------------------CHECK %d-----------------\n\n",i+1)
        ClusterAutoscaler.Start()
        time.Sleep(20*time.Second)
    }
}
