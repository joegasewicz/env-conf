package env_conf

import (
	"errors"
	"os"
	"reflect"
)

// Update
func Update(c interface{}) error {
	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)
	// c must be a pointer
	if reflect.ValueOf(c).Kind() != reflect.Pointer {
		return errors.New("c must be a pointer")
	}
	e := v.Elem()
	te := t.Elem()
	// create map
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
