# billtop

Open source GCP cost optimization. See where your money goes. Stop wasting it.

> **Alpha** — under active development. Not production ready.

[![CI](https://github.com/jopnelli/billtop/actions/workflows/ci.yml/badge.svg)](https://github.com/jopnelli/billtop/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/jopnelli/billtop)](LICENSE)

## What is billtop?

billtop connects to your GCP billing export in BigQuery and makes it legible:

- **Cost breakdown** by project, service, and SKU — in human-readable names
- **Custom labels** — group costs by team, environment, or app without re-tagging GCP resources
- **Credits decoded** — see exactly how SUDs, CUDs, and promotions affect your bill
- **Trends** — what's going up, what's going down, and why
- **Web dashboard** — visual drill-down into your spending

No enterprise sales call. No billing account delegation.

## Quick Start

```bash
# Install
bun install -g @billtop/cli

# Set up (creates BigQuery dataset, enables APIs, guides you through billing export)
billtop setup

# See where your money goes
billtop report
```

## Requirements

- [bun](https://bun.sh) runtime
- A GCP project with [billing export to BigQuery](https://cloud.google.com/billing/docs/how-to/export-data-bigquery-setup) enabled
- `gcloud` CLI authenticated (`gcloud auth application-default login`)

billtop reads your billing data from BigQuery. It never writes to it, never stores credentials, and never phones home.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[Apache 2.0](LICENSE)
