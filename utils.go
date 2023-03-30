package utils

import (
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func merge_maps(map1 *map[string]any, map2 *map[string]any) *map[string]any {
	if map2 == nil {
		return map1
	} else if map1 == nil {
		return map2
	}
	result := make(map[string]any)
	for key := range *map1 {
		result[key] = (*map1)[key]
	}
	for key := range *map2 {
		result[key] = (*map2)[key]
	}
	return &result
}

var RegCamel, _ = regexp.Compile(`([A-Z])`)
var RegSnake, _ = regexp.Compile(`_+([a-zA-Z0-9])`)
var RegSpaces, _ = regexp.Compile(`\s+`)
var RegAfterSpace, _ = regexp.Compile(`\s+([a-zA-Z0-9])`)

func HumanizeStr(src string) string {
	result := src
	result = RegSnake.ReplaceAllString(result, " $1")
	result = RegCamel.ReplaceAllString(result, " $1")
	result = cases.Title(language.English).String(result)
	result = RegSpaces.ReplaceAllString(result, " ")
	return strings.TrimSpace(result)
}

func ToCamelCase(src string) string {
	result := HumanizeStr(src)
	result = RegAfterSpace.ReplaceAllString(result, "$1")
	result = strings.ToLower(result[0:1]) + result[1:]
	return strings.TrimSpace(result)
}

func ToPascalCase(src string) string {
	result := ToCamelCase(src)
	return strings.ToUpper(result[0:1]) + result[1:]
}

func ToSnakeCase(src string) string {
	result := HumanizeStr(src)
	result = RegSpaces.ReplaceAllString(result, "_")
	return strings.ToLower(result)
}

func ToPlural(src string) string {
	lastChar := src[len(src)-1:]
	if lastChar == "s" {
		return src + "es"
	} else if lastChar == "y" {
		return src[:len(src)-1] + "ies"
	}
	return src + "s"
}

func file_exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
