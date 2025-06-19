package robhancdklibawsgrafana


// If this is SERVICE_MANAGED, and the workplace was created through the Amazon Managed Grafana console, then Amazon Managed Grafana automatically creates the IAM roles and provisions the permissions that the workspace needs to use AWS data sources and notification channels.
//
// If this is CUSTOMER_MANAGED, you must manage those roles and permissions yourself.
//
// If you are working with a workspace in a member account of an organization and that account is
// not a delegated administrator account, and you want the workspace to access data sources in
// other AWS accounts in the organization, this parameter must be set to CUSTOMER_MANAGED.
type PermissionTypes string

const (
	// Customer-managed permissions where you manage user access to Grafana.
	PermissionTypes_CUSTOMER_MANAGED PermissionTypes = "CUSTOMER_MANAGED"
	// Service-managed permissions where AWS manages user access to Grafana.
	PermissionTypes_SERVICE_MANAGED PermissionTypes = "SERVICE_MANAGED"
)

