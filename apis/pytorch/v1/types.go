// Copyright 2019 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	common "github.com/alibaba/kubedl/pkg/job_controller/api/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:defaulter-gen=TypeMeta
// +resource:path=pytorchjob
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.conditions[-1:].type`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Finished-TTL",type=integer,JSONPath=`.spec.ttlSecondsAfterFinished`
// +kubebuilder:printcolumn:name="Max-Lifetime",type=integer,JSONPath=`.spec.activeDeadlineSeconds`

// Represents a PyTorchJob resource.
type PyTorchJob struct {
	// Standard Kubernetes type metadata.
	metav1.TypeMeta `json:",inline"`

	// Standard Kubernetes object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired state of the PyTorchJob.
	Spec PyTorchJobSpec `json:"spec,omitempty"`

	// Most recently observed status of the PyTorchJob.
	// Read-only (modified by the system).
	Status common.JobStatus `json:"status,omitempty"`
}

// PyTorchJobSpec is a desired state description of the PyTorchJob.
type PyTorchJobSpec struct {
	// RunPolicy encapsulates various runtime policies of the distributed training
	// job, for example how to clean up resources and how long the job can stay
	// active.
	common.RunPolicy `json:",inline"`

	// A map of PyTorchReplicaType (type) to ReplicaSpec (value). Specifies the PyTorch cluster configuration.
	// For example,
	//   {
	//     "Master": PyTorchReplicaSpec,
	//     "Worker": PyTorchReplicaSpec,
	//   }
	PyTorchReplicaSpecs map[common.ReplicaType]*common.ReplicaSpec `json:"pytorchReplicaSpecs"`
}

const (
	// PyTorchReplicaTypeMaster is the type of Master of distributed PyTorch
	PyTorchReplicaTypeMaster common.ReplicaType = "Master"

	// PyTorchReplicaTypeWorker is the type for workers of distributed PyTorch.
	PyTorchReplicaTypeWorker common.ReplicaType = "Worker"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=pytorchjobs

// PyTorchJobList is a list of PyTorchJobs.
type PyTorchJobList struct {
	// Standard type metadata.
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	metav1.ListMeta `json:"metadata,omitempty"`

	// List of PyTorchJobs.
	Items []PyTorchJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PyTorchJob{}, &PyTorchJobList{})
}
