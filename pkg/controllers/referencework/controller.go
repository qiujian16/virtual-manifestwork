package referencework

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"open-cluster-management.io/addon-framework/pkg/basecontroller/factory"
	workclientset "open-cluster-management.io/api/client/work/clientset/versioned"
	workv1alpha1informer "open-cluster-management.io/api/client/work/informers/externalversions/work/v1alpha1"
	workv1alpha1lister "open-cluster-management.io/api/client/work/listers/work/v1alpha1"
	workapiv1alpha1 "open-cluster-management.io/api/work/v1alpha1"
)

const (
	workReferenceByWorkReplicaSet = "workReferenceByWorkReplicaSet"
)

type referenceWorkController struct {
	client               workclientset.Interface
	workReplicaSetLister workv1alpha1lister.ManifestWorkReplicaSetLister
	referenceWorkIndexer cache.Indexer
}

func NewReferenceWorkController(
	client workclientset.Interface,
	referenceWorkInformer workv1alpha1informer.ReferenceWorkInformer,
	workReplicaSetInformer workv1alpha1informer.ManifestWorkReplicaSetInformer,
) factory.Controller {
	c := &referenceWorkController{
		client:               client,
		referenceWorkIndexer: referenceWorkInformer.Informer().GetIndexer(),
		workReplicaSetLister: workReplicaSetInformer.Lister(),
	}

	err := referenceWorkInformer.Informer().AddIndexers(
		cache.Indexers{
			workReferenceByWorkReplicaSet: indexReferenceByWorkReplicaSet,
		})
	if err != nil {
		utilruntime.HandleError(err)
	}

	return factory.New().WithInformersQueueKeysFunc(
		func(obj runtime.Object) []string {
			key, _ := cache.MetaNamespaceKeyFunc(obj)
			return []string{key}
		}, workReplicaSetInformer.Informer()).
		WithInformersQueueKeysFunc(workReplicaSetByRefernceQueueKey, referenceWorkInformer.Informer()).
		WithSync(c.sync).ToController("referencework-controller")
}

func (c *referenceWorkController) sync(ctx context.Context, syncCtx factory.SyncContext, key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		// ignore addon whose key is invalid
		return nil
	}

	klog.V(4).Infof("Reconciling addon %q", key)

	mwrs, err := c.workReplicaSetLister.ManifestWorkReplicaSets(namespace).Get(name)
	switch {
	case errors.IsNotFound(err):
		return nil
	case err != nil:
		return err
	}

	// get all related reference work
	refWorks, err := c.getReferenceWorksByWorkReplicaSet(key)
	if err != nil {
		return err
	}

	desiredSpecHash := hashOfResourceStruct(mwrs.Spec.ManifestWorkTemplate)
	var errs []error
	for _, refWork := range refWorks {
		err := c.applySpecHash(ctx, refWork, desiredSpecHash)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return utilerrors.NewAggregate(errs)
}

func (c *referenceWorkController) applySpecHash(ctx context.Context, refWork *workapiv1alpha1.ReferenceWork, specHash string) error {
	if refWork.Status.ReferenceSpecHash == specHash {
		return nil
	}

	oldData, err := json.Marshal(&workapiv1alpha1.ReferenceWork{
		Status: workapiv1alpha1.ReferenceWorkStatus{
			ReferenceSpecHash: refWork.Status.ReferenceSpecHash,
		},
	})
	if err != nil {
		return err
	}

	newData, err := json.Marshal(&workapiv1alpha1.ReferenceWork{
		ObjectMeta: metav1.ObjectMeta{
			ResourceVersion: refWork.ResourceVersion,
			UID:             refWork.UID,
		},
		Status: workapiv1alpha1.ReferenceWorkStatus{
			ReferenceSpecHash: specHash,
		},
	})

	patchBytes, err := jsonpatch.CreateMergePatch(oldData, newData)
	if err != nil {
		return fmt.Errorf("failed to create patch for referencework %s/%s: %w", refWork.Namespace, refWork.Name, err)
	}

	_, err = c.client.WorkV1alpha1().ReferenceWorks(refWork.Namespace).Patch(ctx, refWork.Name, types.MergePatchType, patchBytes, metav1.PatchOptions{}, "status")
	return err
}

func (c *referenceWorkController) getReferenceWorksByWorkReplicaSet(key string) ([]*workapiv1alpha1.ReferenceWork, error) {
	objs, err := c.referenceWorkIndexer.ByIndex(workReferenceByWorkReplicaSet, key)
	if err != nil {
		utilruntime.HandleError(err)
		return nil, err
	}

	var refWorks []*workapiv1alpha1.ReferenceWork
	for _, o := range objs {
		refWork := o.(*workapiv1alpha1.ReferenceWork)
		refWorks = append(refWorks, refWork)
	}
	return refWorks, nil
}

func indexReferenceByWorkReplicaSet(obj interface{}) ([]string, error) {
	refWork, ok := obj.(*workapiv1alpha1.ReferenceWork)

	if !ok {
		return []string{}, fmt.Errorf("obj %T is not a ReferenceWork", obj)
	}
	return []string{fmt.Sprintf("%s/%s", refWork.Spec.Reference.Namespace, refWork.Spec.Reference.Name)}, nil
}

func workReplicaSetByRefernceQueueKey(obj runtime.Object) []string {
	keys, err := indexReferenceByWorkReplicaSet(obj)
	if err != nil {
		utilruntime.HandleError(err)
	}

	return keys
}

func hashOfResourceStruct(o interface{}) string {
	oString := fmt.Sprintf("%v", o)
	h := md5.New()
	io.WriteString(h, oString)
	rval := fmt.Sprintf("%x", h.Sum(nil))
	return rval
}
