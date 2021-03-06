/*
Copyright 2018 thejasbabu.

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
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
type Selector struct {
	AppName string `json:"appName"`
}

type Metadata struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type Template struct {
	Metadata Metadata       `json:"metadata"`
	PodSpec  corev1.PodSpec `json:"podSpec"`
}

type Deployment struct {
	Replicas int32                     `json:"replicas,omitempty"`
	Selector Selector                  `json:"selector"`
	Template Template                  `json:"template"`
	Strategy appsv1.DeploymentStrategy `json:"strategy,omitempty"`
	Metadata Metadata                  `json:"metadata"`
}

// FullStackSpec defines the desired state of FullStack
type FullStackSpec struct {
	Namespace  string     `json:"namespace"`
	AppName    string     `json:"appName"`
	Deployment Deployment `json:"deployment"`
}

// FullStackStatus defines the observed state of FullStack
type FullStackStatus struct {
	Namespace string `json:"namespace"`
	AppName   string `json:"appName"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FullStack is the Schema for the fullstacks API
// +k8s:openapi-gen=true
type FullStack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FullStackSpec   `json:"spec,omitempty"`
	Status FullStackStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FullStackList contains a list of FullStack
type FullStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FullStack `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FullStack{}, &FullStackList{})
}
