# go-utils

Basic utilities for go.

- [go-utils](#go-utils)
	- [Installation](#installation)
	- [Usage](#usage)
		- [Configuration](#configuration)
			- [Config from Environment Variables](#config-from-environment-variables)
		- [Errors](#errors)
		- [String Utilities](#string-utilities)

## Installation

Install using:

```
go get github.com/lgirma/go-utils
```

## Usage

### Configuration

Build a config from multiple sources as:

```go
import "utils"

config, err := utils.GetConfigFrom(utils.JsonStringConfigSource(`{
		"app": "test", 
		"age": 1, 
		"check": true,
		"price": 20.5,
		"sub_section": {
			"sub_key": "val"
		}
	}`)).
	Add(utils.JsonStringConfigSource(`{
		"age": 5,
		"check": false
	}`)).
	Add(utils.YamlStringConfigSource(`
		age: 10,
		sub_section:
		  sub_key: other_val
	`)).
	Add(utils.StaticMapConfigSource(mpa[string]any{
		"age": 15,
		"sub_section": map[string]any {
			"sub_key": "new_val",
		},
	}))
	Add(utils.FileConfigSource("config.local.json")).
	Add(utils.FileConfigSource("~/.config/my-app.yaml")).
	Build()
```

Then get config value as:

```go
config.GetStr("app") //returns "test"
config.GetInt64("age") //returns int64(15)
config.GetBool("check") //returns true
config.GetFloat64("price") //returns float64(20.5)

sub_config := config.GetSubSection("sub_section")
sub_config.GetStr("sub_key") //returns "new_val"
```

#### Config from Environment Variables

Configuration also supports environment variables. You will have to add a prefix to your environment variable to avoid collusion. Imagine the following environment variables (note the prefix `MYAPP`) :

* `MYAPP_USER` with value `sa`
* `MYAPP_DB_PASS` with value `pa$$w0rd`

```go
Add(utils.YamlStringConfigSource(`
		user: sa
		db_pass: pass123
	`)).
Add(utils.GetEnvVarConfigSource("MYAPP")).
Build()
```

will result in the config values:

```go
config.GetStr("user") //returns "sa"
config.GetStr("db_pass") //returns "pa$$w0rd"
```

### Errors

You can separate business errors from internal ones using:

```go
err := utils.NewBusinessError("AUTH_FAILED")

utils.IsBusinessError(err) //returns true
utils.IsBusinessError(errors.New("")) //returns false
```

### String Utilities

Use string utilities as:

```go
HumanizeStr("sub_nameStrCase") // returns "Sub name str case"
ToCamelCase("a_nameStrCase") // returns "subNameStrCase"
ToPascalCase("a_nameStrCase") // returns "SubNameStrCase"
ToSnakeCase("a_nameStrCase") // returns "sub_name_str_case"
ToPlural("bus") // returns "buses"
```