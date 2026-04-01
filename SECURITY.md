# Security Policy

## Reporting a Vulnerability

Report via [GitHub Security Advisories](https://github.com/jopnelli/billtop/security/advisories/new). Do not open a public issue.

## Scope

billtop handles GCP billing data (sensitive financial information). In scope:

- SQL injection in BigQuery query construction
- Credential leakage (logging, error messages, config files)
- Authentication bypass on the web dashboard
- Secrets exposed in build artifacts

## Design Principles

- Never stores GCP credentials — uses Application Default Credentials
- Config files contain no secrets — webhook URLs use `${ENV_VAR}` interpolation
- Web dashboard binds to localhost by default
- BigQuery queries use parameterized queries, never string interpolation
