package version_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/paganotoni/doco/cmd/doco/version"
)

func TestVersion(t *testing.T) {
	bb := bytes.NewBuffer([]byte{})
	if err := version.Run(bb); err != nil {
		t.Fatal(err)
	}

	exp := fmt.Sprintf("Doco v1.0.0")
	if bb.String() != exp {
		t.Fatalf("Expected version to be `%s`, got %s", exp, bb.String())
	}
}
