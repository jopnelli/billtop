package port

import (
	"context"
	"time"
)

// MetricPoint is a single data point in a time series.
type MetricPoint struct {
	Timestamp time.Time
	Value     float64
}

// MetricsClient reads utilization metrics from GCP Cloud Monitoring.
type MetricsClient interface {
	// CPUUtilization returns the CPU utilization time series for a resource over the given period.
	CPUUtilization(ctx context.Context, projectID string, resourceName string, start, end time.Time) ([]MetricPoint, error)
}
