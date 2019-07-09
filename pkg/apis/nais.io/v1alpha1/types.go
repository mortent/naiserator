package v1alpha1

import (
	"fmt"
	"strconv"

	hash "github.com/mitchellh/hashstructure"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	LastSyncedHashAnnotation = "nais.io/lastSyncedHash"
	SecretTypeEnv            = "env"
	SecretTypeFiles          = "files"
	DefaultSecretType        = SecretTypeEnv
	DefaultSecretMountPath   = "/var/run/secrets"
)

// Application defines a NAIS application.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Team",type="string",JSONPath=".metadata.labels.team"
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ApplicationSpec `json:"spec"`
}

// ApplicationSpec contains the NAIS manifest.
type ApplicationSpec struct {
	AccessPolicy    AccessPolicy         `json:"accessPolicy,omitempty"`
	ConfigMaps      ConfigMaps           `json:"configMaps,omitempty"`
	Env             []EnvVar             `json:"env,omitempty"`
	Image           string               `json:"image"`
	Ingresses       []string             `json:"ingresses,omitempty"`
	LeaderElection  bool                 `json:"leaderElection,omitempty"`
	Liveness        Probe                `json:"liveness,omitempty"`
	Logformat       string               `json:"logformat,omitempty"`
	Logtransform    string               `json:"logtransform,omitempty"`
	Port            int                  `json:"port,omitempty"`
	PreStopHookPath string               `json:"preStopHookPath,omitempty"`
	Prometheus      PrometheusConfig     `json:"prometheus,omitempty"`
	Readiness       Probe                `json:"readiness,omitempty"`
	Replicas        Replicas             `json:"replicas,omitempty"`
	Resources       ResourceRequirements `json:"resources,omitempty"`
	Secrets         []Secret             `json:"secrets,omitempty"`
	SecureLogs      SecureLogs           `json:"secureLogs,omitempty"`
	Service         Service              `json:"service,omitempty"`
	SkipCaBundle    bool                 `json:"skipCaBundle,omitempty"`
	Strategy        Strategy             `json:"strategy,omitempty"`
	Vault           Vault                `json:"vault,omitempty"`
	WebProxy        bool                 `json:"webproxy,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Application `json:"items"`
}

type SecureLogs struct {
	Enabled bool `json:"enabled"`
}

type Probe struct {
	Path             string `json:"path"`
	Port             int    `json:"port"`
	InitialDelay     int    `json:"initialDelay"`
	PeriodSeconds    int    `json:"periodSeconds"`
	FailureThreshold int    `json:"failureThreshold"`
	Timeout          int    `json:"timeout"`
}

type PrometheusConfig struct {
	Enabled bool   `json:"enabled"`
	Port    string `json:"port"`
	Path    string `json:"path"`
}

type Replicas struct {
	Min                    int `json:"min"`
	Max                    int `json:"max"`
	CpuThresholdPercentage int `json:"cpuThresholdPercentage"`
}

type ResourceSpec struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}

type ResourceRequirements struct {
	Limits   ResourceSpec `json:"limits"`
	Requests ResourceSpec `json:"requests"`
}

type ObjectFieldSelector struct {
	FieldPath string `json:"fieldPath"`
}

type EnvVarSource struct {
	FieldRef ObjectFieldSelector `json:"fieldRef"`
}

type EnvVar struct {
	Name      string       `json:"name"`
	Value     string       `json:"value"`
	ValueFrom EnvVarSource `json:"valueFrom"`
}

type SecretPath struct {
	MountPath string `json:"mountPath"`
	KvPath    string `json:"kvPath"`
}

type Vault struct {
	Enabled bool         `json:"enabled"`
	Sidecar bool         `json:"sidecar"`
	Mounts  []SecretPath `json:"paths"`
}

type ConfigMaps struct {
	Files []string `json:"files"`
}

type Strategy struct {
	Type string `json:"type"`
}

type Service struct {
	Port int32 `json:"port"`
}

type AccessPolicyExternalRule struct {
	Host string `json:"host"`
}

type AccessPolicyGressRule struct {
	Application string `json:"application"`
	Namespace   string `json:"namespace"`
}

type AccessPolicyInbound struct {
	Rules []AccessPolicyGressRule `json:"rules"`
}

type AccessPolicyOutbound struct {
	Rules    []AccessPolicyGressRule    `json:"rules"`
	External []AccessPolicyExternalRule `json:"external"`
}

type AccessPolicy struct {
	Inbound  AccessPolicyInbound  `json:"inbound"`
	Outbound AccessPolicyOutbound `json:"outbound"`
}

type Secret struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	MountPath string `json:"mountPath"`
}

func (in *Application) GetObjectKind() schema.ObjectKind {
	return in
}

func (in *Application) GetObjectReference() v1.ObjectReference {
	return v1.ObjectReference{
		APIVersion:      "v1alpha1",
		UID:             in.UID,
		Name:            in.Name,
		Kind:            "Application",
		ResourceVersion: in.ResourceVersion,
		Namespace:       in.Namespace,
	}
}

func (in *Application) GetOwnerReference() metav1.OwnerReference {
	return metav1.OwnerReference{
		APIVersion: "v1alpha1",
		Kind:       "Application",
		Name:       in.Name,
		UID:        in.UID,
	}
}

// NilFix initializes all slices from their nil defaults.
//
// This is done in order to workaround the k8s client serializer
// which crashes when these fields are uninitialized.
func (in *Application) NilFix() {
	if in.Spec.Ingresses == nil {
		in.Spec.Ingresses = make([]string, 0)
	}
	if in.Spec.Env == nil {
		in.Spec.Env = make([]EnvVar, 0)
	}
	if in.Spec.Secrets == nil {
		in.Spec.Secrets = make([]Secret, 0)
	}
	if in.Spec.Vault.Mounts == nil {
		in.Spec.Vault.Mounts = make([]SecretPath, 0)
	}
	if in.Spec.ConfigMaps.Files == nil {
		in.Spec.ConfigMaps.Files = make([]string, 0)
	}
	if in.Spec.AccessPolicy.Inbound.Rules == nil {
		in.Spec.AccessPolicy.Inbound.Rules = make([]AccessPolicyGressRule, 0)
	}
	if in.Spec.AccessPolicy.Outbound.Rules == nil {
		in.Spec.AccessPolicy.Outbound.Rules = make([]AccessPolicyGressRule, 0)
	}
}

func (in Application) Hash() (string, error) {
	// struct including the relevant fields for
	// creating a hash of an Application object
	relevantValues := struct {
		AppSpec     ApplicationSpec
		Labels      map[string]string
		ChangeCause string
	}{
		in.Spec,
		in.Labels,
		in.Annotations["kubernetes.io/change-cause"],
	}

	h, err := hash.Hash(relevantValues, nil)
	return strconv.FormatUint(h, 10), err
}

func (in *Application) LastSyncedHash() string {
	a := in.GetAnnotations()
	if a == nil {
		a = make(map[string]string)
	}
	return a[LastSyncedHashAnnotation]
}

func (in *Application) SetLastSyncedHash(hash string) {
	a := in.GetAnnotations()
	if a == nil {
		a = make(map[string]string)
	}
	a[LastSyncedHashAnnotation] = hash
	in.SetAnnotations(a)
}

func (in *Application) DefaultSecretPath(base string) SecretPath {
	return SecretPath{
		MountPath: DefaultVaultMountPath,
		KvPath:    fmt.Sprintf("%s/%s/%s", base, in.Name, in.Namespace),
	}
}