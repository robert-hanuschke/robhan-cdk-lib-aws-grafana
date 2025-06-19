package robhancdklibawsgrafana

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to.
type VpcConfiguration struct {
	// The list of Amazon EC2 security groups attached to the Amazon VPC for your Grafana workspace to connect.
	//
	// Duplicates not allowed.
	//
	// Array Members: Minimum number of 1 items. Maximum number of 5 items.
	//
	// Required for VPC configuration.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The list of Amazon EC2 subnets created in the Amazon VPC for your Grafana workspace to connect. Duplicates not allowed.
	//
	// Array Members: Minimum number of 2 items. Maximum number of 6 items.
	//
	// Required for VPC configuration.
	Subnets *[]awsec2.ISubnet `field:"required" json:"subnets" yaml:"subnets"`
}

