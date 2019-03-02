// in example_test.go
package example_test

import (
    "testing"

    . "bitbucket.org/splice/blog/example"
)

func TestAdd(t *testing.T) {
    got := Add(1)
    if got != 1 {
        t.Errorf("got %d, want 1", got)
    }
}
