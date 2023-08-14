/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bar

import (
	"context"

	"github.com/Troy876/toolkit-operator-2367/api/v1alpha1"
	"github.com/Troy876/toolkit-operator-2367/loader"

	"github.com/go-logr/logr"
	"github.com/redhat-appstudio/operator-toolkit/controller"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// BarReconciler reconciles a Bar object
type Controller struct {
	client.Client
	Scheme *runtime.Scheme
	log    *logr.Logger
}

//+kubebuilder:rbac:groups=sample.redhat.com,resources=bars,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=sample.redhat.com,resources=bars/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=sample.redhat.com,resources=bars/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Bar object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *Controller) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	logger := r.log.WithValues("Bar", req.NamespacedName)

	// Gets the bar resource
	bar := &v1alpha1.Bar{}
	// Gets the resource from the cluster respective to the controller
	err := r.Client.Get(ctx, req.NamespacedName, bar)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	adapter := newAdapter(ctx, r.Client, &v1alpha1.Bar{}, loader.NewLoader(), &logger)

	return controller.ReconcileHandler([]controller.Operation{
		adapter.EnsureOwnerReferenceIsSet,
		adapter.EnsureBarIsTiedToFoo,
	})
}

func (c *Controller) Register(mgr ctrl.Manager, log *logr.Logger, cluster cluster.Cluster) error {
	c.Client = mgr.GetClient()
	c.log = log
	c.log.WithName("bar")

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Bar{}).
		Complete(c)
}

// SetupWithManager sets up the controller with the Manager.
func (r *Controller) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Bar{}).
		Complete(r)
}
