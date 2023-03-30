package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	ENV_DEV        = "dev"
	ENV_STAGING    = "staging"
	ENV_PRODUCTION = "production"
)

type ConfigOptions struct {
	Env string
}

type ConfigSource interface {
	Load(prev *map[string]any) (*map[string]any, error)
}

type JsonConfigSource struct {
	_jsonSrc string
}

func (service *JsonConfigSource) Load(prev *map[string]any) (*map[string]any, error) {
	result := make(map[string]any)
	err := json.Unmarshal([]byte(service._jsonSrc), &result)
	if err != nil {
		return nil, err
	}
	return merge_maps(prev, &result), err
}

type StaticConfigSource struct {
	_srcMap map[string]any
}

func (service *StaticConfigSource) Load(prev *map[string]any) (*map[string]any, error) {
	return merge_maps(prev, &service._srcMap), nil
}

type YamlConfigSource struct {
	_yamlSrc string
}

func (service *YamlConfigSource) Load(prev *map[string]any) (*map[string]any, error) {
	result := make(map[string]any)
	err := yaml.Unmarshal([]byte(service._yamlSrc), &result)
	if err != nil {
		return nil, err
	}
	return merge_maps(prev, &result), err
}

type EnvVarConfigSource struct {
	_prefix string
}

func (service *EnvVarConfigSource) Load(prev *map[string]any) (*map[string]any, error) {
	if prev == nil {
		return nil, nil
	}
	for k := range *prev {
		envKey := strings.ToUpper(service._prefix + "_" + ToSnakeCase(k))
		val, ok := os.LookupEnv(envKey)
		if ok {
			(*prev)[k] = val
		}
	}
	return prev, nil
}

type ConfigSourceList struct {
	_list []ConfigSource
}

func GetConfigFrom(src ConfigSource) *ConfigSourceList {
	return &ConfigSourceList{
		_list: []ConfigSource{src},
	}
}

func (srcList *ConfigSourceList) Add(src ConfigSource) *ConfigSourceList {
	srcList._list = append(srcList._list, src)
	return srcList
}

func (srcList *ConfigSourceList) Build(options ...ConfigOptions) (ConfigService, error) {
	result := make(map[string]any)
	for i := range srcList._list {
		resultNew, err := srcList._list[i].Load(&result)
		if err != nil {
			return nil, err
		}
		result = *resultNew
	}
	return GetConfig(result, options...), nil
}

func FileConfigSource(fileName string) (ConfigSource, error) {
	if !file_exists(fileName) {
		return nil, nil
	}
	ext := strings.ToLower(path.Ext(fileName))
	if ext == ".json" {
		jsonContent, err := os.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
		return &JsonConfigSource{_jsonSrc: string(jsonContent)}, nil
	} else if ext == ".yaml" {
		yamlContent, err := os.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
		return &YamlConfigSource{_yamlSrc: string(yamlContent)}, nil
	} else {
		return nil, fmt.Errorf("config format '%s' not supported", ext)
	}
}

func JsonStringConfigSource(src string) ConfigSource {
	return &JsonConfigSource{_jsonSrc: src}
}

func YamlStringConfigSource(src string) ConfigSource {
	return &YamlConfigSource{_yamlSrc: src}
}

func GetEnvVarConfigSource(prefix string) ConfigSource {
	return &EnvVarConfigSource{_prefix: prefix}
}

func StaticMapConfigSource(src map[string]any) ConfigSource {
	return &StaticConfigSource{_srcMap: src}
}

type ConfigService interface {
	GetStr(string) string
	GetInt64(string) int64
	GetFloat64(string) float64
	GetBool(string) bool
	GetAny(string) any
	SubSection(string) ConfigService
}

// TODO:
// - Support multiple (n) sources: json, yaml, env_var, custom
// - Support merging from multiple sources
// - Support order (preceedence) for sources
// - Support default values (or not for performance?)
// - Avoid dependency baking
// - Test final bundle impact

func GetConfig(configMap map[string]any, options ...ConfigOptions) ConfigService {
	return &DefaultConfigService{
		_config: configMap,
	}
}

type DefaultConfigService struct {
	_config map[string]any
}

/* func (service *DefaultConfigService) AddSource(src ConfigSource) ConfigService {
	result, err := src.Load(&service._config)
	service._config = *result
	if err != nil {
		log.Printf("failed to load config source: %v", err)
	}
	return service
} */

func (service *DefaultConfigService) GetAny(key string) any {
	return service._config[key]
}

func (service *DefaultConfigService) GetBool(key string) bool {
	val := service._config[key]
	if val == nil {
		return false
	}
	return val.(bool)
}

func (service *DefaultConfigService) GetFloat64(key string) float64 {
	val := service._config[key]
	if val == nil {
		return 0
	}
	return val.(float64)
}

func (service *DefaultConfigService) GetInt64(key string) int64 {
	val := service._config[key]
	if val == nil {
		return 0
	}
	return int64(val.(float64))
}

func (service *DefaultConfigService) GetStr(key string) string {
	val := service._config[key]
	if val == nil {
		return ""
	}
	return val.(string)
}

func (service *DefaultConfigService) SubSection(key string) ConfigService {
	val := service._config[key]
	if val == nil {
		return nil
	}
	return GetConfig(val.(map[string]any))
}
