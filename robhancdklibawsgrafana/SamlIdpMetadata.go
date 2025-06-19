package robhancdklibawsgrafana


// A structure containing the identity provider (IdP) metadata used to integrate the identity provider with this workspace.
type SamlIdpMetadata struct {
	// The URL of the location containing the IdP metadata.
	//
	// Must be a string with length between 1 and 2048 characters.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The full IdP metadata, in XML format.
	Xml *string `field:"optional" json:"xml" yaml:"xml"`
}

