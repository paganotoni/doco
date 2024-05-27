----
# Name of the website, is used on the pages description and 
# In the meta tags.
name: Doco
description: "Doco is a CLI tool to generate static documentation websites from markdown files."
keywords: "keywords,for,seo"

# Index defines where the pages in the root will be positioned.
index: -1
favicon: "/assets/favicon.png"

logo: 
  src: "/assets/logo.png"
  link: "/"


# Announcement is shown on the top of the site
# next to the site logo
announcement: 
  text: "Check our Github repository."
  link: "https://github.com/paganotoni/doco"

# Social links are shown on the top right of the site
# next to the CTA button.
github: "https://github.com/paganotoni/doco" 

# External links go in the top navigation bar.
external_links:
  - text: "Documentation"
    link: "/"

# QuickLinks go on top of the left navigation bar and
# show on the quick search modal by default.
quick_links:
  - text: "Documentation"
    link: "/"
    # Icons use google material icons
    icon: "menu_book"
  - text: "Repository"
    link: "https://github.com/paganotoni/doco"
    icon: "code"


# CTA button shows on the top right of the site.
# cta: 
#  text: "Star on Github"
#  link: "https://github.com/paganotoni/doco"

# Shows on each of the documentation pages
# $YEAR is replaced with the current year.
copy: "© $YEAR Doco"
----