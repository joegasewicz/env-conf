package env_conf

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

const TAG_NAME = "env_conf"

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
	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)
	// c must be a pointer
	if reflect.ValueOf(c).Kind() != reflect.Pointer {
		return errors.New("c must be a pointer")
	}
	e := v.Elem()
	te := t.Elem()
	count := e.NumField()

	for i := 0; i < count; i++ {
		f := e.Field(i)
		tag := te.Field(i).Tag.Get(TAG_NAME)
		tagSplit := strings.Split(tag, ":")
		if len(tagSplit) == 1 {
			env := os.Getenv(tag)
			if env != "" {
				f.SetString(env)
			}
		} else {
			// handle env var exists
			env := os.Getenv(tagSplit[0])
			if env != "" {
				f.SetString(env)
			} else {
				// handle multiple colons in default tag value
				var defaultVal string
				if len(tagSplit) > 2 {
					for j := 1; j < len(tagSplit); j++ {
						if j == 1 {
							defaultVal += tagSplit[j]
						} else {
							defaultVal += fmt.Sprintf(":%s", tagSplit[j])
						}
					}
					f.SetString(defaultVal)
				} else {
					f.SetString(tagSplit[1])
				}
			}
		}
	}
	return nil
}
