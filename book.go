package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"

	"github.com/aquilax/cooklang-go"
	"github.com/go-shiori/go-epub"
)

func createBook(tmpl *template.Template, recipes []cooklang.RecipeV2, file string) error {
	e, err := epub.NewEpub("Recipes")
	if err != nil {
		return err
	}

	e.SetAuthor("Patrick Kohan")

	if len(recipes) == 0 {
		return errors.New("no recipes found")
	}

	for _, r := range recipes {
		buf := &bytes.Buffer{}
		err := tmpl.ExecuteTemplate(buf, "recipe.tmpl", r)
		if err != nil {
			return err
		}

		_, err = e.AddSection(buf.String(), fmt.Sprint(r.Metadata["title"]), "", "")
		if err != nil {
			return err
		}
	}

	fmt.Printf("writing file to '%s'\n", file)
	return e.Write(file)
}
