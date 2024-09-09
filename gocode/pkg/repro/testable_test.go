package repro_test

import (
	"testing"

	"foo/pkg/repro"
)

func TestTestable(t *testing.T) {
	result := repro.Testable()
	if !result {
		t.Fatalf("it should be true")
	}
}

