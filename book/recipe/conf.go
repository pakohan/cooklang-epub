package recipe

import (
	"github.com/aquilax/cooklang-go"
)

// Recipe is passed to the template of the subsections of the section that is marked as isRecipeSection: true
type Recipe struct {
	Steps    []cooklang.StepV2
	Metadata RecipeMetadata
}

// RecipeMetadata is the cooklang recipe metadata.
type RecipeMetadata struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Cuisine     string   `yaml:"cuisine"`
	Locale      string   `yaml:"locale"`
	Source      Source   `yaml:"source"`
	Tags        []string `yaml:"tags"`
}

type Source struct {
	URL string `yaml:"url"`
}
