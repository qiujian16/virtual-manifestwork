package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	workapiv1 "open-cluster-management.io/api/work/v1"
)

const GroupName = "virtual.open-cluster-management.io"

var (
	GroupVersion  = schema.GroupVersion{Group: GroupName, Version: "v1alph1"}
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// Install is a function which adds this version to a scheme
	Install = schemeBuilder.AddToScheme

	// SchemeGroupVersion generated code relies on this name
	// Deprecated
	SchemeGroupVersion = GroupVersion
	// AddToScheme exists solely to keep the old generators creating valid code
	// DEPRECATED
	AddToScheme = schemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion,
		&workapiv1.ManifestWork{},
		&workapiv1.ManifestWorkList{},
	)
	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}
