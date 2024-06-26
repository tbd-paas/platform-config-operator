/*
Copyright 2024.

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
	"errors"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	"github.com/nukleros/operator-builder-tools/pkg/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ErrUnableToConvertPlatformConfig = errors.New("unable to convert to PlatformConfig")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PlatformConfigSpec defines the desired state of PlatformConfig.
type PlatformConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	Platform PlatformConfigSpecPlatform `json:"platform,omitempty"`

	// +kubebuilder:validation:Optional
	Cloud PlatformConfigSpecCloud `json:"cloud,omitempty"`
}

type PlatformConfigSpecPlatform struct {
	// +kubebuilder:validation:Optional
	Certificates PlatformConfigSpecPlatformCertificates `json:"certificates,omitempty"`

	// +kubebuilder:validation:Optional
	Identity PlatformConfigSpecPlatformIdentity `json:"identity,omitempty"`
}

type PlatformConfigSpecPlatformCertificates struct {
	// +kubebuilder:default="tbd-certificates-system"
	// +kubebuilder:validation:Optional
	// (Default: "tbd-certificates-system")
	// Namespace where
	//
	//	the capability components will be deployed.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:default="small"
	// +kubebuilder:validation:Optional
	// (Default: "small")
	// Size of the
	//
	//	deployment for the underlying capability.  Must be one of small, medium, or large.
	//	+kubebuilder:validation:Enum:small;medium;large
	DeploymentSize string `json:"deploymentSize,omitempty"`
}

type PlatformConfigSpecPlatformIdentity struct {
	// +kubebuilder:default="tbd-identity-system"
	// +kubebuilder:validation:Optional
	// (Default: "tbd-identity-system")
	// Namespace where
	//
	//	the capability components will be deployed.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:default="small"
	// +kubebuilder:validation:Optional
	// (Default: "small")
	// Size of the
	//
	//	+kubebuilder:validation:Enum:small;medium;large
	//	deployment for the underlying capability.  Must be one of small, medium, or large.
	DeploymentSize string `json:"deploymentSize,omitempty"`
}

type PlatformConfigSpecCloud struct {
	// +kubebuilder:default="aws"
	// +kubebuilder:validation:Optional
	// (Default: "aws")
	//
	//	+kubebuilder:validation:Enum=aws
	//	Underlying cloud type this platform is deployed upon.  Currently, only AWS is supported.
	Type string `json:"type,omitempty"`

	// +kubebuilder:default=true
	// +kubebuilder:validation:Optional
	// (Default: true)
	//
	//	Whether this cloud is deployed as a local cloud to use for testing scenarios.
	Local bool `json:"local,omitempty"`
}

// PlatformConfigStatus defines the observed state of PlatformConfig.
type PlatformConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Created               bool                     `json:"created,omitempty"`
	DependenciesSatisfied bool                     `json:"dependenciesSatisfied,omitempty"`
	Conditions            []*status.PhaseCondition `json:"conditions,omitempty"`
	Resources             []*status.ChildResource  `json:"resources,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// PlatformConfig is the Schema for the platformconfigs API.
type PlatformConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlatformConfigSpec   `json:"spec,omitempty"`
	Status            PlatformConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlatformConfigList contains a list of PlatformConfig.
type PlatformConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlatformConfig `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *PlatformConfig) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *PlatformConfig) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *PlatformConfig) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *PlatformConfig) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *PlatformConfig) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *PlatformConfig) SetPhaseCondition(condition *status.PhaseCondition) {
	for i, currentCondition := range component.GetPhaseConditions() {
		if currentCondition.Phase == condition.Phase {
			component.Status.Conditions[i] = condition

			return
		}
	}

	// phase not found, lets add it to the list.
	component.Status.Conditions = append(component.Status.Conditions, condition)
}

// GetResources returns the child resource status for a component.
func (component *PlatformConfig) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *PlatformConfig) SetChildResourceCondition(resource *status.ChildResource) {
	for i, currentResource := range component.GetChildResourceConditions() {
		if currentResource.Group == resource.Group && currentResource.Version == resource.Version && currentResource.Kind == resource.Kind {
			if currentResource.Name == resource.Name && currentResource.Namespace == resource.Namespace {
				component.Status.Resources[i] = resource

				return
			}
		}
	}

	// phase not found, lets add it to the collection
	component.Status.Resources = append(component.Status.Resources, resource)
}

// GetDependencies returns the dependencies for a component.
func (*PlatformConfig) GetDependencies() []workload.Workload {
	return []workload.Workload{}
}

// GetComponentGVK returns a GVK object for the component.
func (*PlatformConfig) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("PlatformConfig")
}

func init() {
	SchemeBuilder.Register(&PlatformConfig{}, &PlatformConfigList{})
}
