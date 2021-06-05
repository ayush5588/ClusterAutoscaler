# ClusterAutoscaler
Implementation of Kubernetes Cluster Autoscaler in Golang
### NOTE
1. Change the kubeConfig file with your cluster config file

### STEPS
1. To get the name and number of pods or nodes in the cluster, first build the project **go build**.
2. Then give the object name i.e **./ClusterAutoscaler <object name>** where object name can be **pod** or **node**
