package robhancdklibawsgrafana

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The configuration settings for network access to your workspace.
type NetworkAccessControl struct {
	// An array of prefix list IDs.
	//
	// A prefix list is a list of CIDR ranges of IP addresses. The IP
	// addresses specified are allowed to access your workspace. If the list is not included in the
	// configuration (passed an empty array) then no IP addresses are allowed to access the
	// workspace.
	//
	// Maximum of 5 prefix lists allowed.
	PrefixLists *[]awsec2.IPrefixList `field:"optional" json:"prefixLists" yaml:"prefixLists"`
	// An array of Amazon VPC endpoint IDs for the workspace.
	//
	// You can create VPC endpoints to your
	// Amazon Managed Grafana workspace for access from within a VPC. If a NetworkAccessConfiguration
	// is specified then only VPC endpoints specified here are allowed to access the workspace. If
	// you pass in an empty array of strings, then no VPCs are allowed to access the workspace.
	//
	// Maximum of 5 VPC endpoints allowed.
	VpcEndpoints *[]awsec2.IVpcEndpoint `field:"optional" json:"vpcEndpoints" yaml:"vpcEndpoints"`
}

