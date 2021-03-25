// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"github.com/vmware-tanzu/vm-operator-api/api/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ContentProviderInterface contains the info to locate a content provider resource
type ContentProviderReference struct {
	// API version of the referent.
	APIVersion string `json:"apiVersion,omitempty"`
	// Kind is the type of resource being referenced.
	Kind string `json:"kind"`
	// Name is the name of resource being referenced.
	Name string `json:"name"`
	// Namespace of the resource being referenced. If empty, cluster scoped resource is assumed.
	Namespace string `json:"namespace,omitempty"`
}

// ContentSourceSpec defines the desired state of ContentSource
type ContentSourceSpec struct {
	// ProviderRef is a reference to a content provider object that describes a provider.
	ProviderRef ContentProviderReference `json:"providerRef,omitempty"`
}

// ContentSourceStatus defines the observed state of ContentSource
type ContentSourceStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster

// ContentSource is the Schema for the contentsources API.
// A ContentSource represents the desired specification and the observed status of a ContentSource instance.
type ContentSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContentSourceSpec   `json:"spec,omitempty"`
	Status ContentSourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContentSourceList contains a list of ContentSource
type ContentSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentSource `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&ContentSource{}, &ContentSourceList{})
}

func (src *ContentSource) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.ContentSource)
	return Convert_v1alpha1_ContentSource_To_v1alpha2_ContentSource(src, dst, nil)
}

func (dst *ContentSource) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.ContentSource)
	return Convert_v1alpha2_ContentSource_To_v1alpha1_ContentSource(src, dst, nil)
}

func (src *ContentSourceList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.ContentSourceList)
	return Convert_v1alpha1_ContentSourceList_To_v1alpha2_ContentSourceList(src, dst, nil)
}

func (dst *ContentSourceList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.ContentSourceList)
	return Convert_v1alpha2_ContentSourceList_To_v1alpha1_ContentSourceList(src, dst, nil)
}
