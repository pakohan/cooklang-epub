package generate

import (
	"fmt"
	"html/template"

	"github.com/go-shiori/go-epub"
	"github.com/pakohan/cooklang-epub/book"
	"github.com/pakohan/cooklang-epub/book/recipe"
)

// Book generates an EPUB book using the provided template, book metadata, sections,
// and writes it to the specified file. It organizes the content into sections and
// subsections based on the provided structure.
//
// Parameters:
//
//	book - A pointer to a recipe.Book containing metadata and recipes for the book.
//	sections - A slice of book.Section defining the structure and content of the book.
//	tmpl - A parsed Go template used to render the content for sections and recipes.
//	file - The output file path where the generated EPUB will be written.
//
// Returns:
//
//	An error if any step in the EPUB generation process fails, or nil if successful.
func Book(book *recipe.Book, sections []book.Section, tmpl *template.Template, file string) error {
	e, err := epub.NewEpub("Rezepte")
	if err != nil {
		return err
	}
	e.SetTitle(book.Metadata.Title)
	e.SetDescription(book.Metadata.Description)
	e.SetAuthor(book.Metadata.Author)
	e.SetIdentifier(book.Metadata.Identifier)
	e.SetLang(book.Metadata.Language)

	for _, section := range sections {
		if section.IsRecipeSection {
			recipeSection, err := e.AddSection("", section.Title, "", "")
			if err != nil {
				return err
			}

			for i, r := range book.Recipes {
				s, err := getTemplateString(tmpl, section.TemplateName, r)
				if err != nil {
					return err
				}

				_, err = e.AddSubSection(recipeSection, s, fmt.Sprint(r.Metadata.Title), fmt.Sprintf("recipe_%d", i), "")
				if err != nil {
					return err
				}
			}
		} else {
			s, err := getTemplateString(tmpl, section.TemplateName, book)
			if err != nil {
				return err
			}

			_, err = e.AddSection(s, section.Title, section.ID, "")
			if err != nil {
				return err
			}
		}
	}

	return e.Write(file)
}
