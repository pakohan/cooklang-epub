# cooklang-epub

[![Go Reference](https://pkg.go.dev/badge/github.com/pakohan/cooklang-epub.svg)](https://pkg.go.dev/github.com/pakohan/cooklang-epub)

Command line tool and GitHub Action to generate ePUB books from cooklang recipes.

## Usage:

cooklang-epub --folder=<recipe folder> --template-folder=<template folder> --output=<output.epub>

[example template repository](https://github.com/pakohan/cooklang-epub-template)

[Example usage of the GitHub Action](https://github.com/pakohan/recipes/blob/main/.github/workflows/generate_epub.yml)

The recipe folder should contain at least one [cooklang](https://cooklang.org/) recipe file and the [meta.yml](https://pkg.go.dev/github.com/pakohan/cooklang-epub/book#BookMetadata) file.

The template folder should contain the Golang template files and the [meta.yml](https://pkg.go.dev/github.com/pakohan/cooklang-epub/book#TemplateConfig) file.
