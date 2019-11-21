/*

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

// Package controllers contains the controllers
package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	fizzbuzzv1beta1 "github.com/julz/cloudnativefizzbuzz/api/v1beta1"
)

// QueryReconciler reconciles a Query object
type QueryReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=fizzbuzz.my.domain,resources=queries,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=fizzbuzz.my.domain,resources=queries/status,verbs=get;update;patch

func (r *QueryReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("query", req.NamespacedName)

	query := &fizzbuzzv1beta1.Query{}
	if err := r.Get(ctx, req.NamespacedName, query); client.IgnoreNotFound(err) != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *QueryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&fizzbuzzv1beta1.Query{}).
		Complete(r)
}
