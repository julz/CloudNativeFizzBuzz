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

// QuerySpec defines the desired state of Query
type QuerySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Query. Edit Query_types.go to remove/update
	Input int `json:"input,omitempty"`
}

// QueryStatus defines the observed state of Query
type QueryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Fizz bool `json:"fizz"`
	Buzz bool `json:"buzz"`

	Conditions []QueryCondition `json:"conditions"`
}

type QueryCondition struct {
	LastTransitionTime metav1.Time `json:"last_transition_time"`
	Message            string      `json:"message"`
	Reason             string      `json:"reason"`
	Status             string      `json:"status"`
	Type               string      `json:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Query is the Schema for the queries API
type Query struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuerySpec   `json:"spec,omitempty"`
	Status QueryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueryList contains a list of Query
type QueryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Query `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Query{}, &QueryList{})
}
