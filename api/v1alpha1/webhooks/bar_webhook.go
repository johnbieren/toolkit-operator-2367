package webhooks

import (
	"context"
	"github/troy/sample-operator/api/v1alpha1"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Webhook describes the data structure for the author webhook
type Webhook struct {
	client client.Client
	log    logr.Logger
}

// Register registers the webhook with the passed manager and log.
func (w *Webhook) Register(mgr ctrl.Manager, log *logr.Logger) error {
	w.client = mgr.GetClient()
	w.log = log.WithName("bar")

	return ctrl.NewWebhookManagedBy(mgr).
		For(&v1alpha1.Bar{}).
		WithDefaulter(w).
		//WithValidator(w).
		Complete()

}

type WebhookInterface interface {
	Register(mgr ctrl.Manager, log *logr.Logger) error
}

// Default implements webhook.Defaulter so a webhook will be registered for the type.
func (w *Webhook) Default(ctx context.Context, obj runtime.Object) error {
	return nil
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (w *Webhook) ValidateCreate(ctx context.Context, obj runtime.Object) error {
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
