package internal

import (
	_ "embed"
	"html/template"
	"io"

	. "github.com/delaneyj/gostar/elements"
	"github.com/paganotoni/doco/internal/config"
)

var (
	//go:embed assets/doco.css
	style []byte

	//go:embed assets/doco.js
	docoJS []byte
)

type navlink struct {
	Title string `json:"-"`
	Link  string `json:"-"`
}

// generatedPage is the data passed to the template
// to generate the static html files.
type generatedPage struct {
	filePath   string          `json:"-"`
	Prev       navlink         `json:"-"`
	Next       navlink         `json:"-"`
	Navigation ElementRenderer `json:"-"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`

	SectionName string        `json:"section_name"`
	Content     template.HTML `json:"-"`
	Link        string        `json:"link"`

	Tokens string `json:"content"`
}

func (g generatedPage) html(s config.Site, w io.Writer) error {
	content := Group(
		MAIN().CLASS("flex-grow lg:px-5 px-5 py-5 text-md lg:text-lg pb-10").IfChildren(
			g.SectionName != "",
			SPAN().CLASS("text-sm").Text(g.SectionName),
		).Children(
			H1().ID("welcome").CLASS("font-bold text-4xl mb-2").Text(g.Title),
			DIV().ID("htmlcontainer").CLASS("max-w-5xl").Text(string(g.Content)),

			DIV().CLASS("grid grid-cols-2 gap-2 justify-between text-gray-600 pt-8").Children(
				DIV().IfChildren(
					g.Prev.Link != "",
					A().HREF("/"+g.Prev.Link).CLASS("p-3.5 px-5 flex border rounded-lg hover:shadow").Children(
						SPAN().CLASS("flex flex-col").Children(
							SPAN().CLASS("text-base").Text("Previous "),
							SPAN().CLASS("text-lg flex flex-row gap-3 font-bold").Children(
								SPAN().Text(
									`<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                             			<path stroke-linecap="round" stroke-linejoin="round" d="M10.5 19.5 3 12m0 0 7.5-7.5M3 12h18" />
                                   	</svg>`,
								),
								SPAN().Text(g.Prev.Title),
							),
						),
					),
				),

				DIV().IfChildren(
					g.Next.Link != "",
					A().HREF("/"+g.Next.Link).CLASS("p-3.5 px-5 border rounded-lg flex flex-row flex-row-reverse hover:shadow").Children(
						SPAN().CLASS("flex flex-col text-right").Children(
							SPAN().CLASS("text-base").Text("Next "),
							SPAN().CLASS("text-lg flex flex-row gap-2 font-bold").Children(
								SPAN().Text(g.Next.Title),
								SPAN().Text(
									`<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                             			<path stroke-linecap="round" stroke-linejoin="round" d="M13.5 4.5 21 12m0 0-7.5 7.5M21 12H3" />
                               		</svg>`,
								),
							),
						),
					),
				),
			),
		),
	)

	return page(s, g, content).Render(w)
}

// page generates the html page with the given content.
func page(s config.Site, g generatedPage, content ElementRenderer) ElementRenderer {
	return HTML().LANG("en").CLASS("h-full").Children(
		HEAD().Children(
			TITLE().TextF("%s - %s", s.Name, g.Title),

			META().CHARSET("utf-8"),
			META().NAME("viewport").CONTENT("width=device-width, initial-scale=1"),
			META().NAME("description").CONTENT(g.Description),
			META().NAME("keywords").CONTENT(g.Keywords),

			LINK().REL("icon").TYPE("image/png").HREF(s.Favicon),
			LINK().HREF("https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:wght@100").REL("stylesheet"),
			LINK().REL("stylesheet").HREF("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/default.min.css"),

			SCRIPT().SRC("https://cdn.tailwindcss.com"),
			STYLE().TYPE("text/tailwindcss").Text(string(style)),
		),
		BODY().CLASS("bg-gray-50 max-w-[1400px] mx-auto flex flex-col min-h-full").Children(
			HEADER().CLASS("sticky top-0 z-40 w-full px-3 border-gray-200 bg-gray-50 border-b flex flex-row gap-8 py-4 items-center").Children(
				SPAN().IfChildren(
					s.Announcement.Text != "",
					IMG().CLASS("h-7").SRC(s.Logo.ImageSrc),
				).IfChildren(
					s.Announcement.Text == "",
					A().HREF("/").CLASS("font-bold text-xl flex flex-row gap-1 items-center").Children(
						SPAN().Text(
							`<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                    			<path stroke-linecap="round" stroke-linejoin="round" d="M12 6.042A8.967 8.967 0 0 0 6 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 0 1 6 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 0 1 6-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0 0 18 18a8.967 8.967 0 0 0-6 2.292m0-14.25v14.25" />
                       		</svg>`,
						),
						SPAN().Text("Doco"),
					),
				),
				// Announcement banner
				DIV().CLASS("flex items-center gap-8").IfChildren(
					s.Announcement.Text != "",
					A().HREF(s.Announcement.Link).CLASS("hidden md:flex rounded-full p-3 py-1.5 bg-white border text-sm flex flex-row items-center gap-2 hover:border-gray-300").Children(
						SVG_SVG().CLASS("w-5 h-5").VIEW_BOX("0 0 24 24").Attr("stroke-width", "1.2").Attr("stroke", "currentColor").Children(
							SVG_PATH().Attr("stroke-linecap", "round").Attr("stroke-linejoin", "round").D("M10.34 15.84c-.688-.06-1.386-.09-2.09-.09H7.5a4.5 4.5 0 1 1 0-9h.75c.704 0 1.402-.03 2.09-.09m0 9.18c.253.962.584 1.892.985 2.783.247.55.06 1.21-.463 1.511l-.657.38c-.551.318-1.26.117-1.527-.461a20.845 20.845 0 0 1-1.44-4.282m3.102.069a18.03 18.03 0 0 1-.59-4.59c0-1.586.205-3.124.59-4.59m0 9.18a23.848 23.848 0 0 1 8.835 2.535M10.34 6.66a23.847 23.847 0 0 0 8.835-2.535m0 0A23.74 23.74 0 0 0 18.795 3m.38 1.125a23.91 23.91 0 0 1 1.014 5.395m-1.014 8.855c-.118.38-.245.754-.38 1.125m.38-1.125a23.91 23.91 0 0 0 1.014-5.395m0-3.46c.495.413.811 1.035.811 1.73 0 .695-.316 1.317-.811 1.73m0-3.46a24.347 24.347 0 0 1 0 3.46"),
						),

						SPAN().Text(s.Announcement.Text),
						SVG_SVG().CLASS("w-3 h-3").VIEW_BOX("0 0 24 24").Attr("stroke-width", "1.2").Attr("stroke", "currentColor").Children(
							SVG_PATH().Attr("stroke-linecap", "round").Attr("stroke-linejoin", "round").D("M8.25 4.5 15.75 12 8.25 19.5"),
						),
					),
				),

				// External Links
				NAV().CLASS("flex-grow flex justify-end").Children(
					UL().CLASS("flex flex-row gap-7 font-medium hidden lg:flex border-r").Children(
						Range(s.ExternalLinks, func(link config.Link) ElementRenderer {
							return LI().Children(
								A().HREF(link.Link).CLASS("hover:underline flex flex-row gap-1 items-center mr-8").Children(
									SPAN().Text(link.Text),
								),
							)
						}),
					),
				),

				// Search button
				DIV().CLASS("md:hidden flex flex-row gap-4 items-center").Children(
					SPAN().CLASS("search-button").Children(
						SVG_SVG().CLASS("w-6 h-6").VIEW_BOX("0 0 24 24").Attr("stroke-width", "1.5").Attr("stroke", "currentColor").Children(
							SVG_PATH().Attr("stroke-linecap", "round").Attr("stroke-linejoin", "round").D("m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"),
						),

						SPAN().CLASS("toggle-mobile-nav").Children(
							SVG_SVG().CLASS("w-8 h-8").VIEW_BOX("0 0 24 24").Attr("stroke-width", "2").Attr("stroke", "currentColor").Children(
								SVG_PATH().Attr("stroke-linecap", "round").Attr("stroke-linejoin", "round").D("M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"),
							),
						),
					),
				),
			).IfChildren(
				s.Github != "",

				DIV().CLASS("flex items-center gap-4").Children(
					A().HREF(s.Github).CLASS("hidden md:flex rounded-full p-3 py-1.5 bg-white border text-sm flex flex-row items-center gap-2 hover:border-gray-300").Children(
						SVG_SVG().CLASS("w-5 h-5").VIEW_BOX("0 0 24 24").Attr("stroke-width", "1.2").Attr("stroke", "currentColor").Children(
							SVG_PATH().Attr("stroke-linecap", "round").Attr("stroke-linejoin", "round").D("M10.34 15.84c-.688-.06-1.386-.09-2.09-.09H7.5a4.5 4.5 0 1 1 0-9h.75c.704 0 1.402-.03 2.09-.09m0 9.18c.253.962.584 1.892.985 2.783.247.55.06 1.21-.463 1.511l-.657.38c-.551.318-1.26.117-1.527-.461a20.845 20.845 0 0 1-1.44-4.282m3.102.069a18.03 18.03 0 0 1-.59-4.59c0-1.586.205-3.124.59-4.59m0 9.18a23.848 23.848 0 0 1 8.835 2.535M10.34 6.66a23.847 23.847 0 0 0 8.835-2.535m0 0A23.74 23.74 0 0 0 18.795 3m.38 1.125a23.91 23.91 0 0 1 1.014 5.395m-1.014 8.855c-.118.38-.245.754-.38 1.125m.38-1.125a23.91 23.91 0 0 0 1.014-5.395m0-3.46c.495.413.811 1.035.811 1.73 0 .695-.316 1.317-.811 1.73m0-3.46a24.347 24.347 0 0 1 0 3.46"),
						),
					),
				),
			),
			SECTION().CLASS("flex-grow flex flex-row px-3").Children(
				ASIDE().CLASS("hidden lg:block lg:fixed min-w-[19rem] pr-5 flex flex-col gap-1 pt-6").Children(
					BUTTON().CLASS("search-button border min-w-[17rem] rounded-lg p-2 px-3 mb-6 bg-gray-50 text-left flex flex-row hover:border-gray-400 hover:bg-white items-center").Text(
						`<svg width="24" height="24" fill="none" aria-hidden="true" class="mr-3 flex-none">
                    		<path d="m19 19-3.5-3.5" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path><circle cx="11" cy="11" r="6" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></circle>
                      	</svg>`,
					).Children(
						SPAN().Text("Quick Search"),
						SPAN().CLASS("ml-auto pl-3 flex-none text-sm font-semibold").Text("⌘K"),
					),

					DIV().CLASS("max-h-[calc(100vh-180px)] overflow-y-scroll pb-10").Children(
						UL().CLASS("flex flex-col mb-8").Children(
							Range(s.QuickLinks, func(l config.Link) ElementRenderer {
								return LI().Children(
									A().HREF(l.Link).CLASS("flex flex-row gap-2 p-1.5 rounded-lg hover:bg-gray-200/80").Children(
										SPAN().CLASS("material-symbols-outlined").Text(l.Icon),
										SPAN().Text(l.Text),
									),
								)
							}),
						),

						NAV().ID("desktop-navigation").Children(
							g.Navigation,
						),
					),
				),

				SECTION().CLASS("flex-grow flex flex-col lg:ml-[19rem]").Children(
					content,

					FOOTER().CLASS("px-5 py-8 items-center flex flex-row text-gray-400 border-t").Children(
						SPAN().CLASS("flex-grow text-sm").ID("copy").Text(s.Copy),
						A().HREF("https://doco.sh").CLASS("text-sm underline flex flex-row items-center gap-1").TARGET("_blank").Children(
							SPAN().Text("Powered By"),
							STRONG().CLASS("font-bold").Text("Doco"),

							SPAN().Text(
								`<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
			         				<path stroke-linecap="round" stroke-linejoin="round" d="M13.5 6H5.25A2.25 2.25 0 0 0 3 8.25v10.5A2.25 2.25 0 0 0 5.25 21h10.5A2.25 2.25 0 0 0 18 18.75V10.5m-10.5 6L21 3m0 0h-5.25M21 3v5.25" />
			              		</svg>`,
							),
						),
					),
				),
			),

			DIV().CLASS("hidden fixed inset-0 z-[100] overflow-y-auto p-4 sm:p-6 md:p-20").ID("search-palette").ROLE("dialog").Attr("aria-modal", "true").Children(
				DIV().CLASS("fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity").ID("search-overlay").Attr("aria-hidden", "true"),
				DIV().CLASS("z-[100] mx-auto max-w-2xl transform divide-y divide-gray-100 overflow-hidden rounded-xl bg-white shadow-2xl ring-1 ring-black ring-opacity-5 transition-all").Children(
					DIV().CLASS("relative").Children(
						SVG_SVG().CLASS("pointer-events-none absolute top-3.5 left-4 h-5 w-5 text-gray-400").Attr("xmlns", "http://www.w3.org/2000/svg").VIEW_BOX("0 0 20 20").Attr("fill", "currentColor").Attr("aria-hidden", "true").Children(
							SVG_PATH().Attr("fill-rule", "evenodd").D("M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z").Attr("clip-rule", "evenodd"),
						),

						INPUT().ID("search-input").TYPE("text").CLASS("h-12 w-full border-0 bg-transparent pl-11 pr-4 text-gray-800 placeholder-gray-400 focus:ring-0 outline-none sm:text-sm").PLACEHOLDER(""),
						BUTTON().ID("close-search").CLASS("bg-gray-200 h-8 w-8 text-xs text-gray-600 rounded-md inline-block absolute right-2 top-2").Text("esc"),
					),

					UL().ID("search-quick-actions").CLASS("max-h-80 scroll-py-2 divide-y divide-gray-100 overflow-y-auto").Children(
						LI().CLASS("p-2").Children(
							H2().CLASS("sr-only").Text("Quick actions"),
							UL().CLASS("text-sm text-gray-700").Children(
								Range(s.QuickLinks, func(l config.Link) ElementRenderer {
									return LI().CLASS("group cursor-default select-none items-center rounded-md px-3 py-2").Children(
										A().HREF(l.Link).CLASS("flex flex-row items-center group hover:text-blue-500").Children(
											SPAN().CLASS("material-symbols-outlined").Text(l.Icon),
											SPAN().CLASS("ml-3 flex-auto truncate").Text(l.Text),
											SVG_SVG().CLASS("w-6 h-6").Attrs("fill", "none", "stroke", "currentColor").VIEW_BOX("0 0 24 24").Attr("xmlns", "http://www.w3.org/2000/svg").Children(
												SVG_PATH().Attr("stroke-linecap", "round").Attr("stroke-linejoin", "round").Attr("stroke-width", "2").D("M17 8l4 4m0 0l-4 4m4-4H3"),
											),
										),
									)
								}),
							),
						),
					),

					UL().ID("search-results").CLASS("hidden max-h-96 overflow-y-auto p-2 text-sm text-gray-700"),
					DIV().ID("search-no-results").CLASS("hidden py-14 px-6 text-center sm:px-14").Text(
						`<svg class="mx-auto h-6 w-6 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                    				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"></path>
                        		 </svg>`,
					).Children(
						P().CLASS("mt-4 text-sm text-gray-900"),
					),
				),
			),

			// Search Result Template
			SCRIPT().ID("search-result-template").TYPE("text/x-js-template").Text(
				`<li class="text-gray-700 group cursor-default select-none items-center rounded-md px-3 py-2 hover:text-blue-500">
		            <a class="flex items-center" href="/${link}">
		                <svg class="h-6 w-6 flex-none" fill="none" stroke="currentColor" viewbox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path></svg>
		                <span class="ml-3 flex-grow flex-auto truncate">${title}</span>
		                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewbox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3"></path></svg>
		                <span class="ml-3 hidden flex-none text-indigo-100">Jump to...</span>
		            </a>
		          </li>`,
			),

			// Mobile menu
			NAV().ID("mobile-menu").CLASS("hidden overflow-scroll bg-gray-400/90 z-50 fixed right-0 left-0 top-0 bottom-0 h-screen w-screen").Children(
				SPAN().CLASS("toggle-mobile-nav fixed top-10 right-10").Children(
					SVG_SVG().CLASS("w-6 h-6").Attr("xmlns", "http://www.w3.org/2000/svg").Attr("fill", "none").VIEW_BOX("0 0 24 24").Attr("stroke-width", "2").Attr("stroke", "currentColor").Children(
						SVG_PATH().Attr("stroke-linecap", "round").Attr("stroke-linejoin", "round").D("M6 18 18 6M6 6l12 12"),
					),
				),
				DIV().CLASS("bg-white max-w-[calc(100vw-100px)] py-6 px-4").Children(
					UL().CLASS("quicklinks flex flex-col mb-8").Children(
						Range(s.QuickLinks, func(i config.Link) ElementRenderer {
							return LI().Children(
								A().HREF(i.Link).CLASS("flex flex-row gap-2 p-1.5 rounded-lg hover:bg-gray-200/80").Children(
									SPAN().CLASS("material-symbols-outlined").Text(i.Icon),
								).Text(i.Text),
							)
						}),
					),
				),
			),

			// Zoomed image overlay
			DIV().ID("zoomed-image-overlay").CLASS("hidden fixed bg-white/90 top-0 right-0 w-full h-full z-50 text-center px-[15%] py-[5%] cursor-zoom-out").Children(
				DIV().CLASS("w-full").Children(
					IMG().SRC("https://placehold.co/600x400").ALT("image").CLASS("w-full bg-white rounded-md place-content-center shadow-xl"),
				),
			),

			// General Scripts
			SCRIPT().SRC("https://cdn.jsdelivr.net/npm/fuse.js@7.0.0"),
			SCRIPT().SRC("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js"),
			SCRIPT().SRC("/doco.js"),
		),
	)
}
