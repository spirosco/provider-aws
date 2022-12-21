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

type VPCPeeringConnectionOptionsAccepterObservation struct {
}

type VPCPeeringConnectionOptionsAccepterParameters struct {

	// Allow a local linked EC2-Classic instance to communicate
	// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
	// to the remote VPC. This option is not supported for inter-region VPC peering.
	// +kubebuilder:validation:Optional
	AllowClassicLinkToRemoteVPC *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	// +kubebuilder:validation:Optional
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// Allow a local VPC to communicate with a linked EC2-Classic
	// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
	// connection. This option is not supported for inter-region VPC peering.
	// +kubebuilder:validation:Optional
	AllowVPCToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VPCPeeringConnectionOptionsObservation struct {

	// The ID of the VPC Peering Connection Options.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPCPeeringConnectionOptionsParameters struct {

	// An optional configuration block that allows for [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	// +kubebuilder:validation:Optional
	Accepter []VPCPeeringConnectionOptionsAccepterParameters `json:"accepter,omitempty" tf:"accepter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A optional configuration block that allows for [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	// +kubebuilder:validation:Optional
	Requester []VPCPeeringConnectionOptionsRequesterParameters `json:"requester,omitempty" tf:"requester,omitempty"`

	// The ID of the requester VPC peering connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCPeeringConnection
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`

	// Reference to a VPCPeeringConnection in ec2 to populate vpcPeeringConnectionId.
	// +kubebuilder:validation:Optional
	VPCPeeringConnectionIDRef *v1.Reference `json:"vpcPeeringConnectionIdRef,omitempty" tf:"-"`

	// Selector for a VPCPeeringConnection in ec2 to populate vpcPeeringConnectionId.
	// +kubebuilder:validation:Optional
	VPCPeeringConnectionIDSelector *v1.Selector `json:"vpcPeeringConnectionIdSelector,omitempty" tf:"-"`
}

type VPCPeeringConnectionOptionsRequesterObservation struct {
}

type VPCPeeringConnectionOptionsRequesterParameters struct {

	// Allow a local linked EC2-Classic instance to communicate
	// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
	// to the remote VPC. This option is not supported for inter-region VPC peering.
	// +kubebuilder:validation:Optional
	AllowClassicLinkToRemoteVPC *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`

	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	// +kubebuilder:validation:Optional
	AllowRemoteVPCDNSResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`

	// Allow a local VPC to communicate with a linked EC2-Classic
	// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
	// connection. This option is not supported for inter-region VPC peering.
	// +kubebuilder:validation:Optional
	AllowVPCToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

// VPCPeeringConnectionOptionsSpec defines the desired state of VPCPeeringConnectionOptions
type VPCPeeringConnectionOptionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCPeeringConnectionOptionsParameters `json:"forProvider"`
}

// VPCPeeringConnectionOptionsStatus defines the observed state of VPCPeeringConnectionOptions.
type VPCPeeringConnectionOptionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCPeeringConnectionOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCPeeringConnectionOptions is the Schema for the VPCPeeringConnectionOptionss API. Provides a resource to manage VPC peering connection options.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCPeeringConnectionOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCPeeringConnectionOptionsSpec   `json:"spec"`
	Status            VPCPeeringConnectionOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCPeeringConnectionOptionsList contains a list of VPCPeeringConnectionOptionss
type VPCPeeringConnectionOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCPeeringConnectionOptions `json:"items"`
}

// Repository type metadata.
var (
	VPCPeeringConnectionOptions_Kind             = "VPCPeeringConnectionOptions"
	VPCPeeringConnectionOptions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCPeeringConnectionOptions_Kind}.String()
	VPCPeeringConnectionOptions_KindAPIVersion   = VPCPeeringConnectionOptions_Kind + "." + CRDGroupVersion.String()
	VPCPeeringConnectionOptions_GroupVersionKind = CRDGroupVersion.WithKind(VPCPeeringConnectionOptions_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCPeeringConnectionOptions{}, &VPCPeeringConnectionOptionsList{})
}