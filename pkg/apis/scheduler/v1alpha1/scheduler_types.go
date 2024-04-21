package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
)

type Scheduler struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec SchedulerSpec `json:"spec,omitempty"`
	// +optional
	Status SchedulerStatus `json:"status,omitempty"`
}

var (
	// Check that Scheduler can be validated and defaulted.
	_ apis.Validatable   = (*Scheduler)(nil)
	_ apis.Defaultable   = (*Scheduler)(nil)
	_ kmeta.OwnerRefable = (*Scheduler)(nil)
	// Check that the type conforms to the duck Knative Resource shape.
	_ duckv1.KRShaped = (*Scheduler)(nil)
)

// SchedulerSpec holds the desired state of the Scheduler (from the client).
type SchedulerSpec struct {
	// ServiceName holds the name of the Kubernetes Service to expose as an "addressable".
	ServiceName string `json:"serviceName"`
}

const (
	// SchedulerConditionReady is set when the revision is starting to materialize
	// runtime resources, and becomes true when those resources are ready.
	SchedulerConditionReady = apis.ConditionReady
)

// SchedulerStatus communicates the observed state of the Scheduler (from the controller).
type SchedulerStatus struct {
	duckv1.Status `json:",inline"`

	// Address holds the information needed to connect this Addressable up to receive events.
	// +optional
	Address *duckv1.Addressable `json:"address,omitempty"`
}

// SchedulerList is a list of Scheduler resources
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SchedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Scheduler `json:"items"`
}

// GetStatus retrieves the status of the resource. Implements the KRShaped interface.
func (as *Scheduler) GetStatus() *duckv1.Status {
	return &as.Status.Status
}
