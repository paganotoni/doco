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
func readConfig(folder string) (config, error) {
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

	config.Description = data["description"].(string)
	config.Keywords = data["keywords"].(string)
	config.Copy = data["copy"].(string)
	config.Github = data["github"].(string)
	config.Favicon = data["favicon"].(string)

	logo, ok := data["logo"].(map[any]any)
	if ok {
		config.Logo.Src = logo["src"].(string)
		config.Logo.Link = logo["link"].(string)
	}

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
