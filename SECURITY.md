# Security Policy

## Reporting a Vulnerability

If you discover a security vulnerability in billtop, please report it via [GitHub Security Advisories](https://github.com/jopnelli/billtop/security/advisories/new).

**Do not** open a public issue for security vulnerabilities.

You should receive an initial response within 48 hours. We will work with you to understand the issue and coordinate a fix before any public disclosure.

## Scope

billtop handles GCP billing data, which is sensitive financial information. The following are in scope:

- Authentication bypass on the web dashboard
- SQL injection in BigQuery query construction
- Credential leakage (logging, error messages, config files)
- Arbitrary code execution
- Path traversal in config file handling
- Secrets exposed in Docker image layers

## Security Design

- billtop never stores GCP credentials — it uses Application Default Credentials (ADC)
- Config files do not contain secrets; webhook URLs use `${ENV_VAR}` interpolation
- The web dashboard binds to localhost by default and requires auth when exposed
- BigQuery queries use parameterized queries, never string interpolation
- Docker images use distroless base, run as non-root, with read-only filesystem
