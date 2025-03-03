package internal

import (
	_ "embed"
	"os"
	"path/filepath"
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

func TestLoadConfigFile(t *testing.T) {
	t.Run("loads existing config file", func(t *testing.T) {
		tmpdir := t.TempDir()
		configPath := filepath.Join(tmpdir, metafile)

		err := os.WriteFile(configPath, []byte(completeConfig), 0644)
		if err != nil {
			t.Fatal(err)
		}

		content, err := loadConfigFile(tmpdir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if string(content) != completeConfig {
			t.Error("loaded content doesn't match expected content")
		}
	})

	t.Run("returns error for non-existent file", func(t *testing.T) {
		tmpdir := t.TempDir()
		_, err := loadConfigFile(tmpdir)
		if err == nil {
			t.Error("expected error for non-existent file")
		}
	})
}

func TestParseConfig(t *testing.T) {
	t.Run("parses complete config", func(t *testing.T) {
		config, err := parseConfig([]byte(completeConfig))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Test basic fields
		if config.Name != "Doco" {
			t.Errorf("expected Name to be 'Doco', got %s", config.Name)
		}
		if config.Description != "Doco is a CLI tool to generate static documentation websites from markdown files." {
			t.Errorf("unexpected Description: %s", config.Description)
		}
		if config.Favicon != "/assets/favicon.png" {
			t.Errorf("expected Favicon to be '/assets/favicon.png', got %s", config.Favicon)
		}
		if config.OGImage != "/assets/og.png" {
			t.Errorf("expected OGImage to be '/assets/og.png', got %s", config.OGImage)
		}

		// Test logo configuration
		if config.Logo.ImageSrc != "/assets/logo.png" {
			t.Errorf("expected Logo.ImageSrc to be '/assets/logo.png', got %s", config.Logo.ImageSrc)
		}
		if config.Logo.Link != "/" {
			t.Errorf("expected Logo.Link to be '/', got %s", config.Logo.Link)
		}

		// Test announcement
		if config.Announcement.Text != "Check our Github repository." {
			t.Errorf("unexpected Announcement.Text: %s", config.Announcement.Text)
		}
		if config.Announcement.Link != "https://github.com/paganotoni/doco" {
			t.Errorf("unexpected Announcement.Link: %s", config.Announcement.Link)
		}

		// Test quick links
		if len(config.QuickLinks) != 2 {
			t.Errorf("expected 2 quick links, got %d", len(config.QuickLinks))
		}
		if config.QuickLinks[0].Text != "Documentation" {
			t.Errorf("unexpected QuickLinks[0].Text: %s", config.QuickLinks[0].Text)
		}
		if config.QuickLinks[0].Icon != "menu_book" {
			t.Errorf("unexpected QuickLinks[0].Icon: %s", config.QuickLinks[0].Icon)
		}

		// Test external links
		if len(config.ExternalLinks) != 1 {
			t.Errorf("expected 1 external link, got %d", len(config.ExternalLinks))
		}
		if config.ExternalLinks[0].Text != "Documentation" {
			t.Errorf("unexpected ExternalLinks[0].Text: %s", config.ExternalLinks[0].Text)
		}
	})

	t.Run("handles empty config", func(t *testing.T) {
		config, err := parseConfig([]byte(emptyConfig))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Should use default values
		if config.Name != defaultName {
			t.Errorf("expected default Name '%s', got %s", defaultName, config.Name)
		}
		if config.Description != defaultDescription {
			t.Errorf("expected default Description '%s', got %s", defaultDescription, config.Description)
		}
	})

	t.Run("handles wrong types gracefully", func(t *testing.T) {
		config, err := parseConfig([]byte(wrongMeta))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Should use default values for wrong types
		if config.Name != defaultName {
			t.Errorf("expected default Name '%s', got %s", defaultName, config.Name)
		}
		if config.Description != defaultDescription {
			t.Errorf("expected default Description '%s', got %s", defaultDescription, config.Description)
		}
	})
}

func TestReadConfig(t *testing.T) {
	t.Run("reads complete config", func(t *testing.T) {
		tmpdir := t.TempDir()
		configPath := filepath.Join(tmpdir, metafile)

		err := os.WriteFile(configPath, []byte(completeConfig), 0644)
		if err != nil {
			t.Fatal(err)
		}

		config, err := ReadConfig(tmpdir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Test basic fields
		if config.Name != "Doco" {
			t.Errorf("expected Name to be 'Doco', got %s", config.Name)
		}
		if config.Description != "Doco is a CLI tool to generate static documentation websites from markdown files." {
			t.Errorf("unexpected Description: %s", config.Description)
		}
	})

	t.Run("returns error for non-existent file", func(t *testing.T) {
		tmpdir := t.TempDir()
		_, err := ReadConfig(tmpdir)
		if err == nil {
			t.Error("expected error for non-existent file")
		}
	})
}
