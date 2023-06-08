/*
Copyright 2023 The Kubernetes Authors.

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
// +kubebuilder:object:generate=true
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+genclient
//+genclient:nonNamespaced
//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineGroup specifies parameters that a underlying storage system
// uses when creating a volume group snapshot. A specific VolumeGroupSnapshotClass
// is used by specifying its name in a VolumeGroupSnapshot object.
// VolumeGroupSnapshotClasses are non-namespaced.
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=vgsclass;vgsclasses
type MachineGroup struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Driver is the name of the storage driver expected to handle this VolumeGroupSnapshotClass.
	// Required.
	Name string `json:"driver" protobuf:"bytes,2,opt,name=driver"`

	// Parameters is a key-value map with storage driver specific parameters for
	// creating group snapshots.
	// These values are opaque to Kubernetes and are passed directly to the driver.
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" protobuf:"bytes,3,rep,name=parameters"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineGroupList is a collection of VolumeGroupSnapshotClasses.
// +kubebuilder:object:root=true
type MachineGroupList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of VolumeGroupSnapshotClasses.
	Items []MachineGroup `json:"items" protobuf:"bytes,2,rep,name=items"`
}
