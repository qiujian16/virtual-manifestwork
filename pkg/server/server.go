package server

import (
	"github.com/qiujian16/virtual-manifestwork/pkg/api"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/informers"
	workclientset "open-cluster-management.io/api/client/work/clientset/versioned"
	workinformers "open-cluster-management.io/api/client/work/informers/externalversions"
)

type VirtualServer struct {
	*genericapiserver.GenericAPIServer
}

func NewVirtualServer(
	client workclientset.Interface,
	informerFactory informers.SharedInformerFactory,
	workInformerFactory workinformers.SharedInformerFactory,
	apiServerConfig *genericapiserver.Config,
) (*VirtualServer, error) {
	apiServer, err := apiServerConfig.Complete(informerFactory).New("proxy-server", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	if err := api.InstallVirtualManifestWorkGroup(apiServer, client, workInformerFactory); err != nil {
		return nil, err
	}

	return &VirtualServer{apiServer}, nil
}

func (p *VirtualServer) Run(stopCh <-chan struct{}) error {
	return p.GenericAPIServer.PrepareRun().Run(stopCh)
}
