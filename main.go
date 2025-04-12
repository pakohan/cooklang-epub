package main

import (
	"flag"
	"html/template"
)

func main() {
	folder := flag.String("folder", ".", "Path to the folder containing recipes")
	outputFile := flag.String("output", "recipes.epub", "Path to the output EPUB file")
	templateFolder := flag.String("template-files", "template", "glob pattern to find the template files")
	flag.Parse()

	tmpl, err := template.New("").Funcs(template.FuncMap{
		"getType": getType,
	}).ParseGlob(*templateFolder)
	if err != nil {
		panic(err)
	}

	for _, t := range tmpl.Templates() {
		println(t.Name())
	}

	recipes, err := ParseFolder(*folder)
	if err != nil {
		panic(err)
	}

	err = createBook(tmpl, recipes, *outputFile)
	if err != nil {
		panic(err)
	}
}
