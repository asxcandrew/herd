package config

import (
	"github.com/kelseyhightower/envconfig"
)

type conf struct {
	DBHost           string `required:"true"`
	DBPort           string `default:"5432"`
	DBUser           string `required:"true"`
	DBPass           string `required:"true"`
	DBName           string `required:"true"`
	TLS              bool
	AuthSecretKey    string `split_words:"true" required:"true"`
	StorageAccessKey string `split_words:"true" required:"true"`
	StorageSecretKey string `split_words:"true" required:"true"`
	StorageEndpoint  string `split_words:"true" required:"true"`
	StoragePath      string `split_words:"true" required:"true"`
	StorageBucket    string `split_words:"true" required:"true"`
}

var C conf

func init() {
	C = conf{}
	err := envconfig.Process("api", &C)

	if err != nil {
		panic(err)
	}
}
