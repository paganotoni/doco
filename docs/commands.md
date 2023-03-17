# CLI

The Doco CLI is the one in charge of generating the static website from the markdown files. It can come very handy when writing the documentation for your project since it will allow you to see the changes in real time.

## Installation

To install the Doco CLI you can run the following command:

```sh
go install github.com/paganotoni/doco@latest
```

## Usage

The Doco CLI has a few commands that will help you generate the static website from your markdown files.

### Build
Build generates the static documentation site from the markdown files in the `docs` folder. It will create a `public` folder with the static website.


```sh
$ doco build
```
### Init

Init will create the `docs` folder with some basic files.

```sh
$ doco init
```

The build argument is optional, when the CLI is invoked without argument it performs the `build` command.

### Serve
Serve builds the static website and starts a local server to serve the static website on port 3004. It will also watch for changes in the markdown files and rebuild the website.

```sh
$ doco serve
[INFO] Watching for changes
[INFO] Building site
[INFO] Write > index.html
[INFO] Write > getting-started.html
[INFO] Write > commands.html
[INFO] Serving on http://localhost:3004/
[INFO] Changes detected
[INFO] Server Stopped
[INFO] Building site
[INFO] Write > index.html
[INFO] Write > getting-started.html
[INFO] Write > commands.html
[INFO] Serving on http://localhost:3004/

```

### Version
Version displays the current version of the CLI.

```sh
$ doco version
v1.0.0
```

### Help

The help command prints usage information about the CLI.

```sh
$ doco help
Usage:
  doco [command]
...
```





