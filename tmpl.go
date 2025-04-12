package main

import (
	"fmt"

	"github.com/aquilax/cooklang-go"
)

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
