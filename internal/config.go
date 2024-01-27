package internal

import (
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

	data, err := metadataFrom(content)
	if err != nil {
		return c, err
	}

	c.Name = stringWithDefault(data["name"], "Doco")
	c.Description = stringWithDefault(data["description"], "Documentation site")
	c.Keywords = stringWithDefault(data["keywords"], "documentation, site, doco")
	c.Copy = stringWithDefault(data["copy"], "© $YEAR Doco")
	c.Github = stringWithDefault(data["github"], "https://github.com/paganotoni/doco")
	c.Favicon = stringWithDefault(data["favicon"], "")

	logo, ok := data["logo"].(map[any]any)
	if ok {
		c.Logo.Src = stringWithDefault(logo["src"], "")
		c.Logo.Link = stringWithDefault(logo["link"], "")
	}

	announcement, ok := data["announcement"].(map[any]any)
	if ok {
		c.Announcement.Text = stringWithDefault(announcement["text"], "")
		c.Announcement.Link = stringWithDefault(announcement["link"], "")
	}

	qlinks, ok := data["quick_links"].([]any)
	if ok {
		for _, v := range qlinks {
			link := v.(map[any]any)
			c.QuickLinks = append(c.QuickLinks, struct {
				Text string
				Link string
				Icon string
			}{
				Text: stringWithDefault(link["text"], ""),
				Link: stringWithDefault(link["link"], ""),
				Icon: stringWithDefault(link["icon"], ""),
			})
		}
	}

	elinks, ok := data["external_links"].([]any)
	if ok {
		for _, v := range elinks {
			link := v.(map[any]any)
			c.ExternalLinks = append(c.ExternalLinks, struct {
				Text string
				Link string
			}{
				Text: stringWithDefault(link["text"], ""),
				Link: stringWithDefault(link["link"], ""),
			})
		}
	}

	return c, nil
}

// stringWithDefault allows to parse the value of a key
// and return a default value if the key is not present or
// the value is not a string.
func stringWithDefault(val any, def string) string {
	ss, ok := val.(string)
	if ok {
		return ss
	}

	return def
}
