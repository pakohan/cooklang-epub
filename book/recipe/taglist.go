package recipe

import (
	"sort"
	"strings"
)

// Tag represents a category or label that can be associated with recipes.
// It contains a title for the tag and a list of associated recipes.
type Tag struct {
	// Title is the name of the tag, taken out of the tags array in the metadata section of the recipe.
	Title string
	// Recipes is a list of recipes associated with this tag.
	Recipes []RecipeTag
}

type RecipeTag struct {
	// Index is the index of the recipe in the book's recipe list.
	Index int
	// Recipe is the recipe with the index associated with this tag.
	Recipe *Recipe
}

func createTagList(book *Book) []Tag {
	tagMap := make(map[string][]RecipeTag)

	for i, recipe := range book.Recipes {
		for _, tag := range recipe.Metadata.Tags {
			tagMap[tag] = append(tagMap[tag], RecipeTag{Index: i, Recipe: recipe})
		}
	}

	var tags []Tag
	for title, recipes := range tagMap {
		tags = append(tags, Tag{Title: title, Recipes: recipes})
	}

	sort.Slice(tags, func(i, j int) bool {
		return strings.ToLower(tags[i].Title) < strings.ToLower(tags[j].Title)
	})

	return tags
}
