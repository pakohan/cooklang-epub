package main

func main() {
	recipes, err := ParseFolder("/Users/mogli/Developer/recipes")
	if err != nil {
		panic(err)
	}

	err = createBook("/Users/mogli/Desktop/out.epub", recipes)
	if err != nil {
		panic(err)
	}
}
