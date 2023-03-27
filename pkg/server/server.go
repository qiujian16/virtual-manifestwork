package server

import (
	"context"
	"github.com/qiujian16/virtual-manifestwork/pkg/api"
	"github.com/qiujian16/virtual-manifestwork/pkg/controllers"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/informers"
	"k8s.io/klog/v2"
	workclientset "open-cluster-management.io/api/client/work/clientset/versioned"
	workinformers "open-cluster-management.io/api/client/work/informers/externalversions"
)

type VirtualServer struct {
	*genericapiserver.GenericAPIServer
	client              workclientset.Interface
	workInformerFactory workinformers.SharedInformerFactory
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

	return &VirtualServer{GenericAPIServer: apiServer, client: client, workInformerFactory: workInformerFactory}, nil
}

func (p *VirtualServer) Run(stopCh <-chan struct{}) error {
	if err := p.GenericAPIServer.AddPostStartHook("referencework-controller", func(ctx genericapiserver.PostStartHookContext) error {
		return controllers.RunManager(GoContext(ctx), p.client, p.workInformerFactory)
	}); err != nil {
		klog.Errorf("add controller error %v", err)
	}

	return p.GenericAPIServer.PrepareRun().Run(stopCh)
}

func GoContext(hookContext genericapiserver.PostStartHookContext) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func(done <-chan struct{}) {
		<-done
		cancel()
	}(hookContext.StopCh)
	return ctx
}
