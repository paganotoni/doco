---
index: 0
---

The root `_meta.md` file contains the configuration for the global documentation features. By default, it contains the following:

```markdown
----
# Name of the website, is used on the pages description and 
# In the meta tags.
name: Doco
description: "Doco is a CLI tool to generate static documentation websites from markdown files."
keywords: "keywords,for,seo"

# Index defines where the pages in the root will be positioned.
index: -1

logo: "/assets/logo.png"
favicon: "/assets/favicon.png"


# Announcement is shown on the top of the site
# next to the site logo
announcement: 
  text: "Check our Github repository"
  link: "https://github.com/paganotoni/doco"

# Social links are shown on the top right of the site
github: "https://github.com/paganotoni/doco" 

# External links go in the top navigation bar.
external_links:
  - text: "Documentation"
    link: "/"

# QuickLinks go on top of the left navigation bar and
# show on the quick search modal by default. Icons use google material icons
# so you can there use any of the icon names from https://fonts.google.com/icons.
quick_links:
  - text: "Documentation"
    link: "/"
    icon: "menu_book"
  - text: "Repository"
    link: "https://github.com/paganotoni/doco"
    icon: "code"

# Shows on each of the documentation pages
# $YEAR is replaced with the current year.
copy: "© $YEAR Doco"
----
```
