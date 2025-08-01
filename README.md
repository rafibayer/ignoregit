# **ignoregit**: A command-line utility to create .gitignore files

This CLI embeds the contents of https://github.com/github/gitignore to allow quick creation of .gitignore files from GitHub-maintained templates.

## Installation
```shell
# From VCS
go install github.com/rafibayer/ignoregit@latest

# From Source
make install
```
## Usage

```shell
# list available templates:
# --all flag to include community+global templates
ignoregit list 

# create new .gitignore from one or more templates:
# --out flag to override default output path: .gitignore, --stdout flag to output to stdout
ignoregit new [language...]  instead
```

## How It Works
The `repo` Makefile target has 3 parts:
1. `fetch` downloads the contents of `github.com/github/gitignore` and unzips it into this repository under `source/repo`

2. `flatten` traverses symlinks in `github.com/github/gitignore` and "flattens" them by replacing them with the symlink target content.
This is done because `ignoregit` utilizes [`embed`](https://pkg.go.dev/embed) which does not support irregular files.

3. `clean` removes unused content like the `.github` folder.

`source/repo` is also a go package. This package embeds the flattened and cleaned contents of `github.com/github/gitignore` as a filesystem - which is then walked by the `ignoregit` CLI to list available templates and copy them into new files. 

## Get Help
```shell
Usage:
  ignoregit [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        list available .gitignore templates
  new         create a new .gitignore

Flags:
  -h, --help   help for ignoregit

Use "ignoregit [command] --help" for more information about a command.
```

## TODO
- json output
- namespacing global/community templates
- search / aliases for templates