package robhancdklibawsgrafana


// Status of SAML configuration for a Grafana workspace.
type SamlConfigurationStatuses string

const (
	// SAML is configured for the workspace.
	SamlConfigurationStatuses_CONFIGURED SamlConfigurationStatuses = "CONFIGURED"
	// SAML is not configured for the workspace.
	SamlConfigurationStatuses_NOT_CONFIGURED SamlConfigurationStatuses = "NOT_CONFIGURED"
)

