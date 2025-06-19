package robhancdklibawsgrafana


// If the workspace uses SAML, use this structure to map SAML assertion attributes to workspace user information and define which groups in the assertion attribute are to have the Admin and Editor roles in the workspace.
type SamlConfiguration struct {
	// A structure containing the identity provider (IdP) metadata used to integrate the identity provider with this workspace.
	//
	// Required field for SAML configuration.
	IdpMetadata *SamlIdpMetadata `field:"required" json:"idpMetadata" yaml:"idpMetadata"`
	// Lists which organizations defined in the SAML assertion are allowed to use the Amazon Managed Grafana workspace.
	//
	// If this is empty, all organizations in the assertion attribute have access.
	//
	// Must have between 1 and 256 elements.
	AllowedOrganizations *[]*string `field:"optional" json:"allowedOrganizations" yaml:"allowedOrganizations"`
	// A structure that defines which attributes in the SAML assertion are to be used to define information about the users authenticated by that IdP to use the workspace.
	AssertionAtrributes *SamlAssertionAttributes `field:"optional" json:"assertionAtrributes" yaml:"assertionAtrributes"`
	// How long a sign-on session by a SAML user is valid, before the user has to sign on again.
	//
	// Must be a positive number.
	LoginValidityDuration *float64 `field:"optional" json:"loginValidityDuration" yaml:"loginValidityDuration"`
	// A structure containing arrays that map group names in the SAML assertion to the Grafana Admin and Editor roles in the workspace.
	RoleValues *SamlRoleValues `field:"optional" json:"roleValues" yaml:"roleValues"`
}

