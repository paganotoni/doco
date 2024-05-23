package internal

import (
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func underscore(s string) string {
	var result string
	for i, v := range s {
		if i > 0 && unicode.IsUpper(v) {
			result += "_"
		}
		result += strings.ToLower(string(v))
	}
	return result
}

func humanizeFilename(path string) string {
	// remove the extension
	s := strings.ReplaceAll(path, filepath.Ext(path), "")
	// remove the path
	s = filepath.Base(s)

	return humanize(s)
}

func humanize(s string) string {
	// remove underscores and dashes
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	s = cases.Title(language.English).String(s)

	return s
}
