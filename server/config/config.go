package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

var c *tomlConfig

type tomlConfig struct {
	DB database `toml:"database"`
}

type database struct {
	Addr string
	Name string
	User string
	Pass string
	TLS  bool
}

func Root() *tomlConfig {
	if c == nil {
		c = &tomlConfig{}
		if _, err := toml.DecodeFile("config.toml", c); err != nil {
			return &tomlConfig{}
		}
		if c.DB.Addr == "" {
			c.DB.Addr = os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
		}
	}
	return c
}
