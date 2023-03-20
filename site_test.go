package doco_test

import (
	"testing"

	"github.com/paganotoni/doco"
)

func TestSiteName(t *testing.T) {
	t.Run("Name Specified", func(t *testing.T) {
		s := doco.NewSite([]doco.Document{
			testdocument{"docs/_index.md", "---\nName: Site\n---\n"},
		})

		if s.Name() != "Site" {
			t.Fatal("expected name to be Site")
		}
	})

	t.Run("Name lowercase", func(t *testing.T) {
		s := doco.NewSite([]doco.Document{
			testdocument{"docs/_index.md", "---\nname: Site\n---\n"},
		})

		if s.Name() != "" {
			t.Fatal("expected name to be \"\"")
		}
	})

	t.Run("No Name", func(t *testing.T) {
		s := doco.NewSite([]doco.Document{
			testdocument{"docs/_index.md", "---\nOther: Site\n---\n"},
		})

		if s.Name() != "" {
			t.Fatal("expected name to be \"\"")
		}
	})

	t.Run("No Index", func(t *testing.T) {
		s := doco.NewSite([]doco.Document{})

		if s.Name() != "" {
			t.Fatal("expected name to be \"\"")
		}
	})

}
