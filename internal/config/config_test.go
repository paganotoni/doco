package config

import (
	_ "embed"

	"os"
	"testing"
)

var (
	//go:embed testdata/complete_meta.md
	completeConfig string

	//go:embed testdata/empty_meta.md
	emptyConfig string

	//go:embed testdata/wrong_types_meta.md
	wrongMeta string
)

func TestReadConfig(t *testing.T) {
	t.Run("read complete meta", func(t *testing.T) {
		owd, _ := os.Getwd()
		defer os.Chdir(owd)

		tmpdir := t.TempDir()
		os.Chdir(tmpdir)

		f, err := os.Create("_meta.md")
		if err != nil {
			t.Fatal(err)
		}

		_, err = f.WriteString(completeConfig)
		if err != nil {
			t.Fatal(err)
		}

		err = f.Close()
		if err != nil {
			t.Fatal(err)
		}

		_, err = Read(tmpdir)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("read incomplete meta", func(t *testing.T) {
		owd, _ := os.Getwd()
		defer os.Chdir(owd)

		tmpdir := t.TempDir()
		os.Chdir(tmpdir)

		f, err := os.Create("_meta.md")
		if err != nil {
			t.Fatal(err)
		}

		_, err = f.WriteString(wrongMeta)
		if err != nil {
			t.Fatal(err)
		}

		err = f.Close()
		if err != nil {
			t.Fatal(err)
		}

		_, err = Read(tmpdir)
		if err != nil {
			t.Fatal(err)
		}
	})
}
