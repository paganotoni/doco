package web

import (
	"html/template"
	"sort"
)

// navigation is a slice of Page that implements the htmlFor method
// to generate the HTML for the navigation.
type navigation Pages

func (gn navigation) HTMLFor(page Page) template.HTML {
	sort.Sort(gn)

	result := `<ul>`
	for _, page := range gn {
		if page.SourcePath == "docs/_meta.md" {
			continue
		}

		result += `<li><a href="` + page.resultLink() + `">` + page.title() + `</a></li>`
	}

	result += `</ul>`

	return template.HTML(result)
}

func (gn navigation) Len() int {
	return len(gn)
}

func (gn navigation) Less(i, j int) bool {
	return gn[i].weight() < gn[j].weight()
}

func (gn navigation) Swap(i, j int) {
	gn[i], gn[j] = gn[j], gn[i]
}
