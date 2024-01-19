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

The `--folder` parameter is optional and it will be used to set the name of the documentation folder. If the folder is not provided, it will default to `docs`. The `--output` parameter is optional and it will be used to set the name of the output folder. If the output folder is not provided, it will default to `public`.