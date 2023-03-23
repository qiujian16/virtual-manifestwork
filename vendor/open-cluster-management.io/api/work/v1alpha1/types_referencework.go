package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	workv1 "open-cluster-management.io/api/work/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
type ReferenceWork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec reperesents the desired ManifestWork payload and Placement reference to be reconciled
	Spec ReferenceWorkSpec `json:"spec,omitempty"`

	// Status represent the current status of Placing ManifestWork resources
	Status ReferenceWorkStatus `json:"status,omitempty"`
}

type ReferenceWorkSpec struct {
	Reference workv1.ResourceIdentifier `json:"workReference,omitempty"`
}

type ReferenceWorkStatus struct {
	workv1.ManifestWorkStatus `json:",inline"`
	ReferenceSpecHash         string `json:"referenceSpecHash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
type ReferenceWorkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReferenceWork `json:"items"`
}
