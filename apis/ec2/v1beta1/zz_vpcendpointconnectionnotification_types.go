/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VPCEndpointConnectionNotificationObservation struct {

	// One or more endpoint events for which to receive notifications.
	ConnectionEvents []*string `json:"connectionEvents,omitempty" tf:"connection_events,omitempty"`

	// The ARN of the SNS topic for the notifications.
	ConnectionNotificationArn *string `json:"connectionNotificationArn,omitempty" tf:"connection_notification_arn,omitempty"`

	// The ID of the VPC connection notification.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of notification.
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`

	// The state of the notification.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The ID of the VPC Endpoint to receive notifications for.
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The ID of the VPC Endpoint Service to receive notifications for.
	VPCEndpointServiceID *string `json:"vpcEndpointServiceId,omitempty" tf:"vpc_endpoint_service_id,omitempty"`
}

type VPCEndpointConnectionNotificationParameters struct {

	// One or more endpoint events for which to receive notifications.
	// +kubebuilder:validation:Optional
	ConnectionEvents []*string `json:"connectionEvents,omitempty" tf:"connection_events,omitempty"`

	// The ARN of the SNS topic for the notifications.
	// +kubebuilder:validation:Optional
	ConnectionNotificationArn *string `json:"connectionNotificationArn,omitempty" tf:"connection_notification_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the VPC Endpoint to receive notifications for.
	// +kubebuilder:validation:Optional
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The ID of the VPC Endpoint Service to receive notifications for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCEndpointService
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCEndpointServiceID *string `json:"vpcEndpointServiceId,omitempty" tf:"vpc_endpoint_service_id,omitempty"`

	// Reference to a VPCEndpointService in ec2 to populate vpcEndpointServiceId.
	// +kubebuilder:validation:Optional
	VPCEndpointServiceIDRef *v1.Reference `json:"vpcEndpointServiceIdRef,omitempty" tf:"-"`

	// Selector for a VPCEndpointService in ec2 to populate vpcEndpointServiceId.
	// +kubebuilder:validation:Optional
	VPCEndpointServiceIDSelector *v1.Selector `json:"vpcEndpointServiceIdSelector,omitempty" tf:"-"`
}

// VPCEndpointConnectionNotificationSpec defines the desired state of VPCEndpointConnectionNotification
type VPCEndpointConnectionNotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointConnectionNotificationParameters `json:"forProvider"`
}

// VPCEndpointConnectionNotificationStatus defines the observed state of VPCEndpointConnectionNotification.
type VPCEndpointConnectionNotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointConnectionNotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointConnectionNotification is the Schema for the VPCEndpointConnectionNotifications API. Provides a VPC Endpoint connection notification resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCEndpointConnectionNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.connectionEvents)",message="connectionEvents is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.connectionNotificationArn)",message="connectionNotificationArn is a required parameter"
	Spec   VPCEndpointConnectionNotificationSpec   `json:"spec"`
	Status VPCEndpointConnectionNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointConnectionNotificationList contains a list of VPCEndpointConnectionNotifications
type VPCEndpointConnectionNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpointConnectionNotification `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpointConnectionNotification_Kind             = "VPCEndpointConnectionNotification"
	VPCEndpointConnectionNotification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpointConnectionNotification_Kind}.String()
	VPCEndpointConnectionNotification_KindAPIVersion   = VPCEndpointConnectionNotification_Kind + "." + CRDGroupVersion.String()
	VPCEndpointConnectionNotification_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpointConnectionNotification_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpointConnectionNotification{}, &VPCEndpointConnectionNotificationList{})
}
