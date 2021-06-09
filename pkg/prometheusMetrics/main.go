/* This is under development prometheus metrics collection file */

package main

import (
   "io/ioutil"
   "encoding/json"
   "log"
   "net/http"
   "fmt"

   metrics "github.com/ayush5588/ClusterAutoscaler/pkg/metrics"
)


func main() {
   /* 10.103.151.144 is the ClusterIP address of the prometheus-server service */
   resp, err := http.Get("http://10.103.151.144:80/api/v1/query?query=kube_node_info")
   /* metrics.NodeInfoStruct is a structure for kube_node_info which is defined in the metricsStruct.go of metrics package */
   var node metrics.NodeInfoStruct
   if err != nil {
      log.Fatalln(err)
   }
//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }

   err = json.Unmarshal(body,&node)
   if err != nil {
        log.Fatalln(err)
   }
   for _, d := range node.Data.Result {
        fmt.Println(d.Metric.Node)
        fmt.Println(d.Metric.InternalIP)
   }

}
