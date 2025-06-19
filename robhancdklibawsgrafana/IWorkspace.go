package robhancdklibawsgrafana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-grafana/robhancdklibawsgrafana/internal"
)

// Represents an Amazon Managed Service for Grafana workspace.
type IWorkspace interface {
	awscdk.IResource
	// Specifies whether the workspace can access AWS resources in this AWS account only, or whether it can also access AWS resources in other accounts in the same organization.
	//
	// If this is
	// ORGANIZATION, the OrganizationalUnits parameter specifies which organizational units the
	// workspace can access.
	AccountAccessType() AccountAccessType
	// Specifies whether this workspace uses SAML 2.0, AWS IAM Identity Center, or both to authenticate users for using the Grafana console within a workspace.
	AuthenticationProviders() *[]AuthenticationProviders
	// A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.
	ClientToken() *string
	// Specifies the AWS data sources that have been configured to have IAM roles and permissions created to allow Amazon Managed Grafana to read data from these sources.
	//
	// This list is only used when the workspace was created through the AWS console, and the
	// permissionType is SERVICE_MANAGED.
	DataSources() *[]*string
	// The user-defined description of the workspace.
	Description() *string
	// The name of the workspace.
	Name() *string
	// The configuration settings for network access to your workspace.
	NetworkAccessControl() *NetworkAccessControl
	// The AWS notification channels that Amazon Managed Grafana can automatically create IAM roles and permissions for, to allow Amazon Managed Grafana to use these channels.
	NotificationDestinations() *[]NotificationDestinations
	// Specifies the organizational units that this workspace is allowed to use data sources from, if this workspace is in an account that is part of an organization.
	OrganizationalUnits() *[]*string
	// The name of the IAM role that is used to access resources through Organizations.
	OrganizationRoleName() *string
	// If this is SERVICE_MANAGED, and the workplace was created through the Amazon Managed Grafana console, then Amazon Managed Grafana automatically creates the IAM roles and provisions the permissions that the workspace needs to use AWS data sources and notification channels.
	//
	// If this is CUSTOMER_MANAGED, you must manage those roles and permissions yourself.
	//
	// If you are working with a workspace in a member account of an organization and that account is
	// not a delegated administrator account, and you want the workspace to access data sources in
	// other AWS accounts in the organization, this parameter must be set to CUSTOMER_MANAGED.
	PermissionType() PermissionTypes
	// Whether plugin administration is enabled in the workspace.
	//
	// Setting to true allows workspace
	// admins to install, uninstall, and update plugins from within the Grafana workspace.
	//
	// This option is only valid for workspaces that support Grafana version 9 or newer.
	PluginAdminEnabled() *bool
	// The IAM role that grants permissions to the AWS resources that the workspace will view data from.
	Role() awsiam.IRole
	// If the workspace uses SAML, use this structure to map SAML assertion attributes to workspace user information and define which groups in the assertion attribute are to have the Admin and Editor roles in the workspace.
	SamlConfiguration() *SamlConfiguration
	// The name of the AWS CloudFormation stack set that is used to generate IAM roles to be used for this workspace.
	StackSetName() *string
	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to.
	VpcConfiguration() *VpcConfiguration
	// The ARN of this workspace.
	WorkspaceArn() *string
	// The unique ID of this workspace.
	WorkspaceId() *string
}

// The jsii proxy for IWorkspace
type jsiiProxy_IWorkspace struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IWorkspace) AccountAccessType() AccountAccessType {
	var returns AccountAccessType
	_jsii_.Get(
		j,
		"accountAccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) AuthenticationProviders() *[]AuthenticationProviders {
	var returns *[]AuthenticationProviders
	_jsii_.Get(
		j,
		"authenticationProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) DataSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) NetworkAccessControl() *NetworkAccessControl {
	var returns *NetworkAccessControl
	_jsii_.Get(
		j,
		"networkAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) NotificationDestinations() *[]NotificationDestinations {
	var returns *[]NotificationDestinations
	_jsii_.Get(
		j,
		"notificationDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) OrganizationalUnits() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) OrganizationRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) PermissionType() PermissionTypes {
	var returns PermissionTypes
	_jsii_.Get(
		j,
		"permissionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) PluginAdminEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"pluginAdminEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) SamlConfiguration() *SamlConfiguration {
	var returns *SamlConfiguration
	_jsii_.Get(
		j,
		"samlConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) StackSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) VpcConfiguration() *VpcConfiguration {
	var returns *VpcConfiguration
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) WorkspaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

