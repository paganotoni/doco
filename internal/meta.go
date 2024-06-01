package internal

import (
	yaml "gopkg.in/yaml.v3"
)

func parseMeta(doc []byte) (map[string]interface{}, error) {
	t := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(doc), &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
