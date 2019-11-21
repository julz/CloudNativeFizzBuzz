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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CalculationSpec defines the desired state of Calculation
type CalculationSpec struct {
	Calc string `json:"calc,omitempty"`

	Vars []Var `json:"vars,omitempty"`
}

type Var struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// CalculationStatus defines the observed state of Calculation
type CalculationStatus struct {
	Result string `json:"result"`

	Conditions []CalculationCondition `json:"conditions"`
}

type CalculationCondition struct {
	LastTransitionTime metav1.Time `json:"last_transition_time"`
	Message            string      `json:"message"`
	Reason             string      `json:"reason"`
	Status             string      `json:"status"`
	Type               string      `json:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Calculation is the Schema for the calculations API
type Calculation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CalculationSpec   `json:"spec,omitempty"`
	Status CalculationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CalculationList contains a list of Calculation
type CalculationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Calculation `json:"items,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Calculation{}, &CalculationList{})
}
