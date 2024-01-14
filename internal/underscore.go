package internal

import (
	"strings"
	"unicode"
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
