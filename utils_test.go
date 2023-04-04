package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHumanizeStr(t *testing.T) {
	assert.Equal(t, "Go To Func", HumanizeStr("goToFunc"))
	assert.Equal(t, "Go To Func", HumanizeStr("go_to_func"))
	assert.Equal(t, "Go To Func", HumanizeStr("go_to__func"))
	assert.Equal(t, "Go To Func", HumanizeStr("go_To_Func"))
	assert.Equal(t, "Go To Func", HumanizeStr("GoToFunc"))
	assert.Equal(t, "Go To Func", HumanizeStr("Go To Func"))
}

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, "goToFunc", ToCamelCase("goToFunc"))
	assert.Equal(t, "goToFunc", ToCamelCase("go to func"))
	assert.Equal(t, "goToFunc", ToCamelCase("go_to_func"))
	assert.Equal(t, "goToFunc", ToCamelCase("go to  func"))
	assert.Equal(t, "goToFunc", ToCamelCase("go To Func"))
	assert.Equal(t, "goToFunc", ToCamelCase("GoToFunc"))
	assert.Equal(t, "goToFunc", ToCamelCase("Go To Func"))
}

func TestToSnakeCase(t *testing.T) {
	assert.Equal(t, "go_to_func", ToSnakeCase("goToFunc"))
	assert.Equal(t, "go_to_func", ToSnakeCase("go to func"))
	assert.Equal(t, "go_to_func", ToSnakeCase("go_to_func"))
	assert.Equal(t, "go_to_func", ToSnakeCase("go to  func"))
	assert.Equal(t, "go_to_func", ToSnakeCase("go To Func"))
	assert.Equal(t, "go_to_func", ToSnakeCase("GoToFunc"))
	assert.Equal(t, "go_to_func", ToSnakeCase("Go To Func"))
}

func TestToPlural(t *testing.T) {
	assert.Equal(t, "tests", ToPlural("test"))
	assert.Equal(t, "flies", ToPlural("fly"))
	assert.Equal(t, "buses", ToPlural("bus"))
}

func TestParse(t *testing.T) {
	parsedInt32, ok := Parse[int32]("34")
	assert.Equal(t, int32(34), parsedInt32)
	assert.True(t, ok)

	parsedInt64, ok := Parse[int64]("34")
	assert.Equal(t, int64(34), parsedInt64)
	assert.True(t, ok)

	parsedBool, ok := Parse[bool]("true")
	assert.Equal(t, true, parsedBool)
	assert.True(t, ok)

	parsedStr, ok := Parse[string]("str")
	assert.Equal(t, "str", parsedStr)
	assert.True(t, ok)

	parsedFloat32, ok := Parse[float32]("1.35")
	assert.Equal(t, float32(1.35), parsedFloat32)
	assert.True(t, ok)

	parsedInvalid, ok := Parse[int32]("invalid_number")
	assert.Equal(t, int32(0), parsedInvalid)
	assert.False(t, ok)

	parsedUnsupported, ok := Parse[struct{name string}]("unsupported_type")
	assert.Equal(t, struct{name string}{name: ""}, parsedUnsupported)
	assert.False(t, ok)
}


func TestParseToType(t *testing.T) {
	parsedInt32, ok := ParseToType("34", reflect.TypeOf(int32(1)))
	assert.Equal(t, int32(34), parsedInt32)
	assert.True(t, ok)

	parsedInt64, ok := ParseToType("34", reflect.TypeOf(int64(1)))
	assert.Equal(t, int64(34), parsedInt64)
	assert.True(t, ok)

	parsedBool, ok := ParseToType("true", reflect.TypeOf(true))
	assert.Equal(t, true, parsedBool)
	assert.True(t, ok)

	parsedStr, ok := ParseToType("str", reflect.TypeOf(""))
	assert.Equal(t, "str", parsedStr)
	assert.True(t, ok)

	parsedFloat32, ok := ParseToType("2.35", reflect.TypeOf(float32(1)))
	assert.Equal(t, float32(2.35), parsedFloat32)
	assert.True(t, ok)

	parsedInvalid, ok := ParseToType("invalid_number", reflect.TypeOf(int32(1)))
	assert.Equal(t, int32(0), parsedInvalid)
	assert.False(t, ok)

	parsedUnsupported, ok := ParseToType("unsupported_type", reflect.TypeOf(struct{name string}{name: ""}))
	assert.Equal(t, nil, parsedUnsupported)
	assert.False(t, ok)
}