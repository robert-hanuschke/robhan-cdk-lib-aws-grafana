package robhancdklibawsgrafana


// Specifies whether this workspace uses SAML 2.0, AWS IAM Identity Center, or both to authenticate users for using the Grafana console within a workspace.
// See: https://docs.aws.amazon.com/grafana/latest/APIReference/API_CreateWorkspace.html
//
type AuthenticationProviders string

const (
	// AWS Single Sign-On authentication provider.
	AuthenticationProviders_AWS_SSO AuthenticationProviders = "AWS_SSO"
	// Security Assertion Markup Language (SAML) authentication provider.
	AuthenticationProviders_SAML AuthenticationProviders = "SAML"
)

