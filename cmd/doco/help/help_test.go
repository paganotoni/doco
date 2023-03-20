package help_test

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/paganotoni/doco/cmd/doco/help"
)

//go:embed help.txt
var content string

func TestRun(t *testing.T) {
	bb := bytes.NewBuffer([]byte{})
	if err := help.Run(bb); err != nil {
		t.Fatal(err)
	}

	if bb.String() != content {
		t.Fatal("expected output")
	}
}
