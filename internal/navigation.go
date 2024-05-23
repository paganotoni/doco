package internal

import (
	"path/filepath"
	"strings"

	. "github.com/delaneyj/gostar/elements"
)

// buildNavigation generates the html for the desktop navigation
// it is used in the template.html file.
func buildNavigation(s *site, doc document) ElementRenderer {
	return NAV().CLASS("documents").Children(
		Range(s.sections, func(s section) ElementRenderer {
			return SECTION().IfChildren(s.name != "", H3().Text(s.name)).Children(
				UL().Children(
					Range(s.documents, func(d document) ElementRenderer {
						link := "/" + filepath.Join(s.path, strings.TrimSuffix(d.filename, ".md")+".html")
						return LI().IfCLASS(doc.filename == d.filename, "active").Children(
							A().HREF(link).Text(d.title),
						)
					}),
				),
			)
		}),
	)
}
