// Copyright 2020-2025 Politecnico di Torino
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:validation:Enum="";"Starting";"Pending";"ResourceQuotaExceeded";"Bound";"Terminating"

// SharedVolumePhase is an enumeration of the different phases associated with a SharedVolume.
type SharedVolumePhase string

const (
	// SharedVolumePhaseUnset -> the shared volume phase is unknown.
	SharedVolumePhaseUnset SharedVolumePhase = ""
	// SharedVolumePhaseCreating -> the shared volume is creating.
	SharedVolumePhaseCreating SharedVolumePhase = "Creating"
	// SharedVolumePhasePending -> the shared volume's PVC is pending.
	SharedVolumePhasePending SharedVolumePhase = "Pending"
	// SharedVolumePhaseProvisioning -> the shared volume's PVC is under provisioning.
	SharedVolumePhaseProvisioning SharedVolumePhase = "Provisioning"
	// SharedVolumePhaseReady -> the shared volume is bound and ready to be accessed.
	SharedVolumePhaseReady SharedVolumePhase = "Ready"
	// SharedVolumePhaseResourceQuotaExceeded -> the shared volume could not be created because the resource quota is exceeded.
	SharedVolumePhaseResourceQuotaExceeded SharedVolumePhase = "ResourceQuotaExceeded"
	// SharedVolumePhaseError -> the shared volume had an error during reconcile.
	SharedVolumePhaseError SharedVolumePhase = "Error"

	// SharedVolumeErrorReasonSmaller -> the shared volume had an error since the size is smaller than before.
	SharedVolumeErrorReasonSmaller string = "Forbidden: Size cannot be less than previous value"
)

// SharedVolumeSpec is the specification of the desired state of the Shared Volume.
type SharedVolumeSpec struct {
	// The human-readable name of the Shared Volume.
	PrettyName string `json:"prettyName"`

	// The size of the volume.
	Size resource.Quantity `json:"size"`
}

// SharedVolumeStatus reflects the most recently observed status of the Shared Volume.
type SharedVolumeStatus struct {
	// The server address to reach the PV //TODO: Va bene come descrizione?
	ServerAddress string `json:"serverAddress,omitempty"`

	// The actual name of the volume retrieved from the PV.
	ExportPath string `json:"exportPath,omitempty"`

	// The current phase of the lifecycle of the Shared Volume.
	Phase SharedVolumePhase `json:"phase,omitempty"`

	// The reason why the Shared Volume is in Error phase.
	ErrorReason string `json:"errorReason,omitempty"` //TODO: OPPURE GUARDA INSTANCE PER FARE DEI LOG???
	//TODO: https://gitkraken.dev/link/dnNjb2RlOi8vZWFtb2Rpby5naXRsZW5zL2xpbmsvci84NzgzZTM5ODNkNzZlZjFjNWFhMDljZGViOThhYTdmODliZjg5OWFkL2Yvb3BlcmF0b3JzL3BrZy9leGFtYWdlbnQvaW5zdGFuY2UuZ28%2FdXJsPWh0dHBzJTNBJTJGJTJGZ2l0aHViLmNvbSUyRm5ldGdyb3VwLXBvbGl0byUyRkNyb3duTGFicyZsaW5lcz0xMjU%3D?origin=gitlens
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName="shvol"
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Pretty Name",type=string,JSONPath=`.spec.prettyName`
// +kubebuilder:printcolumn:name="Size",type=string,JSONPath=`.spec.size`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// SharedVolume describes a shared volume between tenants in CrownLabs.
type SharedVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SharedVolumeSpec   `json:"spec,omitempty"`
	Status SharedVolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedVolumeList contains a list of SharedVolume objects.
type SharedVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []SharedVolume `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SharedVolume{}, &SharedVolumeList{})
}
