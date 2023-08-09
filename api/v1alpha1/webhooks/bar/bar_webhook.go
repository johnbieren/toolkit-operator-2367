package webhooks

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/redhat-appstudio/operator-toolkit-example/loader"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/redhat-appstudio/operator-toolkit-example/api/v1alpha1"
)

// Webhook describes the data structure for the author webhook
type Webhook struct {
	client client.Client
	loader loader.ObjectLoader
	log    logr.Logger
}

// Register registers the webhook with the passed manager and log.
func (w *Webhook) Register(mgr ctrl.Manager, log *logr.Logger) error {
	w.client = mgr.GetClient()
	w.loader = loader.NewLoader()
	w.log = log.WithName("bar")

	return ctrl.NewWebhookManagedBy(mgr).
		For(&v1alpha1.Bar{}).
		WithDefaulter(w).
		Complete()

}

type WebhookInterface interface {
	Register(mgr ctrl.Manager, log *logr.Logger) error
}

func (w *Webhook) Default(ctx context.Context, obj runtime.Object) error {
	return nil
}

// +kubebuilder:webhook:path=/validate-appstudio-redhat-com-v1alpha1-bar,mutating=false,failurePolicy=fail,sideEffects=None,groups=appstudio.redhat.com,resources=bars,verbs=create;update,versions=v1alpha1,name=vbar.kb.io,admissionReviewVersions=v1

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (w *Webhook) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	bar := obj.(*v1alpha1.Bar)

	_, err := w.loader.GetFoo(ctx, w.client, bar.Spec.Foo, bar.Namespace)
	if err != nil {
		return fmt.Errorf("resource references an unexistent Foo resource (%s/%s)", bar.Namespace, bar.Spec.Foo)
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (w *Webhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (w *Webhook) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	return nil
}
