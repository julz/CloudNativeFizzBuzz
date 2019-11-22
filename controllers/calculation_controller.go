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

package controllers

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/julz/calc"
	fizzbuzzv1beta1 "github.com/julz/cloudnativefizzbuzz/api/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CalculationReconciler reconciles a Calculation object
type CalculationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=fizzbuzz.my.domain,resources=calculations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=fizzbuzz.my.domain,resources=calculations/status,verbs=get;update;patch

func (r *CalculationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("calculation", req.NamespacedName)

	calculation := &fizzbuzzv1beta1.Calculation{}
	if err := r.Get(ctx, req.NamespacedName, calculation); client.IgnoreNotFound(err) != nil {
		return ctrl.Result{}, err
	}

	varMap := make(map[string]interface{})
	for _, v := range calculation.Spec.Vars {
		varMap[v.Name] = v.Value
	}

	result, err := calc.EvalVars(calculation.Spec.Calc, varMap)
	if err != nil {
		return ctrl.Result{}, err
	}

	calculation.Status.Result = result
	calculation.Status.Conditions = []fizzbuzzv1beta1.CalculationCondition{
		{
			LastTransitionTime: metav1.Time{Time: time.Now()},
			Type:               "CalculationReady",
			Status:             v1.ConditionTrue,
			Message:            "Calculation complete",
			Reason:             "CalculationComplete",
		},
	}
	if err := r.Status().Update(ctx, calculation); client.IgnoreNotFound(err) != nil {
		return ctrl.Result{}, err
	}

	log.Info("updated", "result", result)

	return ctrl.Result{}, nil
}

func (r *CalculationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&fizzbuzzv1beta1.Calculation{}).
		Complete(r)
}
