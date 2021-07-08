# ClusterAutoscaler
Implementation of Kubernetes Cluster Autoscaler in Golang
### NOTE
1. Change the kubeConfig file present in the root directory with your cluster config file and also the path of the config file specified in the pkg/metrics/metricsAccumulate.go file

### STEPS
1. To run the project, either first build the project **go build** or do *go run main.go*
<br />

## ```What is Cluster Autoscaler?```
The cluster autoscaler is a Kubernetes tool that increases or decreases the size of a Kubernetes cluster (by adding or removing nodes), based on the presence of pending pods and node utilization metrics.

### ```When Cluster Autoscaler scales up?```
Cluster Autoscaler scales up when it sees pending pods which are unscheduled as the resources requested by them can't be allocated because of lack of resources OR there are no amount of resource left of a particular type such as Memory, CPU time, Process IDs, etc. so it scales up the number of nodes and the pending pods gets scheduled on the newly created nodes with the help of scheduler.

### ```When Cluster Autoscaler scales down?```
When the cluster autoscaler sees that certain nodes are underutilized (i.e resources such as CPU, Memory is being utilized lesser than the specified minimum threshold for them), the CA does the pre-removal checks on them. These pre-removal checks involve making sure that the node that has being selected to be removed does not contain pods which have their local storage attached or the pods have PDB (Pod Disruption Budget) defined for them, etc. If the selected nodess pass these pre-removal checks then they are taken down. 

### ```Architecture Diagram```
#### 1. k8s Cluster Architecture
![image](https://user-images.githubusercontent.com/48388639/124967507-22214080-e042-11eb-99f5-3efbb879e621.png)
<br />
#### 2. Metrics Collection
![image](https://user-images.githubusercontent.com/48388639/114090511-29b33b00-98d5-11eb-8e0e-ead61a5a28bd.png)
<br />
#### 3. Metrics Exposing
![image](https://user-images.githubusercontent.com/48388639/114090716-6a12b900-98d5-11eb-9d4e-e35c8c6bbec7.png)
#### 4. Decision Making Process 
![image](https://user-images.githubusercontent.com/48388639/121774729-62a3b080-cba1-11eb-88f2-b9766d3cc6b0.png)

