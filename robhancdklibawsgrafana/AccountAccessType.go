package robhancdklibawsgrafana


// Specifies whether the workspace can access AWS resources in this AWS account only, or whether it can also access AWS resources in other accounts in the same organization.
//
// If this is
// ORGANIZATION, the OrganizationalUnits parameter specifies which organizational units the
// workspace can access.
type AccountAccessType string

const (
	// Access is limited to the current AWS account only.
	AccountAccessType_CURRENT_ACCOUNT AccountAccessType = "CURRENT_ACCOUNT"
	// Access is extended to the entire AWS organization.
	AccountAccessType_ORGANIZATION AccountAccessType = "ORGANIZATION"
)

