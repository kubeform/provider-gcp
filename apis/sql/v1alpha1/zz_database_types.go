// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DatabaseInitParameters struct {

	// The charset value. See MySQL's
	// Supported Character Sets and Collations
	// and Postgres' Character Set Support
	// for more details and supported values. Postgres databases only support
	// a value of UTF8 at creation time.
	// The charset value. See MySQL's
	// [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
	// and Postgres' [Character Set Support](https://www.postgresql.org/docs/9.6/static/multibyte.html)
	// for more details and supported values. Postgres databases only support
	// a value of 'UTF8' at creation time.
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	// The collation value. See MySQL's
	// Supported Character Sets and Collations
	// and Postgres' Collation Support
	// for more details and supported values. Postgres databases only support
	// a value of en_US.UTF8 at creation time.
	// The collation value. See MySQL's
	// [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
	// and Postgres' [Collation Support](https://www.postgresql.org/docs/9.6/static/collation.html)
	// for more details and supported values. Postgres databases only support
	// a value of 'en_US.UTF8' at creation time.
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// The deletion policy for the database. Setting ABANDON allows the resource
	// to be abandoned rather than deleted. This is useful for Postgres, where databases cannot be
	// deleted from the API if there are users other than cloudsqlsuperuser with access. Possible
	// values are: "ABANDON", "DELETE". Defaults to "DELETE".
	// The deletion policy for the database. Setting ABANDON allows the resource
	// to be abandoned rather than deleted. This is useful for Postgres, where databases cannot be
	// deleted from the API if there are users other than cloudsqlsuperuser with access. Possible
	// values are: "ABANDON", "DELETE". Defaults to "DELETE".
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type DatabaseObservation struct {

	// The charset value. See MySQL's
	// Supported Character Sets and Collations
	// and Postgres' Character Set Support
	// for more details and supported values. Postgres databases only support
	// a value of UTF8 at creation time.
	// The charset value. See MySQL's
	// [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
	// and Postgres' [Character Set Support](https://www.postgresql.org/docs/9.6/static/multibyte.html)
	// for more details and supported values. Postgres databases only support
	// a value of 'UTF8' at creation time.
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	// The collation value. See MySQL's
	// Supported Character Sets and Collations
	// and Postgres' Collation Support
	// for more details and supported values. Postgres databases only support
	// a value of en_US.UTF8 at creation time.
	// The collation value. See MySQL's
	// [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
	// and Postgres' [Collation Support](https://www.postgresql.org/docs/9.6/static/collation.html)
	// for more details and supported values. Postgres databases only support
	// a value of 'en_US.UTF8' at creation time.
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// The deletion policy for the database. Setting ABANDON allows the resource
	// to be abandoned rather than deleted. This is useful for Postgres, where databases cannot be
	// deleted from the API if there are users other than cloudsqlsuperuser with access. Possible
	// values are: "ABANDON", "DELETE". Defaults to "DELETE".
	// The deletion policy for the database. Setting ABANDON allows the resource
	// to be abandoned rather than deleted. This is useful for Postgres, where databases cannot be
	// deleted from the API if there are users other than cloudsqlsuperuser with access. Possible
	// values are: "ABANDON", "DELETE". Defaults to "DELETE".
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// an identifier for the resource with format projects/{{project}}/instances/{{instance}}/databases/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Cloud SQL instance. This does not include the project
	// ID.
	// The name of the Cloud SQL instance. This does not include the project
	// ID.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type DatabaseParameters struct {

	// The charset value. See MySQL's
	// Supported Character Sets and Collations
	// and Postgres' Character Set Support
	// for more details and supported values. Postgres databases only support
	// a value of UTF8 at creation time.
	// The charset value. See MySQL's
	// [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
	// and Postgres' [Character Set Support](https://www.postgresql.org/docs/9.6/static/multibyte.html)
	// for more details and supported values. Postgres databases only support
	// a value of 'UTF8' at creation time.
	// +kubebuilder:validation:Optional
	Charset *string `json:"charset,omitempty" tf:"charset,omitempty"`

	// The collation value. See MySQL's
	// Supported Character Sets and Collations
	// and Postgres' Collation Support
	// for more details and supported values. Postgres databases only support
	// a value of en_US.UTF8 at creation time.
	// The collation value. See MySQL's
	// [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
	// and Postgres' [Collation Support](https://www.postgresql.org/docs/9.6/static/collation.html)
	// for more details and supported values. Postgres databases only support
	// a value of 'en_US.UTF8' at creation time.
	// +kubebuilder:validation:Optional
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// The deletion policy for the database. Setting ABANDON allows the resource
	// to be abandoned rather than deleted. This is useful for Postgres, where databases cannot be
	// deleted from the API if there are users other than cloudsqlsuperuser with access. Possible
	// values are: "ABANDON", "DELETE". Defaults to "DELETE".
	// The deletion policy for the database. Setting ABANDON allows the resource
	// to be abandoned rather than deleted. This is useful for Postgres, where databases cannot be
	// deleted from the API if there are users other than cloudsqlsuperuser with access. Possible
	// values are: "ABANDON", "DELETE". Defaults to "DELETE".
	// +kubebuilder:validation:Optional
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// The name of the Cloud SQL instance. This does not include the project
	// ID.
	// The name of the Cloud SQL instance. This does not include the project
	// ID.
	// +crossplane:generate:reference:type=DatabaseInstance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a DatabaseInstance to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a DatabaseInstance to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DatabaseInitParameters `json:"initProvider,omitempty"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Database is the Schema for the Databases API. Represents a SQL database inside the Cloud SQL instance, hosted in Google's cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	Database_Kind             = "Database"
	Database_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Database_Kind}.String()
	Database_KindAPIVersion   = Database_Kind + "." + CRDGroupVersion.String()
	Database_GroupVersionKind = CRDGroupVersion.WithKind(Database_Kind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
