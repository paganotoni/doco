package internal

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark/parser"
)

const (
	// metafile is the name of the file that contains the
	// configuration of the site.
	metafile = "_meta.md"

	// Default values for configuration
	defaultName        = "Doco"
	defaultDescription = "Documentation site"
	defaultKeywords    = "documentation, site, doco"
	defaultCopy        = "© $YEAR Doco"
	defaultGithub      = "https://github.com/paganotoni/doco"
)

// siteConfig represents the configuration of the site
type siteConfig struct {
	Name        string
	Favicon     string
	Description string
	Keywords    string
	Github      string // Github link to display, empty means no link

	Logo          Link
	Announcement  Link
	ExternalLinks []Link
	QuickLinks    []Link

	Copy    string
	OGImage string
	// Files and folders to be ignored
	Ignore []string
}

type Link struct {
	Text     string
	Link     string
	Icon     string
	ImageSrc string
}

// loadConfigFile reads the configuration file from the given folder
func loadConfigFile(folder string) ([]byte, error) {
	file, err := os.Open(filepath.Join(folder, metafile))
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	return content, nil
}

// parseConfig parses the configuration from the given content
func parseConfig(content []byte) (siteConfig, error) {
	var c siteConfig

	// Parse the metadata and apply it to the document
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := mparser.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return c, fmt.Errorf("failed to convert markdown: %w", err)
	}

	meta, err := parseMeta(content)
	if err != nil {
		return c, fmt.Errorf("failed to parse meta: %w", err)
	}

	def := func(val any, defs string) string {
		v, ok := val.(string)
		if !ok || v == "" {
			return defs
		}
		return v
	}

	// Parse basic configuration
	c.Name = def(meta["name"], defaultName)
	c.Description = def(meta["description"], defaultDescription)
	c.Keywords = def(meta["keywords"], defaultKeywords)
	c.Copy = def(meta["copy"], defaultCopy)
	c.Github = def(meta["github"], defaultGithub)
	c.Favicon = def(meta["favicon"], "")
	c.OGImage = def(meta["ogimage"], "")

	// Parse logo configuration
	if logo, ok := meta["logo"].(map[string]any); ok {
		c.Logo.ImageSrc = def(logo["src"], "")
		c.Logo.Link = def(logo["link"], "")
	}

	// Parse announcement configuration
	if announcement, ok := meta["announcement"].(map[string]any); ok {
		c.Announcement.Text = def(announcement["text"], "")
		c.Announcement.Link = def(announcement["link"], "")
	}

	// Parse quick links
	if qlinks, ok := meta["quick_links"].([]any); ok {
		for _, v := range qlinks {
			if l, ok := v.(map[string]any); ok {
				c.QuickLinks = append(c.QuickLinks, Link{
					Text: def(l["text"], ""),
					Link: def(l["link"], ""),
					Icon: def(l["icon"], ""),
				})
			}
		}
	}

	// Parse external links
	if elinks, ok := meta["external_links"].([]any); ok {
		for _, v := range elinks {
			if l, ok := v.(map[string]any); ok {
				c.ExternalLinks = append(c.ExternalLinks, Link{
					Text: def(l["text"], ""),
					Link: def(l["link"], ""),
				})
			}
		}
	}

	// Set default ignore patterns
	if len(c.Ignore) == 0 {
		c.Ignore = []string{"README.md", "README"}
	}

	return c, nil
}

// ReadConfig parses the _meta.md file and returns the config for the site.
func ReadConfig(folder string) (siteConfig, error) {
	content, err := loadConfigFile(folder)
	if err != nil {
		return siteConfig{}, err
	}

	return parseConfig(content)
}
