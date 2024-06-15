---
title: Getting Started
index: 2
description: Learn how to get started with Doco, from installation to initializing the documentation folder.
---

Doco is delivered as a self-contained binary, making its usage and integration straightforward. The first step in utilizing Doco's full potential is to install this binary, which enables access to all its features, from converting markdown files to integrating with continuous integration workflows. This binary-focused approach ensures a streamlined and efficient user experience.

## Installing
To install download the binary from the releases page. You can find the latest release [here](https://github.com/paganotoni/doco/releases).

Below are some examples on how to download depending on your OS and architecture.

### On Mac
```sh
$ curl -OL https://github.com/paganotoni/doco/releases/latest/doco_Darwin_arm64.tar.gz
$ tar -xvzf doco_Darwin_arm64.tar.gz
$ mv doco /usr/local/bin/doco
```

### On Linux
```sh
$ wget https://github.com/paganotoni/doco/releases/latest/doco_Linux_x86_64.tar.gz
$ tar -xvzf doco_Linux_x86_64.tar.gz
$ sudo mv doco /usr/local/bin/doco
```

Doco builds the following OSs and architectures on each release:

- Darwin_arm64
- Darwin_x86_64
- Linux_arm64
- Linux_armv6
- Linux_armv7
- Linux_i386
- Linux_x86_64
- Windows_arm64
- Windows_armv6
- Windows_armv7
- Windows_i386
- Windows_x86_64

### Installing from source
Alternatively you have Go installed in your system you can also use Go to install the Doco binary.

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

This will generate the static website in the `public` folder. and serve it. You can see your docs site at [http://localhost:3000/](http://localhost:3000/). Once you're done you can stop the server with `ctrl+c`.
