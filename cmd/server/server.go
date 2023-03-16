package main

import (
	"fmt"
	"github.com/qiujian16/virtual-manifestwork/cmd/server/options"
	"github.com/qiujian16/virtual-manifestwork/pkg/server"
	"github.com/spf13/pflag"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
	"math/rand"
	workclientset "open-cluster-management.io/api/client/work/clientset/versioned"
	workinformers "open-cluster-management.io/api/client/work/informers/externalversions"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	opts := options.NewOptions()
	opts.AddFlags(pflag.CommandLine)

	klog.InitFlags(nil)
	flag.InitFlags()

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := run(opts, wait.NeverStop); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(s *options.Options, stopCh <-chan struct{}) error {
	if err := s.SetDefaults(); err != nil {
		return err
	}

	if errs := s.Validate(); len(errs) != 0 {
		return utilerrors.NewAggregate(errs)
	}

	clusterCfg, err := clientcmd.BuildConfigFromFlags("", s.KubeConfigFile)
	if err != nil {
		return err
	}

	kubeClient, err := kubernetes.NewForConfig(clusterCfg)
	if err != nil {
		return err
	}

	workClient, err := workclientset.NewForConfig(clusterCfg)
	if err != nil {
		return err
	}

	workInformers := workinformers.NewSharedInformerFactory(workClient, 10*time.Minute)

	informerFactory := informers.NewSharedInformerFactory(kubeClient, 10*time.Minute)

	apiServerConfig, err := s.APIServerConfig()
	if err != nil {
		return err
	}

	virtualServer, err := server.NewVirtualServer(workClient, informerFactory, workInformers, apiServerConfig)
	if err != nil {
		return err
	}

	workInformers.Start(stopCh)
	informerFactory.Start(stopCh)

	return virtualServer.Run(stopCh)
}
