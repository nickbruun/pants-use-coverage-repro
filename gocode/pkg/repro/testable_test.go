package repro

import (
	"testing"
)

func TestTestable(t *testing.T) {
	result := Testable()
	if !result {
		t.Fatalf("it should be true")
	}
}
