/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	launchconfiguration "github.com/upbound/provider-aws/internal/controller/autoscaling/launchconfiguration"
	app "github.com/upbound/provider-aws/internal/controller/deploy/app"
	deploymentconfig "github.com/upbound/provider-aws/internal/controller/deploy/deploymentconfig"
	deploymentgroup "github.com/upbound/provider-aws/internal/controller/deploy/deploymentgroup"
	ami "github.com/upbound/provider-aws/internal/controller/ec2/ami"
	amicopy "github.com/upbound/provider-aws/internal/controller/ec2/amicopy"
	amilaunchpermission "github.com/upbound/provider-aws/internal/controller/ec2/amilaunchpermission"
	availabilityzonegroup "github.com/upbound/provider-aws/internal/controller/ec2/availabilityzonegroup"
	capacityreservation "github.com/upbound/provider-aws/internal/controller/ec2/capacityreservation"
	carriergateway "github.com/upbound/provider-aws/internal/controller/ec2/carriergateway"
	customergateway "github.com/upbound/provider-aws/internal/controller/ec2/customergateway"
	defaultnetworkacl "github.com/upbound/provider-aws/internal/controller/ec2/defaultnetworkacl"
	defaultroutetable "github.com/upbound/provider-aws/internal/controller/ec2/defaultroutetable"
	defaultsecuritygroup "github.com/upbound/provider-aws/internal/controller/ec2/defaultsecuritygroup"
	defaultsubnet "github.com/upbound/provider-aws/internal/controller/ec2/defaultsubnet"
	defaultvpc "github.com/upbound/provider-aws/internal/controller/ec2/defaultvpc"
	defaultvpcdhcpoptions "github.com/upbound/provider-aws/internal/controller/ec2/defaultvpcdhcpoptions"
	ebsdefaultkmskey "github.com/upbound/provider-aws/internal/controller/ec2/ebsdefaultkmskey"
	ebsencryptionbydefault "github.com/upbound/provider-aws/internal/controller/ec2/ebsencryptionbydefault"
	ebssnapshot "github.com/upbound/provider-aws/internal/controller/ec2/ebssnapshot"
	ebssnapshotcopy "github.com/upbound/provider-aws/internal/controller/ec2/ebssnapshotcopy"
	ebssnapshotimport "github.com/upbound/provider-aws/internal/controller/ec2/ebssnapshotimport"
	ebsvolume "github.com/upbound/provider-aws/internal/controller/ec2/ebsvolume"
	egressonlyinternetgateway "github.com/upbound/provider-aws/internal/controller/ec2/egressonlyinternetgateway"
	eip "github.com/upbound/provider-aws/internal/controller/ec2/eip"
	eipassociation "github.com/upbound/provider-aws/internal/controller/ec2/eipassociation"
	flowlog "github.com/upbound/provider-aws/internal/controller/ec2/flowlog"
	host "github.com/upbound/provider-aws/internal/controller/ec2/host"
	instance "github.com/upbound/provider-aws/internal/controller/ec2/instance"
	instancestate "github.com/upbound/provider-aws/internal/controller/ec2/instancestate"
	internetgateway "github.com/upbound/provider-aws/internal/controller/ec2/internetgateway"
	keypair "github.com/upbound/provider-aws/internal/controller/ec2/keypair"
	launchtemplate "github.com/upbound/provider-aws/internal/controller/ec2/launchtemplate"
	mainroutetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/mainroutetableassociation"
	managedprefixlist "github.com/upbound/provider-aws/internal/controller/ec2/managedprefixlist"
	managedprefixlistentry "github.com/upbound/provider-aws/internal/controller/ec2/managedprefixlistentry"
	natgateway "github.com/upbound/provider-aws/internal/controller/ec2/natgateway"
	networkacl "github.com/upbound/provider-aws/internal/controller/ec2/networkacl"
	networkaclrule "github.com/upbound/provider-aws/internal/controller/ec2/networkaclrule"
	networkinsightsanalysis "github.com/upbound/provider-aws/internal/controller/ec2/networkinsightsanalysis"
	networkinsightspath "github.com/upbound/provider-aws/internal/controller/ec2/networkinsightspath"
	networkinterface "github.com/upbound/provider-aws/internal/controller/ec2/networkinterface"
	networkinterfaceattachment "github.com/upbound/provider-aws/internal/controller/ec2/networkinterfaceattachment"
	networkinterfacesgattachment "github.com/upbound/provider-aws/internal/controller/ec2/networkinterfacesgattachment"
	placementgroup "github.com/upbound/provider-aws/internal/controller/ec2/placementgroup"
	route "github.com/upbound/provider-aws/internal/controller/ec2/route"
	routetable "github.com/upbound/provider-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/upbound/provider-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/upbound/provider-aws/internal/controller/ec2/securitygrouprule"
	serialconsoleaccess "github.com/upbound/provider-aws/internal/controller/ec2/serialconsoleaccess"
	snapshotcreatevolumepermission "github.com/upbound/provider-aws/internal/controller/ec2/snapshotcreatevolumepermission"
	spotdatafeedsubscription "github.com/upbound/provider-aws/internal/controller/ec2/spotdatafeedsubscription"
	spotfleetrequest "github.com/upbound/provider-aws/internal/controller/ec2/spotfleetrequest"
	spotinstancerequest "github.com/upbound/provider-aws/internal/controller/ec2/spotinstancerequest"
	subnet "github.com/upbound/provider-aws/internal/controller/ec2/subnet"
	subnetcidrreservation "github.com/upbound/provider-aws/internal/controller/ec2/subnetcidrreservation"
	tag "github.com/upbound/provider-aws/internal/controller/ec2/tag"
	trafficmirrorfilter "github.com/upbound/provider-aws/internal/controller/ec2/trafficmirrorfilter"
	trafficmirrorfilterrule "github.com/upbound/provider-aws/internal/controller/ec2/trafficmirrorfilterrule"
	transitgateway "github.com/upbound/provider-aws/internal/controller/ec2/transitgateway"
	transitgatewayconnect "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayconnect"
	transitgatewayconnectpeer "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayconnectpeer"
	transitgatewaymulticastdomain "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastdomain"
	transitgatewaymulticastdomainassociation "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastdomainassociation"
	transitgatewaymulticastgroupmember "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastgroupmember"
	transitgatewaymulticastgroupsource "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastgroupsource"
	transitgatewaypeeringattachment "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaypeeringattachment"
	transitgatewaypeeringattachmentaccepter "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaypeeringattachmentaccepter"
	transitgatewaypolicytable "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaypolicytable"
	transitgatewayprefixlistreference "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayprefixlistreference"
	transitgatewayroute "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	volumeattachment "github.com/upbound/provider-aws/internal/controller/ec2/volumeattachment"
	vpc "github.com/upbound/provider-aws/internal/controller/ec2/vpc"
	vpcdhcpoptions "github.com/upbound/provider-aws/internal/controller/ec2/vpcdhcpoptions"
	vpcdhcpoptionsassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcdhcpoptionsassociation"
	vpcendpoint "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpoint"
	vpcendpointconnectionnotification "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointconnectionnotification"
	vpcendpointroutetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointroutetableassociation"
	vpcendpointsecuritygroupassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointsecuritygroupassociation"
	vpcendpointservice "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointservice"
	vpcendpointserviceallowedprincipal "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointserviceallowedprincipal"
	vpcendpointsubnetassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointsubnetassociation"
	vpcipam "github.com/upbound/provider-aws/internal/controller/ec2/vpcipam"
	vpcipampool "github.com/upbound/provider-aws/internal/controller/ec2/vpcipampool"
	vpcipampoolcidr "github.com/upbound/provider-aws/internal/controller/ec2/vpcipampoolcidr"
	vpcipampoolcidrallocation "github.com/upbound/provider-aws/internal/controller/ec2/vpcipampoolcidrallocation"
	vpcipamscope "github.com/upbound/provider-aws/internal/controller/ec2/vpcipamscope"
	vpcipv4cidrblockassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/upbound/provider-aws/internal/controller/ec2/vpcpeeringconnection"
	vpcpeeringconnectionaccepter "github.com/upbound/provider-aws/internal/controller/ec2/vpcpeeringconnectionaccepter"
	vpcpeeringconnectionoptions "github.com/upbound/provider-aws/internal/controller/ec2/vpcpeeringconnectionoptions"
	vpnconnection "github.com/upbound/provider-aws/internal/controller/ec2/vpnconnection"
	vpnconnectionroute "github.com/upbound/provider-aws/internal/controller/ec2/vpnconnectionroute"
	vpngateway "github.com/upbound/provider-aws/internal/controller/ec2/vpngateway"
	vpngatewayattachment "github.com/upbound/provider-aws/internal/controller/ec2/vpngatewayattachment"
	vpngatewayroutepropagation "github.com/upbound/provider-aws/internal/controller/ec2/vpngatewayroutepropagation"
	accesskey "github.com/upbound/provider-aws/internal/controller/iam/accesskey"
	accountalias "github.com/upbound/provider-aws/internal/controller/iam/accountalias"
	accountpasswordpolicy "github.com/upbound/provider-aws/internal/controller/iam/accountpasswordpolicy"
	group "github.com/upbound/provider-aws/internal/controller/iam/group"
	groupmembership "github.com/upbound/provider-aws/internal/controller/iam/groupmembership"
	grouppolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofile "github.com/upbound/provider-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/upbound/provider-aws/internal/controller/iam/openidconnectprovider"
	policy "github.com/upbound/provider-aws/internal/controller/iam/policy"
	role "github.com/upbound/provider-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/rolepolicyattachment"
	samlprovider "github.com/upbound/provider-aws/internal/controller/iam/samlprovider"
	servercertificate "github.com/upbound/provider-aws/internal/controller/iam/servercertificate"
	servicelinkedrole "github.com/upbound/provider-aws/internal/controller/iam/servicelinkedrole"
	servicespecificcredential "github.com/upbound/provider-aws/internal/controller/iam/servicespecificcredential"
	signingcertificate "github.com/upbound/provider-aws/internal/controller/iam/signingcertificate"
	user "github.com/upbound/provider-aws/internal/controller/iam/user"
	usergroupmembership "github.com/upbound/provider-aws/internal/controller/iam/usergroupmembership"
	userloginprofile "github.com/upbound/provider-aws/internal/controller/iam/userloginprofile"
	userpolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/userpolicyattachment"
	usersshkey "github.com/upbound/provider-aws/internal/controller/iam/usersshkey"
	virtualmfadevice "github.com/upbound/provider-aws/internal/controller/iam/virtualmfadevice"
	alias "github.com/upbound/provider-aws/internal/controller/kms/alias"
	ciphertext "github.com/upbound/provider-aws/internal/controller/kms/ciphertext"
	externalkey "github.com/upbound/provider-aws/internal/controller/kms/externalkey"
	grant "github.com/upbound/provider-aws/internal/controller/kms/grant"
	key "github.com/upbound/provider-aws/internal/controller/kms/key"
	replicaexternalkey "github.com/upbound/provider-aws/internal/controller/kms/replicaexternalkey"
	replicakey "github.com/upbound/provider-aws/internal/controller/kms/replicakey"
	providerconfig "github.com/upbound/provider-aws/internal/controller/providerconfig"
	cluster "github.com/upbound/provider-aws/internal/controller/rds/cluster"
	clusteractivitystream "github.com/upbound/provider-aws/internal/controller/rds/clusteractivitystream"
	clusterendpoint "github.com/upbound/provider-aws/internal/controller/rds/clusterendpoint"
	clusterinstance "github.com/upbound/provider-aws/internal/controller/rds/clusterinstance"
	clusterparametergroup "github.com/upbound/provider-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "github.com/upbound/provider-aws/internal/controller/rds/clusterroleassociation"
	clustersnapshot "github.com/upbound/provider-aws/internal/controller/rds/clustersnapshot"
	dbinstanceautomatedbackupsreplication "github.com/upbound/provider-aws/internal/controller/rds/dbinstanceautomatedbackupsreplication"
	dbsnapshotcopy "github.com/upbound/provider-aws/internal/controller/rds/dbsnapshotcopy"
	eventsubscription "github.com/upbound/provider-aws/internal/controller/rds/eventsubscription"
	globalcluster "github.com/upbound/provider-aws/internal/controller/rds/globalcluster"
	instancerds "github.com/upbound/provider-aws/internal/controller/rds/instance"
	instanceroleassociation "github.com/upbound/provider-aws/internal/controller/rds/instanceroleassociation"
	optiongroup "github.com/upbound/provider-aws/internal/controller/rds/optiongroup"
	parametergroup "github.com/upbound/provider-aws/internal/controller/rds/parametergroup"
	proxy "github.com/upbound/provider-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "github.com/upbound/provider-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/upbound/provider-aws/internal/controller/rds/proxyendpoint"
	proxytarget "github.com/upbound/provider-aws/internal/controller/rds/proxytarget"
	securitygrouprds "github.com/upbound/provider-aws/internal/controller/rds/securitygroup"
	snapshot "github.com/upbound/provider-aws/internal/controller/rds/snapshot"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/rds/subnetgroup"
	groupresourcegroups "github.com/upbound/provider-aws/internal/controller/resourcegroups/group"
	bucket "github.com/upbound/provider-aws/internal/controller/s3/bucket"
	bucketaccelerateconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketaccelerateconfiguration"
	bucketacl "github.com/upbound/provider-aws/internal/controller/s3/bucketacl"
	bucketanalyticsconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketanalyticsconfiguration"
	bucketcorsconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketcorsconfiguration"
	bucketintelligenttieringconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketintelligenttieringconfiguration"
	bucketinventory "github.com/upbound/provider-aws/internal/controller/s3/bucketinventory"
	bucketlifecycleconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketlifecycleconfiguration"
	bucketlogging "github.com/upbound/provider-aws/internal/controller/s3/bucketlogging"
	bucketmetric "github.com/upbound/provider-aws/internal/controller/s3/bucketmetric"
	bucketnotification "github.com/upbound/provider-aws/internal/controller/s3/bucketnotification"
	bucketobject "github.com/upbound/provider-aws/internal/controller/s3/bucketobject"
	bucketobjectlockconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketobjectlockconfiguration"
	bucketownershipcontrols "github.com/upbound/provider-aws/internal/controller/s3/bucketownershipcontrols"
	bucketpolicy "github.com/upbound/provider-aws/internal/controller/s3/bucketpolicy"
	bucketpublicaccessblock "github.com/upbound/provider-aws/internal/controller/s3/bucketpublicaccessblock"
	bucketreplicationconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketreplicationconfiguration"
	bucketrequestpaymentconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketrequestpaymentconfiguration"
	bucketserversideencryptionconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/upbound/provider-aws/internal/controller/s3/bucketversioning"
	bucketwebsiteconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketwebsiteconfiguration"
	object "github.com/upbound/provider-aws/internal/controller/s3/object"
	objectcopy "github.com/upbound/provider-aws/internal/controller/s3/objectcopy"
	accesspoint "github.com/upbound/provider-aws/internal/controller/s3control/accesspoint"
	accesspointpolicy "github.com/upbound/provider-aws/internal/controller/s3control/accesspointpolicy"
	accountpublicaccessblock "github.com/upbound/provider-aws/internal/controller/s3control/accountpublicaccessblock"
	multiregionaccesspoint "github.com/upbound/provider-aws/internal/controller/s3control/multiregionaccesspoint"
	multiregionaccesspointpolicy "github.com/upbound/provider-aws/internal/controller/s3control/multiregionaccesspointpolicy"
	objectlambdaaccesspoint "github.com/upbound/provider-aws/internal/controller/s3control/objectlambdaaccesspoint"
	objectlambdaaccesspointpolicy "github.com/upbound/provider-aws/internal/controller/s3control/objectlambdaaccesspointpolicy"
	storagelensconfiguration "github.com/upbound/provider-aws/internal/controller/s3control/storagelensconfiguration"
	networkperformancemetricsubscription "github.com/upbound/provider-aws/internal/controller/vpc/networkperformancemetricsubscription"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		launchconfiguration.Setup,
		app.Setup,
		deploymentconfig.Setup,
		deploymentgroup.Setup,
		ami.Setup,
		amicopy.Setup,
		amilaunchpermission.Setup,
		availabilityzonegroup.Setup,
		capacityreservation.Setup,
		carriergateway.Setup,
		customergateway.Setup,
		defaultnetworkacl.Setup,
		defaultroutetable.Setup,
		defaultsecuritygroup.Setup,
		defaultsubnet.Setup,
		defaultvpc.Setup,
		defaultvpcdhcpoptions.Setup,
		ebsdefaultkmskey.Setup,
		ebsencryptionbydefault.Setup,
		ebssnapshot.Setup,
		ebssnapshotcopy.Setup,
		ebssnapshotimport.Setup,
		ebsvolume.Setup,
		egressonlyinternetgateway.Setup,
		eip.Setup,
		eipassociation.Setup,
		flowlog.Setup,
		host.Setup,
		instance.Setup,
		instancestate.Setup,
		internetgateway.Setup,
		keypair.Setup,
		launchtemplate.Setup,
		mainroutetableassociation.Setup,
		managedprefixlist.Setup,
		managedprefixlistentry.Setup,
		natgateway.Setup,
		networkacl.Setup,
		networkaclrule.Setup,
		networkinsightsanalysis.Setup,
		networkinsightspath.Setup,
		networkinterface.Setup,
		networkinterfaceattachment.Setup,
		networkinterfacesgattachment.Setup,
		placementgroup.Setup,
		route.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		serialconsoleaccess.Setup,
		snapshotcreatevolumepermission.Setup,
		spotdatafeedsubscription.Setup,
		spotfleetrequest.Setup,
		spotinstancerequest.Setup,
		subnet.Setup,
		subnetcidrreservation.Setup,
		tag.Setup,
		trafficmirrorfilter.Setup,
		trafficmirrorfilterrule.Setup,
		transitgateway.Setup,
		transitgatewayconnect.Setup,
		transitgatewayconnectpeer.Setup,
		transitgatewaymulticastdomain.Setup,
		transitgatewaymulticastdomainassociation.Setup,
		transitgatewaymulticastgroupmember.Setup,
		transitgatewaymulticastgroupsource.Setup,
		transitgatewaypeeringattachment.Setup,
		transitgatewaypeeringattachmentaccepter.Setup,
		transitgatewaypolicytable.Setup,
		transitgatewayprefixlistreference.Setup,
		transitgatewayroute.Setup,
		transitgatewayroutetable.Setup,
		transitgatewayroutetableassociation.Setup,
		transitgatewayroutetablepropagation.Setup,
		transitgatewayvpcattachment.Setup,
		transitgatewayvpcattachmentaccepter.Setup,
		volumeattachment.Setup,
		vpc.Setup,
		vpcdhcpoptions.Setup,
		vpcdhcpoptionsassociation.Setup,
		vpcendpoint.Setup,
		vpcendpointconnectionnotification.Setup,
		vpcendpointroutetableassociation.Setup,
		vpcendpointsecuritygroupassociation.Setup,
		vpcendpointservice.Setup,
		vpcendpointserviceallowedprincipal.Setup,
		vpcendpointsubnetassociation.Setup,
		vpcipam.Setup,
		vpcipampool.Setup,
		vpcipampoolcidr.Setup,
		vpcipampoolcidrallocation.Setup,
		vpcipamscope.Setup,
		vpcipv4cidrblockassociation.Setup,
		vpcpeeringconnection.Setup,
		vpcpeeringconnectionaccepter.Setup,
		vpcpeeringconnectionoptions.Setup,
		vpnconnection.Setup,
		vpnconnectionroute.Setup,
		vpngateway.Setup,
		vpngatewayattachment.Setup,
		vpngatewayroutepropagation.Setup,
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		group.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		instanceprofile.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		servercertificate.Setup,
		servicelinkedrole.Setup,
		servicespecificcredential.Setup,
		signingcertificate.Setup,
		user.Setup,
		usergroupmembership.Setup,
		userloginprofile.Setup,
		userpolicyattachment.Setup,
		usersshkey.Setup,
		virtualmfadevice.Setup,
		alias.Setup,
		ciphertext.Setup,
		externalkey.Setup,
		grant.Setup,
		key.Setup,
		replicaexternalkey.Setup,
		replicakey.Setup,
		providerconfig.Setup,
		cluster.Setup,
		clusteractivitystream.Setup,
		clusterendpoint.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clusterroleassociation.Setup,
		clustersnapshot.Setup,
		dbinstanceautomatedbackupsreplication.Setup,
		dbsnapshotcopy.Setup,
		eventsubscription.Setup,
		globalcluster.Setup,
		instancerds.Setup,
		instanceroleassociation.Setup,
		optiongroup.Setup,
		parametergroup.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		securitygrouprds.Setup,
		snapshot.Setup,
		subnetgroup.Setup,
		groupresourcegroups.Setup,
		bucket.Setup,
		bucketaccelerateconfiguration.Setup,
		bucketacl.Setup,
		bucketanalyticsconfiguration.Setup,
		bucketcorsconfiguration.Setup,
		bucketintelligenttieringconfiguration.Setup,
		bucketinventory.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketlogging.Setup,
		bucketmetric.Setup,
		bucketnotification.Setup,
		bucketobject.Setup,
		bucketobjectlockconfiguration.Setup,
		bucketownershipcontrols.Setup,
		bucketpolicy.Setup,
		bucketpublicaccessblock.Setup,
		bucketreplicationconfiguration.Setup,
		bucketrequestpaymentconfiguration.Setup,
		bucketserversideencryptionconfiguration.Setup,
		bucketversioning.Setup,
		bucketwebsiteconfiguration.Setup,
		object.Setup,
		objectcopy.Setup,
		accesspoint.Setup,
		accesspointpolicy.Setup,
		accountpublicaccessblock.Setup,
		multiregionaccesspoint.Setup,
		multiregionaccesspointpolicy.Setup,
		objectlambdaaccesspoint.Setup,
		objectlambdaaccesspointpolicy.Setup,
		storagelensconfiguration.Setup,
		networkperformancemetricsubscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
