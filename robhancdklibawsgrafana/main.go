// AWS CDK Construct Library for Grafana
package robhancdklibawsgrafana

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_grafana.AccountAccessType",
		reflect.TypeOf((*AccountAccessType)(nil)).Elem(),
		map[string]interface{}{
			"CURRENT_ACCOUNT": AccountAccessType_CURRENT_ACCOUNT,
			"ORGANIZATION": AccountAccessType_ORGANIZATION,
		},
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_grafana.AuthenticationProviders",
		reflect.TypeOf((*AuthenticationProviders)(nil)).Elem(),
		map[string]interface{}{
			"AWS_SSO": AuthenticationProviders_AWS_SSO,
			"SAML": AuthenticationProviders_SAML,
		},
	)
	_jsii_.RegisterInterface(
		"@robhan-cdk-lib/aws_grafana.IWorkspace",
		reflect.TypeOf((*IWorkspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountAccessType", GoGetter: "AccountAccessType"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationProviders", GoGetter: "AuthenticationProviders"},
			_jsii_.MemberProperty{JsiiProperty: "clientToken", GoGetter: "ClientToken"},
			_jsii_.MemberProperty{JsiiProperty: "dataSources", GoGetter: "DataSources"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkAccessControl", GoGetter: "NetworkAccessControl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationDestinations", GoGetter: "NotificationDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "organizationalUnits", GoGetter: "OrganizationalUnits"},
			_jsii_.MemberProperty{JsiiProperty: "organizationRoleName", GoGetter: "OrganizationRoleName"},
			_jsii_.MemberProperty{JsiiProperty: "permissionType", GoGetter: "PermissionType"},
			_jsii_.MemberProperty{JsiiProperty: "pluginAdminEnabled", GoGetter: "PluginAdminEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "samlConfiguration", GoGetter: "SamlConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stackSetName", GoGetter: "StackSetName"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfiguration", GoGetter: "VpcConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
		},
		func() interface{} {
			j := jsiiProxy_IWorkspace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_grafana.NetworkAccessControl",
		reflect.TypeOf((*NetworkAccessControl)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_grafana.NotificationDestinations",
		reflect.TypeOf((*NotificationDestinations)(nil)).Elem(),
		map[string]interface{}{
			"SNS": NotificationDestinations_SNS,
		},
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_grafana.PermissionTypes",
		reflect.TypeOf((*PermissionTypes)(nil)).Elem(),
		map[string]interface{}{
			"CUSTOMER_MANAGED": PermissionTypes_CUSTOMER_MANAGED,
			"SERVICE_MANAGED": PermissionTypes_SERVICE_MANAGED,
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_grafana.SamlAssertionAttributes",
		reflect.TypeOf((*SamlAssertionAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_grafana.SamlConfiguration",
		reflect.TypeOf((*SamlConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_grafana.SamlConfigurationStatuses",
		reflect.TypeOf((*SamlConfigurationStatuses)(nil)).Elem(),
		map[string]interface{}{
			"CONFIGURED": SamlConfigurationStatuses_CONFIGURED,
			"NOT_CONFIGURED": SamlConfigurationStatuses_NOT_CONFIGURED,
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_grafana.SamlIdpMetadata",
		reflect.TypeOf((*SamlIdpMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_grafana.SamlRoleValues",
		reflect.TypeOf((*SamlRoleValues)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_grafana.Status",
		reflect.TypeOf((*Status)(nil)).Elem(),
		map[string]interface{}{
			"ACTIVE": Status_ACTIVE,
			"CREATING": Status_CREATING,
			"DELETING": Status_DELETING,
			"FAILED": Status_FAILED,
			"UPDATING": Status_UPDATING,
			"UPGRADING": Status_UPGRADING,
			"DELETION_FAILED": Status_DELETION_FAILED,
			"CREATION_FAILED": Status_CREATION_FAILED,
			"UPDATE_FAILED": Status_UPDATE_FAILED,
			"UPGRADE_FAILED": Status_UPGRADE_FAILED,
			"LICENSE_REMOVAL_FAILED": Status_LICENSE_REMOVAL_FAILED,
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_grafana.VpcConfiguration",
		reflect.TypeOf((*VpcConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_grafana.Workspace",
		reflect.TypeOf((*Workspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountAccessType", GoGetter: "AccountAccessType"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationProviders", GoGetter: "AuthenticationProviders"},
			_jsii_.MemberProperty{JsiiProperty: "clientToken", GoGetter: "ClientToken"},
			_jsii_.MemberProperty{JsiiProperty: "creationTimestamp", GoGetter: "CreationTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "dataSources", GoGetter: "DataSources"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getWorkspaceArn", GoMethod: "GetWorkspaceArn"},
			_jsii_.MemberMethod{JsiiMethod: "getWorkspaceId", GoMethod: "GetWorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaVersion", GoGetter: "GrafanaVersion"},
			_jsii_.MemberProperty{JsiiProperty: "modificationTimestamp", GoGetter: "ModificationTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkAccessControl", GoGetter: "NetworkAccessControl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationDestinations", GoGetter: "NotificationDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "organizationalUnits", GoGetter: "OrganizationalUnits"},
			_jsii_.MemberProperty{JsiiProperty: "organizationRoleName", GoGetter: "OrganizationRoleName"},
			_jsii_.MemberProperty{JsiiProperty: "permissionType", GoGetter: "PermissionType"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "pluginAdminEnabled", GoGetter: "PluginAdminEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "samlConfiguration", GoGetter: "SamlConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "samlConfigurationStatus", GoGetter: "SamlConfigurationStatus"},
			_jsii_.MemberProperty{JsiiProperty: "ssoClientId", GoGetter: "SsoClientId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stackSetName", GoGetter: "StackSetName"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfiguration", GoGetter: "VpcConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
		},
		func() interface{} {
			j := jsiiProxy_Workspace{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_WorkspaceBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_grafana.WorkspaceAttributes",
		reflect.TypeOf((*WorkspaceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_grafana.WorkspaceBase",
		reflect.TypeOf((*WorkspaceBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountAccessType", GoGetter: "AccountAccessType"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationProviders", GoGetter: "AuthenticationProviders"},
			_jsii_.MemberProperty{JsiiProperty: "clientToken", GoGetter: "ClientToken"},
			_jsii_.MemberProperty{JsiiProperty: "dataSources", GoGetter: "DataSources"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getWorkspaceArn", GoMethod: "GetWorkspaceArn"},
			_jsii_.MemberMethod{JsiiMethod: "getWorkspaceId", GoMethod: "GetWorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkAccessControl", GoGetter: "NetworkAccessControl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationDestinations", GoGetter: "NotificationDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "organizationalUnits", GoGetter: "OrganizationalUnits"},
			_jsii_.MemberProperty{JsiiProperty: "organizationRoleName", GoGetter: "OrganizationRoleName"},
			_jsii_.MemberProperty{JsiiProperty: "permissionType", GoGetter: "PermissionType"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "pluginAdminEnabled", GoGetter: "PluginAdminEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "samlConfiguration", GoGetter: "SamlConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stackSetName", GoGetter: "StackSetName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfiguration", GoGetter: "VpcConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspaceBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWorkspace)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_grafana.WorkspaceProps",
		reflect.TypeOf((*WorkspaceProps)(nil)).Elem(),
	)
}
