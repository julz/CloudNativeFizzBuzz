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
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
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
	log := r.Log.WithValues("query", req.NamespacedName)

	query := &fizzbuzzv1beta1.Query{}
	if err := r.Get(ctx, req.NamespacedName, query); client.IgnoreNotFound(err) != nil {
		return ctrl.Result{}, err
	}

	log.Info("reconciling query", "input", query.Spec.Input)

	calc := &fizzbuzzv1beta1.Calculation{
		ObjectMeta: v1.ObjectMeta{
			Name:      query.Name + "-fizz",
			Namespace: query.Namespace,
		},
		Spec: fizzbuzzv1beta1.CalculationSpec{
			Calc: "input % 3",
			Vars: []fizzbuzzv1beta1.Var{{
				Name:  "input",
				Value: query.Spec.Input,
			}},
		},
	}

	if err := r.Create(ctx, calc); err != nil {
		if !errors.IsAlreadyExists(err) {
			return ctrl.Result{}, err
		}
	} else {
		log.Info("created calculation, waiting for result")
		return ctrl.Result{RequeueAfter: 500 * time.Millisecond}, nil
	}

	if err := r.Get(ctx, types.NamespacedName{Namespace: query.Namespace, Name: query.Name + "-fizz"}, calc); err != nil {
		return ctrl.Result{}, err
	}

	if len(calc.Status.Conditions) == 0 {
		log.Info("calculation not ready, requeuing query")
		return ctrl.Result{Requeue: true}, nil
	}

	log.Info("calculation ready", "ready", calc.Status.Result)
	if calc.Status.Result == "0" {
		// fizz!
		log.Info(".. and it's a fizz")
		query.Status.Fizz = true
		query.Status.Conditions = []fizzbuzzv1beta1.QueryCondition{}
		if err := r.Status().Update(ctx, query); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil

}

func (r *QueryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&fizzbuzzv1beta1.Query{}).
		Complete(r)
}
