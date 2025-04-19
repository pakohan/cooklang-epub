package main

import (
	"flag"
	"fmt"

	"github.com/pakohan/cooklang-epub/book/generate"
	"github.com/pakohan/cooklang-epub/book/recipe"
	"github.com/pakohan/cooklang-epub/book/tmpl"
)

func main() {
	folder := flag.String("folder", ".", "Path to the folder containing recipes")
	outputFile := flag.String("output", "recipes.epub", "Path to the output EPUB file")
	templateFolder := flag.String("template-folder", "template", "folder that contains the template files")
	flag.Parse()

	fmt.Println("folder:", *folder)
	fmt.Println("output:", *outputFile)
	fmt.Println("template-folder:", *templateFolder)

	book, err := recipe.ParseRecipeFolder(*folder)
	if err != nil {
		panic(err)
	}

	tc, tmpl, err := tmpl.ParseTemplateFolder(*templateFolder)
	if err != nil {
		panic(err)
	}

	err = generate.Book(book, tc.Sections, tmpl, *outputFile)
	if err != nil {
		panic(err)
	}
}
