package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/aquilax/cooklang-go"
	"github.com/goccy/go-yaml"
)

func ParseFolder(path string) ([]cooklang.RecipeV2, error) {
	var recipes []cooklang.RecipeV2
	parser := cooklang.NewParserV2(&cooklang.ParseV2Config{})

	return recipes, filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
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

		recipe, err := parser.ParseStream(body)
		if err != nil {
			return err
		}

		err = yaml.NewDecoder(header).Decode(&recipe.Metadata)
		if err != nil {
			return err
		}

		recipes = append(recipes, *recipe)
		return nil
	})
}
