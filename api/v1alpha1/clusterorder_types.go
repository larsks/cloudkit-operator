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

// Important: Run "make" to regenerate code after modifying this file
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterOrderSpec defines the desired state of ClusterOrder
type ClusterOrderSpec struct {
  // TemplateID is the unique identigier of the cluster template to use when creating this cluster
	TemplateID string `json:"templateID,omitempty"`
}

type ClusterOrderPhaseType string

const (
	// Initial value
	ClusterOrderPhaseUnknown ClusterOrderPhaseType = "Unknown"

	// Order has been accepted but work has not yet started
	ClusterOrderPhaseAccepted ClusterOrderPhaseType = "Accepted"

	// Cluster update is in progress
	ClusterOrderPhaseProgressing ClusterOrderPhaseType = "Progressing"

	// Failed to create or update cluster
	ClusterOrderPhaseFailed ClusterOrderPhaseType = "Failed"

	// All resources associated with this ClusterOrder are ready
	ClusterOrderPhaseReady ClusterOrderPhaseType = "Ready"
)

type ClusterOrderConditionType string

const (
	ClusterOrderConditionAccepted              ClusterOrderConditionType = "Accepted"
	ClusterOrderConditionProgressing           ClusterOrderConditionType = "Progressing"
	ClusterOrderConditionControlPlaneAvailable ClusterOrderConditionType = "ControlPlaneAvailable"
	ClusterOrderConditionNodePoolAvailable     ClusterOrderConditionType = "NodePoolAvailable"
	ClusterOrderConditionAvailable             ClusterOrderConditionType = "Available"
)

// ClusterOrderStatus defines the observed state of ClusterOrder
type ClusterOrderStatus struct {
	// Phase provides a single-value overview of the state of the ClusterOrder
	Phase ClusterOrderPhaseType `json:"phase,omitempty"`

	// Conditions holds an array of metav1.Condition that describe the state of the ClusterOrder
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterOrder is the Schema for the clusterorders API
type ClusterOrder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterOrderSpec   `json:"spec,omitempty"`
	Status ClusterOrderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterOrderList contains a list of ClusterOrder
type ClusterOrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterOrder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterOrder{}, &ClusterOrderList{})
}
