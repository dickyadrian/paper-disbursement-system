package config

import "github.com/kelseyhightower/envconfig"

type Redis struct {
	URL           string `required:"true"`
	QueueMaxRetry int    `envconfig:"QUEUE_MAX_RETRY" required:"false" default:"5"`
}

func LoadRedis() Redis {
	var config Redis
	envconfig.MustProcess("REDIS", &config)
	return config
}
