package main

import (
	"fmt"

	"github.com/aquilax/cooklang-go"
	"github.com/go-shiori/go-epub"
)

func createBook(file string, recipes []cooklang.RecipeV2) error {
	e, err := epub.NewEpub("Recipes")
	if err != nil {
		return err
	}

	e.SetAuthor("Patrick Kohan")

	for _, r := range recipes {
		s, err := executeTemplateRecipe(r)
		if err != nil {
			return err
		}

		_, err = e.AddSection(s, fmt.Sprint(r.Metadata["title"]), "", "")
		if err != nil {
			return err
		}
	}

	return e.Write(file)
}
