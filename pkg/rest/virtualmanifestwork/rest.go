package virtualmanifestwork

import (
	"context"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/rest"
	workclientset "open-cluster-management.io/api/client/work/clientset/versioned"
	workapiv1 "open-cluster-management.io/api/work/v1"
)

type REST struct {
	client workclientset.Interface
}

// NewREST returns a RESTStorage object that will work against ManagedCluster resources
func NewREST(
	client workclientset.Interface,
) *REST {
	return &REST{
		client: client,
	}
}

// New returns a new managedCluster
func (s *REST) New() runtime.Object {
	return &workapiv1.ManifestWork{}
}

func (s *REST) NamespaceScoped() bool {
	return true
}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"vmw", "vmws"}
}

func (s *REST) Destroy() {
	return
}

// NewList returns a new managedCluster list
func (*REST) NewList() runtime.Object {
	return &workapiv1.ManifestWorkList{}
}

func (c *REST) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return &metav1.Table{}, nil
}

var _ = rest.Lister(&REST{})

// List retrieves a list of managedCluster that match label.
func (s *REST) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	var v1ListOptions metav1.ListOptions
	if err := metainternalversion.Convert_internalversion_ListOptions_To_v1_ListOptions(options, &v1ListOptions, nil); err != nil {
		return nil, err
	}
	return nil, nil
}

var _ = rest.Watcher(&REST{})

func (s *REST) Watch(ctx context.Context, options *metainternalversion.ListOptions) (watch.Interface, error) {
	var v1ListOptions metav1.ListOptions
	if err := metainternalversion.Convert_internalversion_ListOptions_To_v1_ListOptions(options, &v1ListOptions, nil); err != nil {
		return nil, err
	}
	return nil, nil
}

var _ = rest.Getter(&REST{})

// Get retrieves a managedCluster by name
func (s *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return nil, nil
}

var _ = rest.Patcher(&REST{})

func (s *REST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	return nil, false, nil
}
