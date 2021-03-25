// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"github.com/vmware-tanzu/vm-operator-api/api/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ClassReference contains info to locate a Kind VirtualMachineClass object
type ClassReference struct {
	// API version of the referent.
	APIVersion string `json:"apiVersion,omitempty"`
	// Kind is the type of resource being referenced.
	Kind string `json:"kind,omitempty"`
	// Name is the name of resource being referenced.
	Name string `json:"name"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=vmclassbinding
// +kubebuilder:printcolumn:name="VirtualMachineClass",type="string",JSONPath=".classRef.name"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// VirtualMachineClassBinding is a binding object responsible for
// defining a VirtualMachineClass and a Namespace associated with it
type VirtualMachineClassBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// ClassReference is a reference to a VirtualMachineClass object
	ClassRef ClassReference `json:"classRef,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineClassBindingList contains a list of VirtualMachineClassBinding
type VirtualMachineClassBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineClassBinding `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&VirtualMachineClassBinding{}, &VirtualMachineClassBindingList{})
}

func (src *VirtualMachineClassBinding) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.VirtualMachineClassBinding)
	return Convert_v1alpha1_VirtualMachineClassBinding_To_v1alpha2_VirtualMachineClassBinding(src, dst, nil)
}

func (dst *VirtualMachineClassBinding) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.VirtualMachineClassBinding)
	return Convert_v1alpha2_VirtualMachineClassBinding_To_v1alpha1_VirtualMachineClassBinding(src, dst, nil)
}

func (src *VirtualMachineClassBindingList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.VirtualMachineClassBindingList)
	return Convert_v1alpha1_VirtualMachineClassBindingList_To_v1alpha2_VirtualMachineClassBindingList(src, dst, nil)
}

func (dst *VirtualMachineClassBindingList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.VirtualMachineClassBindingList)
	return Convert_v1alpha2_VirtualMachineClassBindingList_To_v1alpha1_VirtualMachineClassBindingList(src, dst, nil)
}
