// Package port defines the interfaces (ports) that the domain layer requires.
// Implementations live in internal/adapter/.
package port

import (
	"context"

	"github.com/jopnelli/billtop/internal/domain/model"
)

// BillingReader provides read access to GCP billing data.
type BillingReader interface {
	// CostByService returns aggregated costs grouped by GCP service for the given period.
	CostByService(ctx context.Context, period model.Period) ([]model.ServiceCost, error)

	// CostByProject returns aggregated costs grouped by GCP project for the given period.
	CostByProject(ctx context.Context, period model.Period) ([]model.ProjectCost, error)

	// TableExists checks whether the billing export table exists and has data.
	TableExists(ctx context.Context) (bool, error)
}
