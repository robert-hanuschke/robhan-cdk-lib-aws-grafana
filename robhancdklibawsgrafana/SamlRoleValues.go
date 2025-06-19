package robhancdklibawsgrafana


// A structure containing arrays that map group names in the SAML assertion to the Grafana Admin and Editor roles in the workspace.
type SamlRoleValues struct {
	// A list of groups from the SAML assertion attribute to grant the Grafana Admin role to.
	//
	// Maximum of 256 elements.
	Admin *[]*string `field:"optional" json:"admin" yaml:"admin"`
	// A list of groups from the SAML assertion attribute to grant the Grafana Editor role to.
	//
	// Maximum of 256 elements.
	Editor *[]*string `field:"optional" json:"editor" yaml:"editor"`
}

