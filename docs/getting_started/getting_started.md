---
title: Getting Started
---

Doco is a documentation website generator for your projects written in Go. It takes a folder with markdown files and generates a deployable static website.

## Features

- Markdown based
- Generated Navigation
- Responsive Website
- Syntax highlighting
- Quick Search widget

## Installation
To install Doco you can do it with the go tool or download the binary from the [releases page]().

### Install using Go

```go
go install github.com/paganotoni/doco@latest
```
### Download binary
You can download the binary from the [releases page]() and pick the one that matches your OS and architecture.

## Usage
Once the tool is installed you can run the following command to initialize the documentation folder:

```sh
doco init
```

This, based on the `docs` folder with a few basic files:

- _meta.md
- getting_started/getting_started.md
- getting_started/development.md
- getting_started/faq.md

Once the folder is initialized you can run the following command to generate the static website:

```sh
doco build
```

This will generate the static website in the `site` folder. You can then deploy the static website to your favorite hosting provider.

### the `_meta` file

The `_meta` file is a markdown file that contains the metadata for the documentation website. It is used to generate the header links as well as configuring general features such as the search widget.

### markdown files

Each of the markdown files in the `docs` folder will be a page in the documentation website. The folder structure will be used to generate the navigation. However, you can configure the name of these files by adding a `title` metadata to the markdown file. You can also use the `weight` metadata to configure the order of the pages in the navigation menu. Files can be excluded from the navigation by using the `draft` property on them.

A typical markdown file will look like this:

```md
--- 
title: Getting Started
weight: 1
draft: false
---

# Getting Started
...
```