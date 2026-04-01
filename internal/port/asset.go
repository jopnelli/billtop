package port

import "context"

// Resource represents a GCP resource from the Cloud Asset Inventory.
type Resource struct {
	Name        string
	AssetType   string // e.g., "compute.googleapis.com/Instance"
	ProjectID   string
	Zone        string
	Labels      map[string]string
	MachineType string
	Status      string
}

// AssetInventoryClient reads resource metadata from GCP Cloud Asset Inventory.
type AssetInventoryClient interface {
	// ListResources returns resources of the given asset type in a project.
	ListResources(ctx context.Context, projectID string, assetType string) ([]Resource, error)
}
