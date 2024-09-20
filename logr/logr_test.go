//go:build unit

package logr

import (
	"testing"

	"github.com/go-logr/logr"
)

func TestLogrInterface(t *testing.T) {
	var logrLogger Logger

	logrLogger = logr.New(nil)
	t.Logf("logr: %v", logrLogger)
}
