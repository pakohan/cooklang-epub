package recipe

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/aquilax/cooklang-go"
	"github.com/goccy/go-yaml"
	"github.com/pakohan/cooklang-epub/book"
)

// Book is being passed to every temlate that is not marked as isRecipeSection.
type Book struct {
	Recipes  []*Recipe
	Tags     []Tag
	Metadata book.BookMetadata
}

func (b *Book) walkdirFunc(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if d.IsDir() {
		return nil
	}

	if filepath.Ext(path) != ".cook" {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	header, body := &bytes.Buffer{}, &bytes.Buffer{}
	inHeader := false
	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.Equal(line, []byte("---")) {
			inHeader = !inHeader
			continue
		}

		if inHeader {
			fmt.Fprintln(header, string(line))
		} else {
			fmt.Fprintln(body, string(line))
		}
	}

	parser := cooklang.NewParserV2(&cooklang.ParseV2Config{})
	recipe, err := parser.ParseStream(body)
	if err != nil {
		return err
	}

	m := RecipeMetadata{}
	err = yaml.NewDecoder(header).Decode(&m)
	if err != nil {
		return err
	}

	b.Recipes = append(b.Recipes, &Recipe{
		Steps:    recipe.Steps,
		Metadata: m,
	})
	return nil
}
