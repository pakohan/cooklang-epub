package main

import (
	"sort"
	"strings"

	"github.com/aquilax/cooklang-go"
)

type Tag struct {
	Title   string
	Recipes []Recipe
}

type Recipe struct {
	Index int
	Title string
}

func createTagList(recipes []cooklang.RecipeV2) []Tag {
	tagMap := make(map[string][]Recipe)

	for i, recipe := range recipes {
		for _, tagI := range recipe.Metadata["tags"].([]interface{}) {
			tag := tagI.(string)
			tagMap[tag] = append(tagMap[tag], Recipe{Index: i, Title: recipe.Metadata["title"].(string)})
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
