---
title: Getting Started
index: 2
---
Doco is delivered as a self-contained binary, making its usage and integration straightforward. The first step in utilizing Doco's full potential is to install this binary, which enables access to all its features, from converting markdown files to integrating with continuous integration workflows. This binary-focused approach ensures a streamlined and efficient user experience.

## Installing
To install Doco you can do it with the go tool or download the binary from the releases page.

```go
go install github.com/paganotoni/doco/cmd/doco@latest
```

## Initializing the documentation folder
Once the tool is installed you can run the following command to initialize the documentation folder:

```sh
doco init
```

This, will create the `docs` folder with a few files:

- _meta.md
- index.md
- getting_started/getting_started.md


## Browse the documentation

Once the folder is initialized you can run the following command to generate the static website:

```sh
doco serve
```

This will generate the static website in the `public` folder. and serve it. You can see your docs site at http://localhost:3000. Once you're done you can stop the server with `ctrl+c`. 


