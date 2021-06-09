package main

import (
    "context"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
    "fmt"
    "k8s.io/client-go/tools/clientcmd"

)

func main() {
    kubeconfig := "/home/ayush5588/go/src/github.com/ClusterAutoscaler/realKubeConfig.conf"
    config, err := clientcmd.BuildConfigFromFlags("",kubeconfig)
    if err != nil {
        panic(err.Error())
    }
    // creates the clientset
    clientset, err := metricsv.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }
    podMetricsList, err := clientset.MetricsV1beta1().PodMetricses("").List(context.TODO(),metav1.ListOptions{})
    //fmt.Printf("Type: %T\n",podMetricsList)
    if err != nil {
        panic(err.Error())
    }
    for _, m := range podMetricsList.Items {
        fmt.Println(m)
    }
}
