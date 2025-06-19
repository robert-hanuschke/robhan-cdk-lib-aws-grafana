package robhancdklibawsgrafana


// Status of a Grafana workspace.
type Status string

const (
	// Workspace is active and ready to use.
	Status_ACTIVE Status = "ACTIVE"
	// Workspace is being created.
	Status_CREATING Status = "CREATING"
	// Workspace is being deleted.
	Status_DELETING Status = "DELETING"
	// Workspace operation has failed.
	Status_FAILED Status = "FAILED"
	// Workspace is being updated.
	Status_UPDATING Status = "UPDATING"
	// Workspace is being upgraded.
	Status_UPGRADING Status = "UPGRADING"
	// Workspace deletion has failed.
	Status_DELETION_FAILED Status = "DELETION_FAILED"
	// Workspace creation has failed.
	Status_CREATION_FAILED Status = "CREATION_FAILED"
	// Workspace update has failed.
	Status_UPDATE_FAILED Status = "UPDATE_FAILED"
	// Workspace upgrade has failed.
	Status_UPGRADE_FAILED Status = "UPGRADE_FAILED"
	// License removal has failed.
	Status_LICENSE_REMOVAL_FAILED Status = "LICENSE_REMOVAL_FAILED"
)

