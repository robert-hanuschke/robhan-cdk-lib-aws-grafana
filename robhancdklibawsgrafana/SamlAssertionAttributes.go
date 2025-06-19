package robhancdklibawsgrafana


// A structure that defines which attributes in the IdP assertion are to be used to define information about the users authenticated by the IdP to use the workspace.
//
// Each attribute must be a string with length between 1 and 256 characters.
type SamlAssertionAttributes struct {
	// The name of the attribute within the SAML assertion to use as the email names for SAML users.
	//
	// Must be between 1 and 256 characters long.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The name of the attribute within the SAML assertion to use as the user full "friendly" names for user groups.
	//
	// Must be between 1 and 256 characters long.
	Groups *string `field:"optional" json:"groups" yaml:"groups"`
	// The name of the attribute within the SAML assertion to use as the login names for SAML users.
	//
	// Must be between 1 and 256 characters long.
	Login *string `field:"optional" json:"login" yaml:"login"`
	// The name of the attribute within the SAML assertion to use as the user full "friendly" names for SAML users.
	//
	// Must be between 1 and 256 characters long.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the attribute within the SAML assertion to use as the user full "friendly" names for the users' organizations.
	//
	// Must be between 1 and 256 characters long.
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The name of the attribute within the SAML assertion to use as the user roles.
	//
	// Must be between 1 and 256 characters long.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

