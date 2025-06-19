package robhancdklibawsgrafana


// The AWS notification channels that Amazon Managed Grafana can automatically create IAM roles and permissions for, to allow Amazon Managed Grafana to use these channels.
type NotificationDestinations string

const (
	// Amazon Simple Notification Service (SNS) as notification destination.
	NotificationDestinations_SNS NotificationDestinations = "SNS"
)

