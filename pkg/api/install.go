package api

import (
	"github.com/qiujian16/virtual-manifestwork/pkg/rest/virtualmanifestwork"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	workclientset "open-cluster-management.io/api/client/work/clientset/versioned"
	workinformers "open-cluster-management.io/api/client/work/informers/externalversions"
	workapiv1 "open-cluster-management.io/api/work/v1"
	workapiv1alpha1 "open-cluster-management.io/api/work/v1alpha1"
)

var (
	// Scheme contains the types needed by the resource metrics API.
	Scheme = runtime.NewScheme()
	// ParameterCodec handles versioning of objects that are converted to query parameters.
	ParameterCodec = runtime.NewParameterCodec(Scheme)
	// Codecs is a codec factory for serving the resource metrics API.
	Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {
	// we need to add the options to empty v1
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	workapiv1.Install(Scheme)
	workapiv1alpha1.Install(Scheme)
}

func InstallVirtualManifestWorkGroup(server *genericapiserver.GenericAPIServer, client workclientset.Interface, factory workinformers.SharedInformerFactory) error {
	v1storage := map[string]rest.Storage{
		"manifestworks":        virtualmanifestwork.NewREST(client),
		"manifestworks/status": virtualmanifestwork.NewStatusRest(client),
	}
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(workapiv1.GroupName, Scheme, ParameterCodec, Codecs)
	apiGroupInfo.VersionedResourcesStorageMap[workapiv1.GroupVersion.Version] = v1storage
	return server.InstallAPIGroup(&apiGroupInfo)
}
