package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jopnelli/billtop/internal/version"
)

func TestVersionCmd_Output(t *testing.T) {
	version.Version = "1.2.3"
	version.Commit = "abc1234"
	version.Date = "2026-01-01T00:00:00Z"
	t.Cleanup(func() {
		version.Version = "dev"
		version.Commit = "none"
		version.Date = "unknown"
	})

	cmd := NewRootCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"version"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "1.2.3") {
		t.Errorf("expected version 1.2.3 in output, got: %s", out)
	}
	if !strings.Contains(out, "abc1234") {
		t.Errorf("expected commit abc1234 in output, got: %s", out)
	}
}
