# ENV Conf
Grabs environmental variables &amp; maps them to your config type via struct tags

```
go get -u github.com/joegasewicz/env-conf
```

### Usage
If an environment variable returns an empty string then the default config 
struct member will not be overridden.
```go
os.Setenv("ENV_ONE", "apples")
os.Setenv("ENV_TWO", "bananas")

type Config struct {
    EnvOne string `env_conf:"ENV_ONE"`
    EnvTwo string `env_conf:"ENV_TWO"`
}

c := Config{}
err := env_conf.Update(&c)

fmt.Println("ENV_ONE: ", c.EnvOne) // apples
```

#### Set Defaults
Using a `:` colon to set defaults, 
the environment variable is on the left side & the default is on the right side.
```go
type Config struct {
    EnvOne string `env_conf:"ENV_ONE:http://127.0.0.1:8080"`
}
// If `ENV_ONE` does not exist then `EnvOne` will equal http://127.0.0.1:8080
```