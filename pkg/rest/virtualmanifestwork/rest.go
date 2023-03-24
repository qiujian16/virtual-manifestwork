package virtualmanifestwork

import (
	"context"
	vwclient "github.com/qiujian16/virtual-manifestwork/pkg/client"
	"k8s.io/apimachinery/pkg/api/meta"
	metatable "k8s.io/apimachinery/pkg/api/meta/table"
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
	headers := []metav1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: "Name is the name of manifestwork."},
		{Name: "Age", Type: "date", Description: "Age represents the age of the manifestworks until created."},
	}
	table := &metav1.Table{}
	opt, ok := tableOptions.(*metav1.TableOptions)
	noHeaders := ok && opt != nil && opt.NoHeaders
	if !noHeaders {
		table.ColumnDefinitions = headers
	}

	if m, err := meta.ListAccessor(object); err == nil {
		table.ResourceVersion = m.GetResourceVersion()
		table.Continue = m.GetContinue()
		table.RemainingItemCount = m.GetRemainingItemCount()
	} else {
		if m, err := meta.CommonAccessor(object); err == nil {
			table.ResourceVersion = m.GetResourceVersion()
		}
	}
	var err error
	table.Rows, err = metatable.MetaToTableRow(object, func(obj runtime.Object, m metav1.Object, name, age string) ([]interface{}, error) {
		return []interface{}{}, nil
	})
	if err != nil {
		return nil, err
	}

	return table, nil
}

var _ = rest.Lister(&REST{})

// List retrieves a list of managedCluster that match label.
func (s *REST) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	var v1ListOptions metav1.ListOptions
	if err := metainternalversion.Convert_internalversion_ListOptions_To_v1_ListOptions(options, &v1ListOptions, nil); err != nil {
		return nil, err
	}
	vclient := vwclient.NewVirtualManifestWork(ctx, s.client)
	return vclient.List(ctx, v1ListOptions)
}

var _ = rest.Watcher(&REST{})

func (s *REST) Watch(ctx context.Context, options *metainternalversion.ListOptions) (watch.Interface, error) {
	var v1ListOptions metav1.ListOptions
	if err := metainternalversion.Convert_internalversion_ListOptions_To_v1_ListOptions(options, &v1ListOptions, nil); err != nil {
		return nil, err
	}
	vclient := vwclient.NewVirtualManifestWork(ctx, s.client)
	return vclient.Watch(ctx, v1ListOptions)
}

var _ = rest.Getter(&REST{})

// Get retrieves a managedCluster by name
func (s *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	vclient := vwclient.NewVirtualManifestWork(ctx, s.client)
	return vclient.Get(ctx, name, *options)
}
