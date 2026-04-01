// Package model defines the core domain types used throughout billtop.
package model

import "time"

// CostRecord represents a single line item from the GCP billing export.
type CostRecord struct {
	BillingAccountID string
	ProjectID        string
	ProjectName      string
	ServiceID        string
	ServiceName      string
	SKUID            string
	SKUDescription   string
	Region           string
	Zone             string
	UsageStartTime   time.Time
	UsageEndTime     time.Time
	Cost             float64
	Currency         string
	Credits          []Credit
	Labels           map[string]string
	ResourceName     string
}

// Credit represents a discount or credit applied to a cost record.
type Credit struct {
	Name   string
	Amount float64
	Type   string // e.g., "SUSTAINED_USAGE_DISCOUNT", "COMMITTED_USAGE_DISCOUNT", "PROMOTION"
}

// ServiceCost is an aggregated cost for a single GCP service.
type ServiceCost struct {
	ServiceName  string
	GrossCost    float64
	TotalCredits float64
	NetCost      float64
	Currency     string
}

// ProjectCost is an aggregated cost for a single GCP project.
type ProjectCost struct {
	ProjectID   string
	ProjectName string
	GrossCost   float64
	NetCost     float64
	Currency    string
}

// Period defines a time range for cost queries.
type Period struct {
	Start time.Time
	End   time.Time
}
