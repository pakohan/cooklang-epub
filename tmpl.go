package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	"github.com/aquilax/cooklang-go"
)

var (
	//go:embed template/*
	templateFS embed.FS

	tmpl = template.Must(template.New("").Funcs(template.FuncMap{
		"getType": getType,
	}).ParseFS(templateFS, "template/*.tmpl"))
)

func executeTemplateRecipe(recipe cooklang.RecipeV2) (string, error) {
	buf := &bytes.Buffer{}
	err := tmpl.ExecuteTemplate(buf, "recipe.tmpl", recipe)
	return buf.String(), err
}

func getType(stepPart any) string {
	switch stepPart.(type) {
	case cooklang.IngredientV2:
		return "ingredient"
	case cooklang.TextV2:
		return "text"
	case cooklang.TimerV2:
		return "timer"
	case cooklang.CookwareV2:
		return "cookware"
	}

	panic(fmt.Sprintf("unknown type: %T", stepPart))
}
