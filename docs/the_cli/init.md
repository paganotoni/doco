---
title: Init
index: 1
---

Init creates the `docs` folder with some basic files.

```sh
$ doco init 
# or a different folder name
$ doco --folder [folder] init 
```

When the docs folder already exists, it will stop the operation and error out with:

```sh
ERROR: folder exists, aborting init
```

The `name` parameter is optional and it will be used to set the name of the documentation site. The `init` command creates the following files:

```sh
docs
├── _meta.md
├── assets
│   ├── logo.png
│   └── favicon.png
├── index.md
└── getting_started.md
```
