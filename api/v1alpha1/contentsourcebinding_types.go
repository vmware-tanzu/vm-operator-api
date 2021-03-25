// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"github.com/vmware-tanzu/vm-operator-api/api/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ContentSourceReference contains info to locate a Kind ContentSource object.
type ContentSourceReference struct {
	// API version of the referent.
	APIVersion string `json:"apiVersion,omitempty"`
	// Kind is the type of resource being referenced.
	Kind string `json:"kind,omitempty"`
	// Name is the name of resource being referenced.
	Name string `json:"name"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="ContentSource",type="string",JSONPath=".contentSourceRef.name"

// ContentSourceBinding is an object that represents a ContentSource to Namespace mapping.
type ContentSourceBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// ContentSourceRef is a reference to a ContentSource object.
	ContentSourceRef ContentSourceReference `json:"contentSourceRef,omitempty"`
}

// +kubebuilder:object:root=true

// ContentSourceBindingList contains a list of ContentSourceBinding.
type ContentSourceBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentSourceBinding `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&ContentSourceBinding{}, &ContentSourceBindingList{})
}

func (src *ContentSourceBinding) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.ContentSourceBinding)
	return Convert_v1alpha1_ContentSourceBinding_To_v1alpha2_ContentSourceBinding(src, dst, nil)
}

func (dst *ContentSourceBinding) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.ContentSourceBinding)
	return Convert_v1alpha2_ContentSourceBinding_To_v1alpha1_ContentSourceBinding(src, dst, nil)
}

func (src *ContentSourceBindingList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.ContentSourceBindingList)
	return Convert_v1alpha1_ContentSourceBindingList_To_v1alpha2_ContentSourceBindingList(src, dst, nil)
}

func (dst *ContentSourceBindingList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.ContentSourceBindingList)
	return Convert_v1alpha2_ContentSourceBindingList_To_v1alpha1_ContentSourceBindingList(src, dst, nil)
}
