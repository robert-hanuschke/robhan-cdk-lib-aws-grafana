package robhancdklibawsgrafana


// A structure containing arrays that map group names in the SAML assertion to the Grafana Admin and Editor roles in the workspace.
type SamlRoleValues struct {
	// A list of groups from the SAML assertion attribute to grant the Grafana Admin role to.
	//
	// A maximum of 256 elements is allowed (validated by CloudFormation at deploy time).
	// Default: - no groups are granted the Admin role.
	//
	Admin *[]*string `field:"optional" json:"admin" yaml:"admin"`
	// A list of groups from the SAML assertion attribute to grant the Grafana Editor role to.
	//
	// A maximum of 256 elements is allowed (validated by CloudFormation at deploy time).
	// Default: - no groups are granted the Editor role.
	//
	Editor *[]*string `field:"optional" json:"editor" yaml:"editor"`
}

