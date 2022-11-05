package env_conf

import (
	"os"
	"testing"
)

func TestUpdate(t *testing.T) {

	os.Setenv("ENV_ONE", "apples")
	os.Setenv("ENV_TWO", "bananas")
	os.Setenv("ENV_FOUR", "https://.google.com")
	os.Setenv("ENV_FIVE", "https://github.com:my::value")

	type Config struct {
		EnvOne   string `env_conf:"ENV_ONE"`
		EnvTwo   string `env_conf:"ENV_TWO:plums"`
		EnvThree string `env_conf:"ENV_THREE:http://127.0.0.1:8080"`
		EnvFour  string `env_conf:"ENV_FOUR:http://localhost:9090"`
		EnvFive  string `env_conf:"ENV_FIVE:http://github.com:my::value"`
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
		t.Logf("expected plums but got %s", c.EnvTwo)
		t.Fail()
	}

	// Check defaults
	if c.EnvThree != "http://127.0.0.1:8080" {
		t.Logf("expected http://127.0.0.1:8080 but got %s", c.EnvThree)
		t.Fail()
	}

	// Fixes bug - variable name is incorrect #5 (https://github.com/joegasewicz/env-conf/issues/5)
	if c.EnvFour != "https://.google.com" {
		t.Logf("expected https://.google.com but got %s", c.EnvFour)
		t.Fail()
	}

	if c.EnvFive != "https://github.com:my::value" {
		t.Logf("expected https://github.com:my::value but got %v", c.EnvFive)
		t.Fail()
	}

	// Check Errors
	err = Update(c)
	if err == nil {
		t.Logf("Expected error to not be nil")
		t.Fail()
	}
}
