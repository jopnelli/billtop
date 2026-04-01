# billtop

Open source GCP cost optimization. See where your money goes. Stop wasting it.

> **Alpha** — billtop is under active development. APIs and config formats may change.

[![CI](https://github.com/jopnelli/billtop/actions/workflows/ci.yml/badge.svg)](https://github.com/jopnelli/billtop/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/jopnelli/billtop)](LICENSE)

## What is billtop?

billtop connects to your GCP billing export in BigQuery and gives you:

- **Cost visibility** — where your money goes, by project, service, and SKU, in human-readable names
- **Waste detection** — idle disks, oversized VMs, unused IPs, and more
- **Trend analysis** — what's going up, what's going down, and why
- **Anomaly alerts** — get notified before costs spiral (Slack, webhooks, PagerDuty)
- **CUD tracking** — monitor commitment utilization (Google doesn't alert on this)
- **Custom labels** — group costs by team, environment, or app without re-tagging GCP resources

No enterprise sales call. No billing account delegation. No 6-week onboarding.

## Quick Start

```bash
# Install
go install github.com/jopnelli/billtop/cmd/billtop@latest

# Set up (creates BigQuery dataset, enables APIs, guides you through billing export)
billtop setup

# See where your money goes
billtop report

# Find savings
billtop scan
```

Or with Docker:

```bash
docker run --rm \
  -v ~/.config/gcloud/application_default_credentials.json:/credentials/adc.json:ro \
  -e GOOGLE_APPLICATION_CREDENTIALS=/credentials/adc.json \
  ghcr.io/jopnelli/billtop:latest report
```

## Requirements

- A GCP project with [billing export to BigQuery](https://cloud.google.com/billing/docs/how-to/export-data-bigquery-setup) enabled
- `gcloud` CLI authenticated (`gcloud auth application-default login`)
- BigQuery Data Viewer permission on the billing dataset

billtop reads your billing data from BigQuery. It never writes to it, never stores credentials, and never phones home.

## Documentation

Coming soon. For now, see `billtop --help`.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[Apache 2.0](LICENSE)
