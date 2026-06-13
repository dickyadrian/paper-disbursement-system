package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Database struct {
	URL string `required:"true"`
}

func LoadDatabase() Database {
	var config Database
	envconfig.MustProcess("DATABASE", &config)
	return config
}
