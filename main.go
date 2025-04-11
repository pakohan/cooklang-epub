package main

import "flag"

func main() {
	folder := flag.String("folder", ".", "Path to the folder containing recipes")
	outputFile := flag.String("output", "recipes.epub", "Path to the output EPUB file")
	flag.Parse()

	recipes, err := ParseFolder(*folder)
	if err != nil {
		panic(err)
	}

	err = createBook(*outputFile, recipes)
	if err != nil {
		panic(err)
	}
}
