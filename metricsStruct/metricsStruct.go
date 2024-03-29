package metricsStruct

import (
    "time"
)


type Metrics struct {
    UnscheduledPodsCNT int
    CpuUtilizationCNT int
    MemUtilizationCNT int
    PIDPressureStatus bool
    DISKPressureStatus bool
    MEMPressureStatus bool
    StructSET bool
}



// Struct for apis/metrics.k8s.io/v1beta1/pods query results which tells pod resource metrics usage from Metrics-Server
type PodMetricsStruct struct {
	Kind       string `json:"kind"`
	Apiversion string `json:"apiVersion"`
	Metadata   struct {
		Selflink string `json:"selfLink"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			Selflink          string    `json:"selfLink"`
			Creationtimestamp time.Time `json:"creationTimestamp"`
		} `json:"metadata"`
		Timestamp  time.Time `json:"timestamp"`
		Window     string    `json:"window"`
		Containers []struct {
			Name  string `json:"name"`
			Usage struct {
				CPU    string `json:"cpu"`
				Memory string `json:"memory"`
			} `json:"usage"`
		} `json:"containers"`
	} `json:"items"`
}



// Struct for apis/metrics.k8s.io/v1beta1/nodes query results which tells nodes resource metrucs usage from Metrics-Server
type NodeMetricsStruct struct {
	Kind       string `json:"kind"`
	Apiversion string `json:"apiVersion"`
	Metadata   struct {
		Selflink string `json:"selfLink"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Selflink          string    `json:"selfLink"`
			Creationtimestamp time.Time `json:"creationTimestamp"`
		} `json:"metadata"`
		Timestamp time.Time `json:"timestamp"`
		Window    string    `json:"window"`
		Usage     struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"usage"`
	} `json:"items"`
}



// Struct for kube_pod_info query results which tells about the pods present in the cluster
type PodInfoStruct struct {
        Status string `json:"status"`
        Data   struct {
                Resulttype string `json:"resultType"`
                Result     []struct {
                        Metric struct {
                                Name                     string `json:"__name__"`
                                AppKubernetesIoInstance  string `json:"app_kubernetes_io_instance"`
                                AppKubernetesIoManagedBy string `json:"app_kubernetes_io_managed_by"`
                                AppKubernetesIoName      string `json:"app_kubernetes_io_name"`
                                CreatedByKind            string `json:"created_by_kind"`
                                CreatedByName            string `json:"created_by_name"`
                                HelmShChart              string `json:"helm_sh_chart"`
                                HostIP                   string `json:"host_ip"`
                                Instance                 string `json:"instance"`
                                Job                      string `json:"job"`
                                KubernetesName           string `json:"kubernetes_name"`
                                KubernetesNamespace      string `json:"kubernetes_namespace"`
                                KubernetesNode           string `json:"kubernetes_node"`
                                Namespace                string `json:"namespace"`
                                Node                     string `json:"node"`
                                Pod                      string `json:"pod"`
                                PodIP                    string `json:"pod_ip"`
                                PriorityClass            string `json:"priority_class"`
                                UID                      string `json:"uid"`
                        } `json:"metric"`
                        Value []interface{} `json:"value"`
                } `json:"result"`
        } `json:"data"`
}



// Struct for kube_node_info query results which tells about the nodes present in the cluster
type NodeInfoStruct struct {
	Status string `json:"status"`
	Data   struct {
		Resulttype string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name                     string `json:"__name__"`
				AppKubernetesIoInstance  string `json:"app_kubernetes_io_instance"`
				AppKubernetesIoManagedBy string `json:"app_kubernetes_io_managed_by"`
				AppKubernetesIoName      string `json:"app_kubernetes_io_name"`
				ContainerRuntimeVersion  string `json:"container_runtime_version"`
				HelmShChart              string `json:"helm_sh_chart"`
				Instance                 string `json:"instance"`
				InternalIP               string `json:"internal_ip"`
				Job                      string `json:"job"`
				KernelVersion            string `json:"kernel_version"`
				KubeletVersion           string `json:"kubelet_version"`
				KubeproxyVersion         string `json:"kubeproxy_version"`
				KubernetesName           string `json:"kubernetes_name"`
				KubernetesNamespace      string `json:"kubernetes_namespace"`
				KubernetesNode           string `json:"kubernetes_node"`
				Node                     string `json:"node"`
				OsImage                  string `json:"os_image"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}



// Struct for kube_node_spec_unschedulable query results which tells whether a node can schedule new pods
type NodeSpecUnschedStruct struct {
	Status string `json:"status"`
	Data   struct {
		Resulttype string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name                     string `json:"__name__"`
				AppKubernetesIoInstance  string `json:"app_kubernetes_io_instance"`
				AppKubernetesIoManagedBy string `json:"app_kubernetes_io_managed_by"`
				AppKubernetesIoName      string `json:"app_kubernetes_io_name"`
				HelmShChart              string `json:"helm_sh_chart"`
				Instance                 string `json:"instance"`
				Job                      string `json:"job"`
				KubernetesName           string `json:"kubernetes_name"`
				KubernetesNamespace      string `json:"kubernetes_namespace"`
				KubernetesNode           string `json:"kubernetes_node"`
				Node                     string `json:"node"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}



// Struct for kube_node_status_condition query results which tells the status of the nodes
type NodeStatusStruct struct {
	Status string `json:"status"`
	Data   struct {
		Resulttype string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name                     string `json:"__name__"`
				AppKubernetesIoInstance  string `json:"app_kubernetes_io_instance"`
				AppKubernetesIoManagedBy string `json:"app_kubernetes_io_managed_by"`
				AppKubernetesIoName      string `json:"app_kubernetes_io_name"`
				Condition                string `json:"condition"`
				HelmShChart              string `json:"helm_sh_chart"`
				Instance                 string `json:"instance"`
				Job                      string `json:"job"`
				KubernetesName           string `json:"kubernetes_name"`
				KubernetesNamespace      string `json:"kubernetes_namespace"`
				KubernetesNode           string `json:"kubernetes_node"`
				Node                     string `json:"node"`
				Status                   string `json:"status"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}



// Struct for kube_node_status_capacity query results which tells capacity of different resources of a node
type NodeResCapacityStruct struct {
	Status string `json:"status"`
	Data   struct {
		Resulttype string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name                     string `json:"__name__"`
				AppKubernetesIoInstance  string `json:"app_kubernetes_io_instance"`
				AppKubernetesIoManagedBy string `json:"app_kubernetes_io_managed_by"`
				AppKubernetesIoName      string `json:"app_kubernetes_io_name"`
				HelmShChart              string `json:"helm_sh_chart"`
				Instance                 string `json:"instance"`
				Job                      string `json:"job"`
				KubernetesName           string `json:"kubernetes_name"`
				KubernetesNamespace      string `json:"kubernetes_namespace"`
				KubernetesNode           string `json:"kubernetes_node"`
				Node                     string `json:"node"`
				Resource                 string `json:"resource"`
				Unit                     string `json:"unit"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}



// Struct for kube_node_status_allocatable query results which tells how much of a resource(which is schedulable) on a node is allocatable
type NodeResAllocatableStruct struct {
	Status string `json:"status"`
	Data   struct {
		Resulttype string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name                     string `json:"__name__"`
				AppKubernetesIoInstance  string `json:"app_kubernetes_io_instance"`
				AppKubernetesIoManagedBy string `json:"app_kubernetes_io_managed_by"`
				AppKubernetesIoName      string `json:"app_kubernetes_io_name"`
				HelmShChart              string `json:"helm_sh_chart"`
				Instance                 string `json:"instance"`
				Job                      string `json:"job"`
				KubernetesName           string `json:"kubernetes_name"`
				KubernetesNamespace      string `json:"kubernetes_namespace"`
				KubernetesNode           string `json:"kubernetes_node"`
				Node                     string `json:"node"`
				Resource                 string `json:"resource"`
				Unit                     string `json:"unit"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}



// Struct for kube_pod_status_unschedulable query results which describes the unschedulable status for the pod

/*
type PodStatusUnschedulableStruct struct {
	Status string `json:"status"`
	Data   struct {
		Resulttype string        `json:"resultType"`
		Result     []interface{} `json:"result"`
	} `json:"data"`
}
*/

type PodStatusUnschedulableStruct struct {
	Status string `json:"status"`
	Data   struct {
		Resulttype string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name                     string `json:"__name__"`
				AppKubernetesIoInstance  string `json:"app_kubernetes_io_instance"`
				AppKubernetesIoManagedBy string `json:"app_kubernetes_io_managed_by"`
				AppKubernetesIoName      string `json:"app_kubernetes_io_name"`
				HelmShChart              string `json:"helm_sh_chart"`
				Instance                 string `json:"instance"`
				Job                      string `json:"job"`
				KubernetesName           string `json:"kubernetes_name"`
				KubernetesNamespace      string `json:"kubernetes_namespace"`
				KubernetesNode           string `json:"kubernetes_node"`
				Namespace                string `json:"namespace"`
				Pod                      string `json:"pod"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}



// Struct for kube_pod_status_phase query results which tells the pod current phase
type PodStatusPhaseStruct struct {
	Status string `json:"status"`
	Data   struct {
		Resulttype string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name                     string `json:"__name__"`
				AppKubernetesIoInstance  string `json:"app_kubernetes_io_instance"`
				AppKubernetesIoManagedBy string `json:"app_kubernetes_io_managed_by"`
				AppKubernetesIoName      string `json:"app_kubernetes_io_name"`
				HelmShChart              string `json:"helm_sh_chart"`
				Instance                 string `json:"instance"`
				Job                      string `json:"job"`
				KubernetesName           string `json:"kubernetes_name"`
				KubernetesNamespace      string `json:"kubernetes_namespace"`
				KubernetesNode           string `json:"kubernetes_node"`
				Namespace                string `json:"namespace"`
				Phase                    string `json:"phase"`
				Pod                      string `json:"pod"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}



// Struct for kube_pod_container_resource_requests query results which tells about requested resource amount for each node
type NodeResRequestStruct struct {
	Status string `json:"status"`
	Data   struct {
		Resulttype string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name                     string `json:"__name__"`
				AppKubernetesIoInstance  string `json:"app_kubernetes_io_instance"`
				AppKubernetesIoManagedBy string `json:"app_kubernetes_io_managed_by"`
				AppKubernetesIoName      string `json:"app_kubernetes_io_name"`
				Container                string `json:"container"`
				HelmShChart              string `json:"helm_sh_chart"`
				Instance                 string `json:"instance"`
				Job                      string `json:"job"`
				KubernetesName           string `json:"kubernetes_name"`
				KubernetesNamespace      string `json:"kubernetes_namespace"`
				KubernetesNode           string `json:"kubernetes_node"`
				Namespace                string `json:"namespace"`
				Node                     string `json:"node"`
				Pod                      string `json:"pod"`
				Resource                 string `json:"resource"`
				Unit                     string `json:"unit"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

