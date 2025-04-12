package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	"github.com/aquilax/cooklang-go"
)

//go:embed templates/*.tmpl
var templates embed.FS

var internalTemplates = template.Must(template.ParseFS(templates, "templates/*.tmpl"))

func getTemplateString(t *template.Template, name string, payload any) (string, error) {
	buf := &bytes.Buffer{}
	err := t.ExecuteTemplate(buf, name, payload)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
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
