/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	esv1alpha1 "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// SensorParameters are the configurable fields of a Sensor.
type SensorParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// SensorObservation are the observable fields of a Sensor.
type SensorObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A SensorSpec defines the desired state of a Sensor.
type SensorSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       esv1alpha1.Sensor `json:"forProvider"`
}

// A SensorStatus represents the observed state of a Sensor.
type SensorStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SensorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Sensor is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,argo}
type Sensor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SensorSpec   `json:"spec"`
	Status SensorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SensorList contains a list of Sensor
type SensorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Sensor `json:"items"`
}

// Sensor type metadata.
var (
	SensorKind             = reflect.TypeOf(Sensor{}).Name()
	SensorGroupKind        = schema.GroupKind{Group: Group, Kind: SensorKind}.String()
	SensorKindAPIVersion   = SensorKind + "." + SchemeGroupVersion.String()
	SensorGroupVersionKind = SchemeGroupVersion.WithKind(SensorKind)
)

func init() {
	SchemeBuilder.Register(&Sensor{}, &SensorList{})
}
