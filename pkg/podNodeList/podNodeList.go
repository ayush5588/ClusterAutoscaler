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
// GetItems: INPUT -> type of object (i.e. pod or node)
// GetItems: OUTPUT -> Slice containing the names in all namespace of the requested object and error (if exist)
func GetItems(itemType string, kubeconfig string) ([]string, error) {
    resourceList := []string{}
    //kubeconfig := "/home/ayush5588/go/src/github.com/ClusterAutoscaler/realKubeConfig.conf"
    config, err := clientcmd.BuildConfigFromFlags("",kubeconfig)
    if err!=nil{
        return resourceList, err
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return resourceList, err
    }
    if itemType == "pod" {
        arr, err := clientset.CoreV1().Pods("").List(context.TODO(),metav1.ListOptions{})
        if err != nil {
            return resourceList, err
        }
        // Insert the Pods name in the array resourceList
        for _, item := range arr.Items {
            resourceList = append(resourceList,item.GetName())
        }
    }else if itemType == "node" {
        arr, err := clientset.CoreV1().Nodes().List(context.TODO(),metav1.ListOptions{})
        if err != nil {
            return resourceList, err
        }
        // Insert the Nodes name in the array resourceList
        for _, item := range arr.Items {
            resourceList = append(resourceList,item.GetName())
        }
    }

    return resourceList, nil
}

