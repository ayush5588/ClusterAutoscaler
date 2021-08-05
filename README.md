# ClusterAutoscaler
Implementation of Kubernetes Cluster Autoscaler in Golang
<br />

## Command to execute the project
```golang

go run main.go -path=pass_kubeConfig_file_path -promIP=Prometheus-Server_IP_with_port
```

## What is Cluster Autoscaler?
The cluster autoscaler is a Kubernetes tool that increases or decreases the size of a Kubernetes cluster (by adding or removing nodes), based on the presence of pending pods and node utilization metrics.

## When Cluster Autoscaler scales up?
Cluster Autoscaler scales up when it sees pending pods which are unscheduled as the resources requested by them can't be allocated because of lack of resources OR there are no amount of resource left of a particular type such as Memory, CPU time, Process IDs, etc. so it scales up the number of nodes and the pending pods gets scheduled on the newly created nodes with the help of scheduler.

## When Cluster Autoscaler scales down?
When the cluster autoscaler sees that certain nodes are underutilized (i.e resources such as CPU, Memory is being utilized lesser than the specified minimum threshold for them), the CA does the pre-removal checks on them. These pre-removal checks involve making sure that the node that has being selected to be removed does not contain pods which have their local storage attached or the pods have PDB (Pod Disruption Budget) defined for them, etc. If the selected nodess pass these pre-removal checks then they are taken down. 

## Architecture Diagram
### 1. ```k8s Cluster Architecture```
![image](https://user-images.githubusercontent.com/48388639/124967507-22214080-e042-11eb-99f5-3efbb879e621.png)
<br />
### 2. ```Upscaling Decision Making Process``` 
![image](https://user-images.githubusercontent.com/48388639/124967694-5694fc80-e042-11eb-84bd-efc051f537ad.png)
<br />
### 3. ```Downscaling Decision Making Process```
![image](https://user-images.githubusercontent.com/48388639/124967994-b8edfd00-e042-11eb-8ce0-b4da09146669.png)


