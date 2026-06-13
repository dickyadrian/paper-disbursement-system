package config

import "github.com/kelseyhightower/envconfig"

type App struct {
	Name string `required:"true" default:"project-template"`
	Port int    `required:"true" default:"3030"`
}

func LoadApp() App {
	var config App
	envconfig.MustProcess("APP", &config)
	return config
}
