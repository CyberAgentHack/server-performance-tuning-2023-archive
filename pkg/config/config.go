package config

import (
	"fmt"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/db/config"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
)

type Config struct {
	DBConfig *config.DBConfig
}

func NewConfig(env string, dbSecretName string) (*Config, error) {
	var cfg *Config
	switch env {
	case "prd":
		cfg = &Config{
			DBConfig: &config.DBConfig{
				SecretsManagerDBConfig: &config.SecretsManagerDBConfig{
					SecretID: dbSecretName,
				},
			},
		}
	case "local":
		cfg = &Config{
			DBConfig: &config.DBConfig{
				RawDBConfig: &config.RawDBConfig{
					Username: "root",
					Password: "",
					Host:     "localhost",
					Port:     3306,
					DB:       "wsperf",
				},
			},
		}
	default:
		return nil, errcode.New(fmt.Errorf("unknown Environment: %s", env))
	}

	return cfg, nil
}
