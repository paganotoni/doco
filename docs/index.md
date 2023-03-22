---
Title: "Home"
Weight: 0
---
# Welcome to Doco

Doco is a documentation website generator for your projects written in Go. It takes a folder with markdown files and generates a deployable static website.

## Installation

```sh
go install github.com/paganotoni/doco@latest
```
## Usage

The Doco CLI has a few commands that will help you generate the static website from your markdown files. To get started you can run the following command:

```sh
doco init
```
This will create the `docs` folder with a few basic files:

- `_meta.md`: This file contains the metadata of your documentation. 
- `index.md`: This is the home page of your documentation.
- `getting-started.md`: This is the getting started page of your documentation.

### Metadata
Beyond the mardown content Doco allows to specify some metadata for each page. The metadata is used to generate the navigation and the title of the pages.
#### Meta fields

Within each markdown file you can add metadata fields. The metadata fields are used to generate the navigation and the title of the pages.

```md
---
Title: Getting Started
Weight: 1
---
```

##### Title
Title allows to define the title of the page. It is used in the title tag of the page and the navigation.

##### Weight
Weight allows to define the order of the page in the navigation. The pages are sorted by weight in the navigation.
