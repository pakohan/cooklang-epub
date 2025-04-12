package main

import (
	"flag"
	"fmt"
	"html/template"
)

func main() {
	folder := flag.String("folder", ".", "Path to the folder containing recipes")
	outputFile := flag.String("output", "recipes.epub", "Path to the output EPUB file")
	templateFolder := flag.String("template-files", "template", "glob pattern to find the template files")
	flag.Parse()

	fmt.Println("folder:", *folder)
	fmt.Println("output:", *outputFile)
	fmt.Println("template-files:", *templateFolder)

	tmpl, err := template.New("").Funcs(template.FuncMap{
		"getType": getType,
	}).ParseGlob(*templateFolder)
	if err != nil {
		panic(err)
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
