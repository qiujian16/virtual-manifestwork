package client

import (
	"context"
	"errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/klog/v2"
	workclientset "open-cluster-management.io/api/client/work/clientset/versioned"
	workapiv1 "open-cluster-management.io/api/work/v1"
	workapiv1alpha1 "open-cluster-management.io/api/work/v1alpha1"
)

type VirtualManifestWork interface {
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*workapiv1.ManifestWork, error)
	List(ctx context.Context, opts metav1.ListOptions) (*workapiv1.ManifestWorkList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type virtualManifestWork struct {
	client workclientset.Interface
	ns     string
}

func NewVirtualManifestWork(ctx context.Context, c workclientset.Interface) *virtualManifestWork {
	ns, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		ns = metav1.NamespaceAll
	}
	return &virtualManifestWork{
		client: c,
		ns:     ns,
	}
}

func (v *virtualManifestWork) Get(ctx context.Context, name string, opts metav1.GetOptions) (*workapiv1.ManifestWork, error) {
	refWork, err := v.client.WorkV1alpha1().ReferenceWorks(v.ns).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}

	return convertToManifestWork(ctx, refWork, v.client)
}

func (v *virtualManifestWork) List(ctx context.Context, opts metav1.ListOptions) (*workapiv1.ManifestWorkList, error) {
	refWorks, err := v.client.WorkV1alpha1().ReferenceWorks(v.ns).List(ctx, opts)
	if err != nil {
		return nil, err
	}

	list := &workapiv1.ManifestWorkList{
		Items: []workapiv1.ManifestWork{},
	}

	for _, refWork := range refWorks.Items {
		mw, err := convertToManifestWork(ctx, &refWork, v.client)
		if err != nil {
			return nil, err
		}
		list.Items = append(list.Items, *mw)
	}
	return list, nil
}

func (v *virtualManifestWork) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	result, err := v.client.WorkV1alpha1().ReferenceWorks(v.ns).Watch(ctx, opts)
	if err != nil {
		return result, err
	}

	transformingWatcher := NewTransformingWatcher(result, func(event watch.Event) *watch.Event {
		transformed := event
		eventType := event.Type
		if eventType == watch.Bookmark || eventType == watch.Error {
			return &transformed
		}
		resource, ok := event.Object.(*workapiv1alpha1.ReferenceWork)
		if !ok {
			errorMessage := "watch expected a resource of type *ReferenceWork"
			klog.Errorf(errorMessage)
			transformed.Type = watch.Error
			transformed.Object = &metav1.Status{
				Status:  "Failure",
				Reason:  metav1.StatusReasonUnknown,
				Message: errorMessage,
				Code:    500,
			}
			return &transformed
		}
		if transformedResource, err := convertToManifestWork(ctx, resource, v.client); err != nil {
			if kerrors.IsNotFound(err) {
				return nil
			}
			transformed.Type = watch.Error
			statusError := &kerrors.StatusError{}
			if errors.As(err, &statusError) {
				transformed.Object = statusError.ErrStatus.DeepCopy()
			} else {
				transformed.Object = &metav1.Status{
					Status:  "Failure",
					Reason:  metav1.StatusReasonUnknown,
					Message: "Watch transformation failed",
					Code:    500,
					Details: &metav1.StatusDetails{
						Name:  resource.GetName(),
						Group: resource.GroupVersionKind().Group,
						Kind:  resource.GroupVersionKind().Kind,
						Causes: []metav1.StatusCause{
							{
								Type:    metav1.CauseTypeUnexpectedServerResponse,
								Message: err.Error(),
							},
						},
					},
				}
			}
		} else {
			transformed.Object = transformedResource
		}
		return &transformed
	})
	return transformingWatcher, nil
}

func convertToManifestWork(ctx context.Context, refWork *workapiv1alpha1.ReferenceWork, client workclientset.Interface) (*workapiv1.ManifestWork, error) {
	workSet, err := client.WorkV1alpha1().ManifestWorkReplicaSets(refWork.Spec.Reference.Namespace).Get(ctx, refWork.Spec.Reference.Name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		return &workapiv1.ManifestWork{
			ObjectMeta: refWork.ObjectMeta,
			Spec:       workapiv1.ManifestWorkSpec{},
			Status:     refWork.Status.ManifestWorkStatus,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	return &workapiv1.ManifestWork{
		ObjectMeta: refWork.ObjectMeta,
		Spec:       workSet.Spec.ManifestWorkTemplate,
		Status:     refWork.Status.ManifestWorkStatus,
	}, nil
}
