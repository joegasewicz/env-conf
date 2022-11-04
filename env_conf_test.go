package env_conf

import (
	"os"
	"testing"
)

func TestUpdate(t *testing.T) {

	os.Setenv("ENV_ONE", "apples")
	os.Setenv("ENV_TWO", "bananas")

	type Config struct {
		EnvOne string `env_conf:"ENV_ONE"`
		EnvTwo string `env_conf:"ENV_TWO"`
	}

	c := Config{}
	err := Update(&c)

	if err != nil {
		t.Logf("expected error to be nil")
		t.Fail()
	}

	if c.EnvOne != "apples" {
		t.Logf("expected apples but got %s", c.EnvOne)
		t.Fail()
	}
	if c.EnvTwo != "bananas" {
		t.Logf("expected bananas but got %s", c.EnvOne)
		t.Fail()
	}

	err = Update(c)
	if err == nil {
		t.Logf("Expected error to not be nil")
		t.Fail()
	}
}
