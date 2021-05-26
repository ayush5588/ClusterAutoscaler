package main

import (
	 "context"
         "k8s.io/client-go/tools/clientcmd"
         "k8s.io/client-go/kubernetes"
         "log"
         "fmt"
         metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    )

func main() {
         kubeconfig :=  "/home/ayush5588/go/src/kubernetes/kubeConfig.conf"
         config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
         if err != nil {
                  log.Fatal(err)
         }
         clientset, err := kubernetes.NewForConfig(config)
         if err != nil {
                  log.Fatal(err)
         }

         nodes, err := clientset.CoreV1().Nodes().List(context.TODO(),metav1.ListOptions{})
         if err != nil {
         	panic(err.Error())
     	 }
     	 fmt.Printf("There are %d nodes in the cluster\n", len(nodes.Items))
     	 for _, node := range nodes.Items {
           	fmt.Printf("Node name=/%s\n",node.GetName())
     	 }
}
