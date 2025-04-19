package book

// TemplateConfig represents the metadata for a template folder.
// It contains a list of sections that lays out the chapters of the book.
type TemplateConfig struct {
	Sections []Section `yaml:"sections"`
}

// Section represents a section in the generated book.
type Section struct {
	// TemplateName is the name of the template file to use for this section.
	TemplateName string `yaml:"templateName"`
	// Title is the title of the section.
	Title string `yaml:"title"`
	// IsRecipeSection indicates whether the section should be used for the recipes.
	// That means for each recipe there will be a subsection added.
	IsRecipeSection bool `yaml:"isRecipeSection"`
	// ID can be set for referencing it later, e.g. the table of contents.
	ID string `yaml:"id"`
}

// BookMetadata represents the metadata for the book.
type BookMetadata struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Author      string `yaml:"author"`
	Identifier  string `yaml:"identifier"`
	Language    string `yaml:"language"`
}
