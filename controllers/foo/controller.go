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

package foo

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/redhat-appstudio/operator-toolkit-example/api/v1alpha1"
	"github.com/redhat-appstudio/operator-toolkit/controller"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	samplev1alpha1 "github/troy/sample-operator/api/v1alpha1"
	"github/troy/sample-operator/loader"
)

// FooReconciler reconciles a Foo object
type Controller struct {
	client.Client
	log *logr.Logger
}

//+kubebuilder:rbac:groups=sample.redhat.com,resources=foos,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=sample.redhat.com,resources=foos/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=sample.redhat.com,resources=foos/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Foo object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *Controller) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.log.WithValues("Foo", req.NamespacedName)
	// TODO(user): your logic here
	ctrl.Log.Info("Sample operator")

	// Gets the foo resource from the API
	foo := &samplev1alpha1.Foo{}
	// Gets the foo resource from the cluster
	err := r.Client.Get(ctx, req.NamespacedName, foo)
	// Checks for errors, if none, continue
	if err != nil {
		if errors.IsNotFound(err) {
			// The "," means several items are returned; results and nil is returned to err
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	adapter := newAdapter(ctx, r.Client, &v1alpha1.Foo{}, loader.NewLoader(), &logger)

	return controller.ReconcileHandler([]controller.Operation{
		adapter.EnsureFinalizersAreCalled,
		adapter.EnsureFinalizerIsAdded,
		adapter.EnsureMaximumReplicas,
		adapter.EnsureMinimumReplicas,
		adapter.EnsureReplicaDataConsistency,
	})

}

func (c *Controller) Register(mgr ctrl.Manager, log *logr.Logger, cluster cluster.Cluster) error {
	c.Client = mgr.GetClient()
	c.log = log
	c.log.WithName("foo")

	return ctrl.NewControllerManagedBy(mgr).
		For(&samplev1alpha1.Foo{}, builder.WithPredicates(predicate.GenerationChangedPredicate{})).
		Complete(c)
}

func (c *Controller) SetupCache(mgr ctrl.Manager) error {
	indexFunc := func(obj client.Object) []string {
		return []string{obj.(*samplev1alpha1.Bar).Spec.Foo}
	}

	return mgr.GetCache().IndexField(context.Background(), &samplev1alpha1.Foo{}, "spec.bar", indexFunc)
}
