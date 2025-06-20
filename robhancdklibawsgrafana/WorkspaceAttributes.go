package robhancdklibawsgrafana

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

type WorkspaceAttributes struct {
	// Specifies whether the workspace can access AWS resources in this AWS account only, or whether it can also access AWS resources in other accounts in the same organization.
	//
	// If this is
	// ORGANIZATION, the OrganizationalUnits parameter specifies which organizational units the
	// workspace can access.
	//
	// Required field.
	AccountAccessType AccountAccessType `field:"required" json:"accountAccessType" yaml:"accountAccessType"`
	// Specifies whether this workspace uses SAML 2.0, AWS IAM Identity Center, or both to authenticate users for using the Grafana console within a workspace.
	//
	// Required field.
	AuthenticationProviders *[]AuthenticationProviders `field:"required" json:"authenticationProviders" yaml:"authenticationProviders"`
	// If this is SERVICE_MANAGED, and the workplace was created through the Amazon Managed Grafana console, then Amazon Managed Grafana automatically creates the IAM roles and provisions the permissions that the workspace needs to use AWS data sources and notification channels.
	//
	// If this is CUSTOMER_MANAGED, you must manage those roles and permissions yourself.
	//
	// If you are working with a workspace in a member account of an organization and that account is
	// not a delegated administrator account, and you want the workspace to access data sources in
	// other AWS accounts in the organization, this parameter must be set to CUSTOMER_MANAGED.
	//
	// Required field.
	PermissionType PermissionTypes `field:"required" json:"permissionType" yaml:"permissionType"`
	// The arn of this workspace.
	WorkspaceArn *string `field:"required" json:"workspaceArn" yaml:"workspaceArn"`
	// A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.
	//
	// Must be 1-64 characters long and contain only printable ASCII characters.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Specifies the AWS data sources that have been configured to have IAM roles and permissions created to allow Amazon Managed Grafana to read data from these sources.
	//
	// This list is only used when the workspace was created through the AWS console, and the
	// permissionType is SERVICE_MANAGED.
	DataSources *[]*string `field:"optional" json:"dataSources" yaml:"dataSources"`
	// The user-defined description of the workspace.
	//
	// Maximum length of 2048 characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the workspace.
	//
	// Must be 1-255 characters long and contain only alphanumeric characters, hyphens, dots,
	// underscores, and tildes.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The configuration settings for network access to your workspace.
	NetworkAccessControl *NetworkAccessControl `field:"optional" json:"networkAccessControl" yaml:"networkAccessControl"`
	// The AWS notification channels that Amazon Managed Grafana can automatically create IAM roles and permissions for, to allow Amazon Managed Grafana to use these channels.
	NotificationDestinations *[]NotificationDestinations `field:"optional" json:"notificationDestinations" yaml:"notificationDestinations"`
	// Specifies the organizational units that this workspace is allowed to use data sources from, if this workspace is in an account that is part of an organization.
	OrganizationalUnits *[]*string `field:"optional" json:"organizationalUnits" yaml:"organizationalUnits"`
	// Name of the IAM role to use for the organization.
	//
	// Maximum length of 2048 characters.
	OrganizationRoleName *string `field:"optional" json:"organizationRoleName" yaml:"organizationRoleName"`
	// Whether plugin administration is enabled in the workspace.
	//
	// Setting to true allows workspace
	// admins to install, uninstall, and update plugins from within the Grafana workspace.
	//
	// This option is only valid for workspaces that support Grafana version 9 or newer.
	//
	// Default: false.
	PluginAdminEnabled *bool `field:"optional" json:"pluginAdminEnabled" yaml:"pluginAdminEnabled"`
	// The IAM role that grants permissions to the AWS resources that the workspace will view data from.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// If the workspace uses SAML, use this structure to map SAML assertion attributes to workspace user information and define which groups in the assertion attribute are to have the Admin and Editor roles in the workspace.
	SamlConfiguration *SamlConfiguration `field:"optional" json:"samlConfiguration" yaml:"samlConfiguration"`
	// The name of the AWS CloudFormation stack set that is used to generate IAM roles to be used for this workspace.
	StackSetName *string `field:"optional" json:"stackSetName" yaml:"stackSetName"`
	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to.
	VpcConfiguration *VpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

