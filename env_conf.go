package env_conf

import (
	"errors"
	"os"
	"reflect"
)

// Update If an environment variable returns an empty string then the default config
// struct member will not be overridden.
//
//	os.Setenv("ENV_ONE", "apples")
//	os.Setenv("ENV_TWO", "bananas")
//
//	type Config struct {
//		EnvOne string `env_conf:"ENV_ONE"`
//		EnvTwo string `env_conf:"ENV_TWO"`
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

	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		tag := te.Field(i).Tag.Get("env_conf")
		env := os.Getenv(tag)
		if env != "" {
			f.SetString(env)
		}
	}
	return nil
}
