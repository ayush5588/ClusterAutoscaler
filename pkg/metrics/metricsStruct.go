// Struct for kube_node_info query results which gives information about nodes
package metrics

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


/* Struct for kube_node_status_capacity query results which tells capacity of different resources of a node */
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


/* Struct for kube_node_status_allocatable query results which tells how much of a resource(which is schedulable) on a node is allocatable */
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

/*Struct for kube_pod_status_unschedulable query results which describes the unschedulable status for the pod*/
type PodStatusUnschedulableStruct struct {
	Status string `json:"status"`
	Data   struct {
		Resulttype string        `json:"resultType"`
		Result     []interface{} `json:"result"`
	} `json:"data"`
}

/* Struct for kube_pod_status_phase query results which tells the pod current phase */
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

