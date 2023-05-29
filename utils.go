package utils

import (
	"os"
	"reflect"
	"regexp"
	"strconv"
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

func Parse[T any](str string) (T, bool) {
	var val any = *new(T)
	if _, ok := val.(string); ok {
		val = str
		return val.(T), true
	} else if _, ok := val.(int32); ok {
		parsedInt64, err := strconv.ParseInt(str, 10, 32)
		val = int32(parsedInt64)
		return val.(T), err == nil
	} else if _, ok := val.(int64); ok {
		parsedInt64, err := strconv.ParseInt(str, 10, 64)
		val = parsedInt64
		return val.(T), err == nil
	} else if _, ok := val.(float32); ok {
		parsedFloat64, err := strconv.ParseFloat(str, 32)
		val = float32(parsedFloat64)
		return val.(T), err == nil
	} else if _, ok := val.(float64); ok {
		parsedFloat64, err := strconv.ParseFloat(str, 64)
		val = parsedFloat64
		return val.(T), err == nil
	} else if _, ok := val.(int16); ok {
		parsedInt64, err := strconv.ParseInt(str, 10, 16)
		val = int16(parsedInt64)
		return val.(T), err == nil
	} else if _, ok := val.(int); ok {
		parsedInt64, err := strconv.ParseInt(str, 10, 16)
		val = int(parsedInt64)
		return val.(T), err == nil
	} else if _, ok := val.(int8); ok {
		parsedInt64, err := strconv.ParseInt(str, 10, 8)
		val = int8(parsedInt64)
		return val.(T), err == nil
	} else if _, ok := val.(uint32); ok {
		parsedInt64, err := strconv.ParseUint(str, 10, 32)
		val = uint32(parsedInt64)
		return val.(T), err == nil
	} else if _, ok := val.(uint64); ok {
		parsedInt64, err := strconv.ParseUint(str, 10, 64)
		val = parsedInt64
		return val.(T), err == nil
	} else if _, ok := val.(uint16); ok {
		parsedInt64, err := strconv.ParseUint(str, 10, 16)
		val = uint16(parsedInt64)
		return val.(T), err == nil
	} else if _, ok := val.(uint8); ok {
		parsedInt64, err := strconv.ParseUint(str, 10, 8)
		val = uint8(parsedInt64)
		return val.(T), err == nil
	} else if _, ok := val.(bool); ok {
		parsedBool, err := strconv.ParseBool(str)
		val = parsedBool
		return val.(T), err == nil
	} else {
		return *new(T), false
	}
}

func TryParse[T any](str string) T {
	result, _ := Parse[T](str)
	return result
}

func ParseToType(str string, t reflect.Type) (any, bool) {
	var val any
	if t.Kind() == reflect.String {
		return str, true
	} else if t.Kind() == reflect.Int32 {
		parsedInt64, err := strconv.ParseInt(str, 10, 32)
		return int32(parsedInt64), err == nil
	} else if t.Kind() == reflect.Int64 {
		parsedInt64, err := strconv.ParseInt(str, 10, 64)
		return parsedInt64, err == nil
	} else if t.Kind() == reflect.Float32 {
		parsedInt32, err := strconv.ParseFloat(str, 32)
		return float32(parsedInt32), err == nil
	} else if t.Kind() == reflect.Float64 {
		parsedInt64, err := strconv.ParseFloat(str, 64)
		return parsedInt64, err == nil
	} else if t.Kind() == reflect.Int16 {
		parsedInt64, err := strconv.ParseInt(str, 10, 16)
		return int16(parsedInt64), err == nil
	} else if t.Kind() == reflect.Int8 {
		parsedInt64, err := strconv.ParseInt(str, 10, 8)
		return int8(parsedInt64), err == nil
	} else if t.Kind() == reflect.Uint32 {
		parsedInt64, err := strconv.ParseUint(str, 10, 32)
		return uint32(parsedInt64), err == nil
	} else if t.Kind() == reflect.Uint64 {
		parsedInt64, err := strconv.ParseUint(str, 10, 64)
		return parsedInt64, err == nil
	} else if t.Kind() == reflect.Uint16 {
		parsedInt64, err := strconv.ParseUint(str, 10, 16)
		return uint16(parsedInt64), err == nil
	} else if t.Kind() == reflect.Uint8 {
		parsedInt64, err := strconv.ParseUint(str, 10, 8)
		return uint8(parsedInt64), err == nil
	} else if t.Kind() == reflect.Bool {
		parsedBool, err := strconv.ParseBool(str)
		return parsedBool, err == nil
	} else {
		return val, false
	}
}
