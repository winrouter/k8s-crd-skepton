/*
Copyright 2020 The Kubernetes Authors.

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
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeSnapshot is a user's request for either creating a point-in-time
// snapshot of a persistent volume, or binding to a pre-existing snapshot.
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=vs
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Name",type=boolean,JSONPath=`.spec.Name`,description="Indicates if the snapshot is ready to be used to restore a volume."
// +kubebuilder:printcolumn:name="Arch",type=string,JSONPath=`.spec.Arch`,description="If a new snapshot needs to be created, this contains the name of the source PVC from which this snapshot was (or will be) created."
// +kubebuilder:printcolumn:name="UpdateTime",type=date,JSONPath=`.status.UpdateTime`,description="Timestamp when the point-in-time snapshot was taken by the underlying storage system."
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type Machine struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// spec defines the desired characteristics of a machine requested by a user.
	// More info: https://kubernetes.io/docs/concepts/storage/volume-snapshots#volumesnapshots
	// Required.
	Spec MachineSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`

	// status represents the current information of a snapshot.
	// Consumers must verify binding between VolumeSnapshot and
	// VolumeSnapshotContent objects is successful (by validating that both
	// VolumeSnapshot and VolumeSnapshotContent point at each other) before
	// using this object.
	// +optional
	Status *MachineStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MachineList is a list of Machine objects
type MachineList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of VolumeSnapshots
	Items []Machine `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// MachineSpec describes the common attributes of a machine.
type MachineSpec struct {
	// source specifies where a machine will be created from.
	// This field is immutable after creation.
	// Required.
	Name *string `json:"source" protobuf:"bytes,1,opt,name=source"`

	// source specifies where a machine will be created from.
	// This field is immutable after creation.
	// Required.
	Arch *string `json:"source" protobuf:"bytes,1,opt,name=source"`

}

// MachineStatus is the status of the Machine
// Note that CreationTime, RestoreSize, ReadyToUse, and Error are in both
// VolumeSnapshotStatus and VolumeSnapshotContentStatus. Fields in VolumeSnapshotStatus
// are updated based on fields in VolumeSnapshotContentStatus. They are eventual
// consistency. These fields are duplicate in both objects due to the following reasons:
//   - Fields in VolumeSnapshotContentStatus can be used for filtering when importing a
//     volumesnapshot.
//   - VolumsnapshotStatus is used by end users because they cannot see VolumeSnapshotContent.
//   - CSI snapshotter sidecar is light weight as it only watches VolumeSnapshotContent
//     object, not VolumeSnapshot object.
type MachineStatus struct {
	// UpdateTime is the timestamp when the point-in-time snapshot is taken
	// by the underlying storage system.
	// In dynamic snapshot creation case, this field will be filled in by the
	// snapshot controller with the "creation_time" value returned from CSI
	// "CreateSnapshot" gRPC call.
	// For a pre-existing snapshot, this field will be filled with the "creation_time"
	// value returned from the CSI "ListSnapshots" gRPC call if the driver supports it.
	// If not specified, it may indicate that the creation time of the snapshot is unknown.
	// +optional
	UpdateTime *metav1.Time `json:"creationTime,omitempty" protobuf:"bytes,2,opt,name=creationTime"`

	// readyToUse indicates if the snapshot is ready to be used to restore a volume.
	// In dynamic snapshot creation case, this field will be filled in by the
	// snapshot controller with the "ready_to_use" value returned from CSI
	// "CreateSnapshot" gRPC call.
	// For a pre-existing snapshot, this field will be filled with the "ready_to_use"
	// value returned from the CSI "ListSnapshots" gRPC call if the driver supports it,
	// otherwise, this field will be set to "True".
	// If not specified, it means the readiness of a snapshot is unknown.
	// +optional
	ReadyToUse *bool `json:"readyToUse,omitempty" protobuf:"varint,3,opt,name=readyToUse"`
}

