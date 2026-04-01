// Package version holds build-time version information.
// Values are set via ldflags during the build process.
package version

var (
	// Version is the semantic version (e.g., "0.1.0").
	Version = "dev"

	// Commit is the git commit SHA at build time.
	Commit = "none"

	// Date is the build timestamp in RFC3339 format.
	Date = "unknown"
)
