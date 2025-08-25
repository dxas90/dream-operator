/*
Copyright 2025.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ServiceConfig struct {

	// Param1 is a sample configuration parameter
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Param1 string `json:"param1,omitempty"`
	// Param2 is another sample configuration parameter
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Param2 string `json:"param2,omitempty"`
	// EnabledOptions is a list of enabled options for the service
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	EnabledOptions []string `json:"enabledOptions,omitempty"`
}

// LearnSpec defines the desired state of Learn.
type LearnSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html
	// +kubebuilder:validation:MinLength=2
	// +kubebuilder:validation:MaxLength=32

	// DisplayName is the display name for the Learn instance
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	DisplayName string `json:"displayName,omitempty"`

	// Description provides a brief description of the Learn instance
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Description string `json:"description,omitempty"`

	// ServiceConfig holds configuration parameters for the service
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ServiceConfig ServiceConfig `json:"serviceConfig,omitempty"`
}

// LearnStatus defines the observed state of Learn.
type LearnStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Represents the observations of a Learn's current state.
	// Learn.status.conditions.type are: "Available", "Progressing", and "Degraded"
	// Learn.status.conditions.status are one of True, False, Unknown.
	// Learn.status.conditions.reason the value should be a CamelCase string and producers of specific
	// condition types may define expected values and meanings for this field, and whether the values
	// are considered a guaranteed API.
	// Learn.status.conditions.Message is a human readable message indicating details about the transition.
	// For further information see: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties

	// Conditions store the status conditions of the Learn instances
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Learn is the Schema for the learns API.
type Learn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LearnSpec   `json:"spec,omitempty"`
	Status LearnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LearnList contains a list of Learn.
type LearnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Learn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Learn{}, &LearnList{})
}
