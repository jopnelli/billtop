package port

import "context"

// Alert represents a cost anomaly or threshold breach to be delivered.
type Alert struct {
	Name        string
	Severity    string // "info", "warning", "critical"
	Message     string
	Service     string
	ProjectID   string
	Amount      float64
	Percentage  float64
	Currency    string
}

// AlertNotifier delivers cost alerts to external systems.
type AlertNotifier interface {
	// Send delivers an alert. Implementations handle formatting for their target (Slack, webhook, etc.).
	Send(ctx context.Context, alert Alert) error
}
