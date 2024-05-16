package internal

import (
	"cmp"
	"io"
	"os"
	"path/filepath"
)

const (
	// metafile is the name of the file that contains the
	// configuration of the site.
	metafile = "_meta.md"
)

// config of the general elements of the site.
type config struct {
	Name        string
	Favicon     string
	Description string
	Keywords    string
	Github      string // Github link to display, empty means no link

	Announcement struct {
		Text string // Text to display, empty means no announcement
		Link string
	}

	Logo struct {
		Link string
		Src  string
	}

	QuickLinks []struct {
		Text string
		Link string
		Icon string
	}

	ExternalLinks []struct {
		Text string
		Link string
	}

	Copy string
}

// readConfig parses the _meta.md file and returns the config
// for the site.
// TODO: change this to receive the file access (fs package?) instead of the folder.
func readConfig(folder string) (c config, err error) {
	file, err := os.Open(filepath.Join(folder, metafile))
	if err != nil {
		return c, err
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return c, err
	}

	meta, err := metadataFrom(content)
	if err != nil {
		return c, err
	}

	c.Name = cmp.Or(meta["name"].(string), "Doco")
	c.Description = cmp.Or(meta["description"].(string), "Documentation site")
	c.Keywords = cmp.Or(meta["keywords"].(string), "documentation, site, doco")
	c.Copy = cmp.Or(meta["copy"].(string), "© $YEAR Doco")
	c.Github = cmp.Or(meta["github"].(string), "https://github.com/paganotoni/doco")
	c.Favicon = cmp.Or(meta["favicon"].(string), "")

	logo, ok := meta["logo"].(map[any]any)
	if ok {
		c.Logo.Src = cmp.Or(logo["src"].(string), "")
		c.Logo.Link = cmp.Or(logo["link"].(string), "")
	}

	announcement, ok := meta["announcement"].(map[any]any)
	if ok {
		c.Announcement.Text = cmp.Or(announcement["text"].(string), "")
		c.Announcement.Link = cmp.Or(announcement["link"].(string), "")
	}

	qlinks, ok := meta["quick_links"].([]any)
	if ok {
		for _, v := range qlinks {
			link := v.(map[any]any)
			c.QuickLinks = append(c.QuickLinks, struct {
				Text string
				Link string
				Icon string
			}{
				Text: cmp.Or(link["text"].(string), ""),
				Link: cmp.Or(link["link"].(string), ""),
				Icon: cmp.Or(link["icon"].(string), ""),
			})
		}
	}

	elinks, ok := meta["external_links"].([]any)
	if ok {
		for _, v := range elinks {
			link := v.(map[any]any)
			c.ExternalLinks = append(c.ExternalLinks, struct {
				Text string
				Link string
			}{
				Text: cmp.Or(link["text"].(string), ""),
				Link: cmp.Or(link["link"].(string), ""),
			})
		}
	}

	return c, nil
}
