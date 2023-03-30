package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromJson(t *testing.T) {
	service, err := GetConfigFrom(JsonStringConfigSource(`{"app": "test_2"}`)).
		Build()
	assert.Nil(t, err)
	assert.NotNil(t, service)

	service, err = GetConfigFrom(JsonStringConfigSource(`{app": "test_2"}`)).
		Build()
	assert.NotNil(t, err)
	assert.Nil(t, service)
}

func TestLoadConfigFromYaml(t *testing.T) {
	service, err := GetConfigFrom(YamlStringConfigSource(`app: "test"`)).
		Build()
	assert.Nil(t, err)
	assert.NotNil(t, service)

	service, err = GetConfigFrom(YamlStringConfigSource(`app "test"`)).
		Build()
	assert.NotNil(t, err)
	assert.Nil(t, service)
}

func TestLoadConfigFromEnvVar(t *testing.T) {
	os.Setenv("GO_UTILS_APP", "test")
	defer os.Unsetenv("GO_UTILS_APP")
	service, err := GetConfigFrom(StaticMapConfigSource(map[string]any{"app": "toasted"})).
		Add(GetEnvVarConfigSource("GO_UTILS")).
		Build()
	assert.Nil(t, err)
	assert.NotNil(t, service)
	assert.Equal(t, "test", service.GetStr("app"))
}

func TestDefaultConfigService(t *testing.T) {
	service, err := GetConfigFrom(JsonStringConfigSource(`{
		"app": "test", 
		"age": 1, 
		"check": true,
		"price": 20.5,
		"sub_section": {
			"sub_key": "val"
		}
	}`)).
		Build()

	assert.Nil(t, err)
	assert.NotNil(t, service)

	assert.Equal(t, "test", service.GetStr("app"))
	assert.Equal(t, int64(1), service.GetInt64("age"))
	assert.Equal(t, float64(20.5), service.GetFloat64("price"))
	assert.Equal(t, true, service.GetBool("check"))
	assert.Nil(t, service.GetAny("non_existing"))
}

func TestMultipleConfigSources(t *testing.T) {
	service, err := GetConfigFrom(JsonStringConfigSource(`{
		"app": "test", 
		"age": 1, 
		"check": true,
		"price": 20.5,
		"sub_section": {
			"sub_key": "val"
		}
	}`)).
	Add(JsonStringConfigSource(`{
		"age": 5,
		"check": false
	}`)).
	Build()

	assert.Nil(t, err)
	assert.NotNil(t, service)

	assert.Equal(t, "test", service.GetStr("app"))
	assert.Equal(t, int64(5), service.GetInt64("age"))
	assert.Equal(t, float64(20.5), service.GetFloat64("price"))
	assert.Equal(t, false, service.GetBool("check"))
	assert.Nil(t, service.GetAny("non_existing"))
}

func TestSubSection(t *testing.T) {
	service, err := GetConfigFrom(JsonStringConfigSource(`{
		"app": "test", 
		"age": 1, 
		"check": true,
		"price": 20.5,
		"sub_section": {
			"sub_key": "val"
		}
	}`)).
	Build()

	assert.Nil(t, err)
	assert.NotNil(t, service)

	section := service.SubSection("sub_section")
	assert.NotNil(t, section)
	assert.Equal(t, "val", section.GetStr("sub_key"))

	section_nil := service.SubSection("sub_section_non_existing")
	assert.Nil(t, section_nil)
}
