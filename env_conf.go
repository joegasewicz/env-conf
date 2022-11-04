package env_conf

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Update If an environment variable returns an empty string then the default config
// struct member will not be overridden.
//
//	os.Setenv("ENV_ONE", "apples")
//	os.Setenv("ENV_TWO", "bananas")
//
//	type Config struct {
//		EnvOne string `env_conf:"ENV_ONE"`
//		EnvTwo string `env_conf:"ENV_TWO:plums"` // Set a default value <ENV_VAR>:<DEFAULT_VALUE>
//	}
//
//	c := Config{}
//	err := Update(&c)
func Update(c interface{}) error {
	var env string
	tagName := "env_conf"
	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)
	// c must be a pointer
	if reflect.ValueOf(c).Kind() != reflect.Pointer {
		return errors.New("c must be a pointer")
	}
	e := v.Elem()
	te := t.Elem()

	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		tag := te.Field(i).Tag.Get(tagName)
		env = os.Getenv(tag)
		if env != "" {
			f.SetString(env)
		} else {
			tagSplit := strings.Split(tag, ":")
			if len(tagSplit) > 1 {
				// handle multiple colons in default tag value
				if len(tagSplit) > 2 {
					for i = 1; i < len(tagSplit); i++ {
						if i == 1 {
							env += tagSplit[i]
						} else {
							env += fmt.Sprintf(":%s", tagSplit[i])
						}
					}
					f.SetString(env)
				} else {
					f.SetString(tagSplit[1])
				}
			}
		}
	}
	return nil
}
