package internal

import (
	"html/template"
	"path/filepath"
	"strings"
)

// desktopNavigation generates the html for the desktop navigation
// it is used in the template.html file.
func desktopNavigation(s *site, doc document) template.HTML {
	var html string
	html += `<nav class="documents">`
	for _, v := range s.sections {
		if v.name != "" {
			html += `<h3>` + v.name + `</h3>`
		}

		html += `<ul>`
		for _, ss := range v.documents {
			var class string
			if doc.filename == ss.filename {
				class = "active"
			}

			link := "/" + filepath.Join(v.path, strings.TrimSuffix(ss.filename, ".md")+".html")
			html += `<li class="` + class + `"><a href="` + link + `">` + ss.title + `</a></li>`
		}

		html += `</ul>`
	}

	html += `</nav>`

	return template.HTML(html)
}
