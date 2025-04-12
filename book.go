package main

import (
	"errors"
	"fmt"
	"html/template"

	"github.com/aquilax/cooklang-go"
	"github.com/go-shiori/go-epub"
)

func createBook(tmpl *template.Template, recipes []cooklang.RecipeV2, file string) error {
	if len(recipes) == 0 {
		return errors.New("no recipes found")
	}

	e, err := epub.NewEpub("Rezepte")
	if err != nil {
		return err
	}

	internalCSSPath, err := e.AddCSS("style.css", "")
	if err != nil {
		return err
	}

	s, err := getTemplateString(internalTemplates, "cover.tmpl", recipes)
	if err != nil {
		return err
	}

	_, err = e.AddSection(s, "Cover", "", internalCSSPath)
	if err != nil {
		return err
	}

	s, err = getTemplateString(internalTemplates, "toc.tmpl", recipes)
	if err != nil {
		return err
	}

	_, err = e.AddSection(s, "Inhaltsverzeichnis", "", "")
	if err != nil {
		return err
	}

	recipeSection, err := e.AddSection("", "Rezepte", "", "")
	if err != nil {
		return err
	}

	for i, r := range recipes {
		s, err = getTemplateString(tmpl, "recipe.tmpl", r)
		if err != nil {
			return err
		}

		_, err = e.AddSubSection(recipeSection, s, fmt.Sprint(r.Metadata["title"]), fmt.Sprintf("recipe_%d", i), "")
		if err != nil {
			return err
		}
	}

	s, err = getTemplateString(internalTemplates, "tags.tmpl", createTagList(recipes))
	if err != nil {
		return err
	}

	_, err = e.AddSection(s, "Tags", "tags.xhtml", "")
	if err != nil {
		return err
	}

	fmt.Printf("writing file to '%s'\n", file)
	return e.Write(file)
}
