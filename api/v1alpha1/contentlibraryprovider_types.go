// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"github.com/vmware-tanzu/vm-operator-api/api/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ContentLibraryProviderSpec defines the desired state of ContentLibraryProvider
type ContentLibraryProviderSpec struct {
	// UUID describes the UUID of a vSphere content library. It is the unique identifier for a
	// vSphere content library.
	UUID string `json:"uuid,omitempty"`
}

// ContentLibraryProviderStatus defines the observed state of ContentLibraryProvider
// Can include fields indicating when was the last time VM images were updated from a library
type ContentLibraryProviderStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Content Library UUID",type="string",JSONPath=".spec.uuid",description="UUID of the vSphere content library"

// ContentLibraryProvider is the Schema for the contentlibraryproviders API
type ContentLibraryProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContentLibraryProviderSpec   `json:"spec,omitempty"`
	Status ContentLibraryProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContentLibraryProviderList contains a list of ContentLibraryProvider
type ContentLibraryProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentLibraryProvider `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&ContentLibraryProvider{}, &ContentLibraryProviderList{})
}

func (src *ContentLibraryProvider) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.ContentLibraryProvider)
	return Convert_v1alpha1_ContentLibraryProvider_To_v1alpha2_ContentLibraryProvider(src, dst, nil)
}

func (dst *ContentLibraryProvider) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.ContentLibraryProvider)
	return Convert_v1alpha2_ContentLibraryProvider_To_v1alpha1_ContentLibraryProvider(src, dst, nil)
}

func (src *ContentLibraryProviderList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.ContentLibraryProviderList)
	return Convert_v1alpha1_ContentLibraryProviderList_To_v1alpha2_ContentLibraryProviderList(src, dst, nil)
}

func (dst *ContentLibraryProviderList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.ContentLibraryProviderList)
	return Convert_v1alpha2_ContentLibraryProviderList_To_v1alpha1_ContentLibraryProviderList(src, dst, nil)
}
