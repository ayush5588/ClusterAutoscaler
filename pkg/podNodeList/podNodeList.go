// Code to get the pods in the cluster
package podNodeList

import (
	     "context"
         "k8s.io/client-go/tools/clientcmd"
         "k8s.io/client-go/kubernetes"
         //"log"
         //"fmt"
         metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    )


// get the list of either Pods or Nodes depending upon the itemType
func getItems(string itemType, []string *arr) {
    kubeconfig := "home/ayush5588/src/github.com/ClusterAutoscaler/kubeConfig.conf"
    config, err := clinetcmd.BuildConfigFromFlags("",kubeconfig)
    if err!=nil{
        panic(err.Error())
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }
    if itemType == "pod" {
        arr, err := clientset.CoreV1().Pods("").List(context.TODO(),metav1.ListOptions{})
        if err != nil {
            panic(err.Error())
        }
    }else if itemType == "node" {
        arr, err := clientset.CoreV1().Nodes("").List(context.TODO(),metav1.ListOptions{})
        if err != nil {
            panic(err.Error())
        }
    }
}

