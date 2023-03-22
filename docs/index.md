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

### _meta.md

The meta file contains the metadata of your documentation. The metadata is used to generate the navigation and the title of the pages.

```md
---
Name: Doco
Logo: https://raw.githubusercontent.com/paganotoni/doco/main/docs/assets/logo.png
---
```

This is useful to have a single source of truth for the metadata of your documentation.
### Build
The build command will generate the static website from the markdown files in the `docs` folder. It will create a `public` folder with the static website.

### Metadata

