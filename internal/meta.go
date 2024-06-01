package internal

import (
	"regexp"

	yaml "gopkg.in/yaml.v3"
)

func parseMeta(doc []byte) (t map[string]interface{}, err error) {
	// Extracting the YAML front matter froim the markdown document
	rg := regexp.MustCompile(`(?s)^---\n(.*?)\n---\n`)
	m := rg.FindSubmatch(doc)
	if len(m) < 2 {
		return nil, nil
	}

	err = yaml.Unmarshal([]byte(m[1]), &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
