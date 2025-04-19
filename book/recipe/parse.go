package recipe

import (
	"errors"
	"os"
	"path/filepath"
	"sort"

	"github.com/goccy/go-yaml"
)

// ParseRecipeFolder parses a folder containing recipe files and metadata,
// and returns a Book object representing the collection of recipes.
//
// The function performs the following steps:
//  1. Walks through the directory specified by the `path` parameter to collect
//     recipe files and populate the Book object.
//  2. Sorts the recipes alphabetically by their title.
//  3. Generates a list of tags based on the recipes in the folder.
//  4. Reads and decodes metadata from a "meta.yml" file in the folder.
//
// Parameters:
//   - path: The path to the folder containing recipe files and metadata.
//
// Returns:
//   - A pointer to a Book object containing the parsed recipes and metadata.
//   - An error if any issues occur during parsing, such as no recipes found
//     or failure to read metadata.
//
// Errors:
//   - Returns an error if the folder contains no recipes.
//   - Returns an error if the "meta.yml" file cannot be opened or decoded.
func ParseRecipeFolder(path string) (*Book, error) {
	var book Book
	err := filepath.WalkDir(path, book.walkdirFunc)
	if err != nil {
		return nil, err
	}

	if len(book.Recipes) == 0 {
		return nil, errors.New("no recipes found")
	}

	sort.Slice(book.Recipes, func(i, j int) bool {
		return book.Recipes[i].Metadata.Title < book.Recipes[j].Metadata.Title
	})

	book.Tags = createTagList(&book)

	f, err := os.Open(filepath.Join(path, "meta.yml"))
	if err != nil {
		return nil, err
	}

	err = yaml.NewDecoder(f).Decode(&book.Metadata)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
