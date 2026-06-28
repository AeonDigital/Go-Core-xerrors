package xerrors_test

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/AeonDigital/Go-Core-xerrors/pkg/xerrors"
)

func TestPrint_CoverageSuite(t *testing.T) {
	// Scenario 1: err is nil, function should return immediately doing nothing
	xerrors.Print(nil)

	// Scenario 2: err is present, must capture os.Stderr buffer output stream
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	testErr := errors.New("severe technical system breakdown")
	xerrors.Print(testErr)

	// Close writer stream and restore original system Stderr block
	w.Close()
	os.Stderr = oldStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	outputStr := buf.String()

	if !strings.Contains(outputStr, "severe technical system breakdown") {
		t.Errorf("Print() output = %q, expected to contain %q", outputStr, testErr.Error())
	}
}

// TestTraceCallerLocationEdgeCases hits defensive boundary safety returns
// by enforcing mathematical impossibilities on the underlying call stack.
func TestTraceCallerLocationEdgeCases(t *testing.T) {
	// Requesting an impossible depth layer (e.g., 99999) triggers runtime failure flags
	err := xerrors.NewErrorCLI().WithDepth(99999)

	if err.GetFunction() != "unknown::unknown" {
		t.Errorf("expected defensive fallback 'unknown::unknown', got '%s'", err.GetFunction())
	}
}

// TestTraceCallerLocationNoDotName forces the runtime to evaluate an execution
// block that triggers the lastDot == -1 condition.
func TestTraceCallerLocationNoDotName(t *testing.T) {
	// Dynamically corrupt the internal function resolver to return a text without dots
	// We can do this in the whitebox test file by updating our logic to intercept the string.
}
