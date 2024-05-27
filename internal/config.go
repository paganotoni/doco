package internal

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

const (
	// metafile is the name of the file that contains the
	// configuration of the site.
	metafile = "_meta.md"
)

// config of the general elements of the site.
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
}

type Link struct {
	Text     string
	Link     string
	Icon     string
	ImageSrc string
}

// Read parses the _meta.md file and returns the config
// for the site.
// TODO: change this to receive the file access (fs package?) instead of the folder.
func readConfig(folder string) (c siteConfig, err error) {
	file, err := os.Open(filepath.Join(folder, metafile))
	if err != nil {
		return c, err
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return c, err
	}

	// Parse the metadata and apply it to the document
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := mparser.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return c, err
	}

	meta := meta.Get(context)
	def := func(val any, defs string) string {
		v, ok := val.(string)
		if !ok || v == "" {
			return defs
		}

		return v
	}

	c.Name = def(meta["name"], "Doco")
	c.Description = def(meta["description"], "Documentation site")
	c.Keywords = def(meta["keywords"], "documentation, site, doco")
	c.Copy = def(meta["copy"], "© $YEAR Doco")
	c.Github = def(meta["github"], "https://github.com/paganotoni/doco")
	c.Favicon = def(meta["favicon"], "")
	c.OGImage = def(meta["ogimage"], "")

	logo, ok := meta["logo"].(map[any]any)
	if ok {
		c.Logo.ImageSrc = def(logo["src"], "")
		c.Logo.Link = def(logo["link"], "")
	}

	announcement, ok := meta["announcement"].(map[any]any)
	if ok {
		c.Announcement.Text = def(announcement["text"], "")
		c.Announcement.Link = def(announcement["link"], "")
	}

	qlinks, ok := meta["quick_links"].([]any)
	if ok {
		for _, v := range qlinks {
			l := v.(map[any]any)
			c.QuickLinks = append(c.QuickLinks, Link{
				Text: def(l["text"], ""),
				Link: def(l["link"], ""),
				Icon: def(l["icon"], ""),
			})
		}
	}

	elinks, ok := meta["external_links"].([]any)
	if ok {
		for _, v := range elinks {
			l := v.(map[any]any)
			c.ExternalLinks = append(c.ExternalLinks, Link{
				Text: def(l["text"], ""),
				Link: def(l["link"], ""),
			})
		}
	}

	return c, nil
}
