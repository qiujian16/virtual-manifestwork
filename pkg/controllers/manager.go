package controllers

import (
	"context"
	"github.com/qiujian16/virtual-manifestwork/pkg/controllers/referencework"
	workclientset "open-cluster-management.io/api/client/work/clientset/versioned"
	workinformers "open-cluster-management.io/api/client/work/informers/externalversions"
)

func RunManager(ctx context.Context, workClient workclientset.Interface, workInformerFactory workinformers.SharedInformerFactory) error {
	refernceWorkController := referencework.NewReferenceWorkController(
		workClient,
		workInformerFactory.Work().V1alpha1().ReferenceWorks(),
		workInformerFactory.Work().V1alpha1().ManifestWorkReplicaSets(),
	)

	go refernceWorkController.Run(ctx, 2)

	<-ctx.Done()
	return nil
}
