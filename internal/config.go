package internal

import (
	"io"
	"os"
	"path/filepath"
)

const (
	metafile = "_meta.md"
)

// config of the site.
type config struct {
	Name        string
	Logo        string
	Favicon     string
	Description string
	Keywords    string

	Announcement struct {
		Text string // Text to display, empty means no announcement
		Link string
	}

	Github     string // Github link to display, empty means no link
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

// parseMeta parses the _meta.md file and returns the config
// for the site.
func parseMeta(folder string) (config, error) {
	file, err := os.Open(filepath.Join(folder, metafile))
	if err != nil {
		return config{}, err
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return config{}, err
	}

	data, err := metadataFrom(content)
	if err != nil {
		return config{}, err
	}

	config := config{}
	config.Name = data["name"].(string)
	config.Logo = data["logo"].(string)
	config.Description = data["description"].(string)
	config.Keywords = data["keywords"].(string)
	config.Copy = data["copy"].(string)
	config.Github = data["github"].(string)
	config.Favicon = data["favicon"].(string)

	announcement := data["announcement"].(map[any]any)
	config.Announcement.Text = announcement["text"].(string)
	config.Announcement.Link = announcement["link"].(string)

	quickLinks := data["quick_links"].([]any)
	for _, v := range quickLinks {
		link := v.(map[any]any)
		config.QuickLinks = append(config.QuickLinks, struct {
			Text string
			Link string
			Icon string
		}{
			Text: link["text"].(string),
			Link: link["link"].(string),
			Icon: link["icon"].(string),
		})
	}

	externalLinks := data["external_links"].([]any)
	for _, v := range externalLinks {
		link := v.(map[any]any)
		config.ExternalLinks = append(config.ExternalLinks, struct {
			Text string
			Link string
		}{
			Text: link["text"].(string),
			Link: link["link"].(string),
		})
	}

	return config, nil
}
