package utils

import (
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
