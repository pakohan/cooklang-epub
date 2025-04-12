package main

import (
	"bytes"
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

	return e.Write(file)
}
