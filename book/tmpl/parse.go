package tmpl

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/aquilax/cooklang-go"
	"github.com/goccy/go-yaml"
	"github.com/pakohan/cooklang-epub/book"
)

// ParseTemplateFolder parses a folder containing template files and a metadata file.
// It expects the folder to contain template files with the ".tmpl" extension and a "meta.yml" file
// with metadata information.
//
// The function returns a TemplateConfig struct parsed from the "meta.yml" file, a compiled
// template.Template object containing all the parsed templates, or an error if any step fails.
//
// Parameters:
//   - path: The path to the folder containing the template files and metadata.
//
// Returns:
//   - *book.TemplateConfig: A pointer to the parsed TemplateConfig struct.
//   - *template.Template: A pointer to the compiled template.Template object.
//   - error: An error if any issue occurs during parsing.
func ParseTemplateFolder(path string) (*book.TemplateConfig, *template.Template, error) {
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"getType": getType,
	}).ParseGlob(filepath.Join(path, "*.tmpl"))
	if err != nil {
		return nil, nil, err
	}

	f, err := os.Open(filepath.Join(path, "meta.yml"))
	if err != nil {
		return nil, nil, err
	}

	tc := book.TemplateConfig{}
	err = yaml.NewDecoder(f).Decode(&tc)
	if err != nil {
		return nil, nil, err
	}

	return &tc, tmpl, nil
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
