// in example_test.go
package example

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1)
	if got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}
