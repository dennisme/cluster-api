package api

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Cluster struct {
	metav1.ObjectMeta
	Spec   ClusterSpec `json:"spec"`
	//Status ClusterStatus
}


type ClusterSpec struct {
	Cloud        string `json:"cloud"` // e.g. aws, azure, gcp
	Project      string `json:"project"`
	SSH          *SSHConfig `json:"ssh"`
}

type SSHConfig struct {
	PublicKeyPath string `json:"publicKeyPath"`
	User      string `json:"user"`
}

type ControlPlaneConfig struct {
	APIServer         APIServerConfig
	ControllerManager ControllerManagerConfig
}

type APIServerConfig struct {
	AdvertiseAddress  string // e.g. "0.0.0.0"
	Port              uint32
	CertExtraSANs     string
	AuthorizationMode string
	KubernetesVersion KubernetesVersionInfo
}

type ControllerManagerConfig struct {
	LeaderElection    bool
	Port              uint32
	KubernetesVersion KubernetesVersionInfo
}

type EtcdConfig struct {
	Version string // full semver
	API     string // e.g. v2, v3
}

type KubernetesVersionInfo struct {
	SemVer   string
	Location string // e.g. https://storageapis.google.com/...
}
