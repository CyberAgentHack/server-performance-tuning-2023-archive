package config

import "github.com/kelseyhightower/envconfig"

const prefix = "CA"

func Load(spec interface{}) error {
	return envconfig.Process(prefix, spec)
}
