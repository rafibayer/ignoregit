# **ignoregit**: A command-line utility to create .gitignore files

This CLI embeds the contents of https://github.com/github/gitignore to allow quick creation of .gitignore files from GitHub-maintained templates.

## Installation
```shell
go install github.com/rafibayer/ignoregit@latest
```
## Usage

```shell
# list available templates:
ignoregit list # --all flag to include community+global templates

# create new .gitignore
ignoregit new <template> # --out flag to override default output path: .gitignore
```

## How It Works
The `repo` Makefile target has 3 parts:
1. `fetch` downloads the contents of `github.com/github/gitignore` and unzips it into this repository under `source/repo`

2. `flatten` traverses symlinks in `github.com/github/gitignore` and "flattens" them by replacing them with the symlink target content.
This is done because `ignoregit` utilizes [`embed`](https://pkg.go.dev/embed) which does not support irregular files.

3. `clean` removes unused content like the `.github` folder.

`source/repo` is also a go package. This package embeds the flattened and cleaned contents of `github.com/github/gitignore` as a filesystem - which is then walked by the `ignoregit` CLI to list available templates and copy them into new files. 

## Full Help-Text
```shell
ignoregit -h
Usage:
  ignoregit [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list
  new

Flags:
  -h, --help   help for ignoregit

Use "ignoregit [command] --help" for more information about a command.

ignoregit list -h
Usage:
  ignoregit list [flags]

Flags:
  -a, --all    include community and global .gitignores
  -h, --help   help for list

ignoregit new -h
Usage:
  ignoregit new [language] [flags]

Flags:
  -h, --help         help for new
  -o, --out string   .gitignore output filename (default ".gitignore")
```

## TODO
- json output
- namespacing global/community templates
- search / aliases for templates