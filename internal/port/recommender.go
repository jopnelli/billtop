package port

import "context"

// Recommendation represents a single cost optimization recommendation from GCP.
type Recommendation struct {
	ID               string
	Type             string // e.g., "VM_RIGHTSIZING", "IDLE_VM", "IDLE_SQL"
	ResourceName     string
	ProjectID        string
	Zone             string
	Description      string
	Currency         string
	CurrentState     string // e.g., "n2-standard-8"
	RecommendedState string // e.g., "n2-standard-4"
	MonthlySavings   float64
}

// RecommenderClient reads optimization recommendations from the GCP Recommender API.
type RecommenderClient interface {
	// ListRecommendations returns cost optimization recommendations for a project.
	ListRecommendations(ctx context.Context, projectID string) ([]Recommendation, error)
}
