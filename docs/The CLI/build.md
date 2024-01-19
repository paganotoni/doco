---
title: Build
index: 2
---

Build generates the static documentation site from the markdown files in the `docs` folder. It will create a `public` folder with the static website.

```sh
$ doco build
# or a different folder name
$ doco --folder [folder] build
# or a different output folder
$ doco --output [output] build
```

When the docs folder does not exist, it stops the operation and error out with:

```sh
ERROR: docs folder does not exist, aborting build
```