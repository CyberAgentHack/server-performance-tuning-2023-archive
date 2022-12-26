package config

import (
	"fmt"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/db/config"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
)

type Config struct {
	DBConfig      *config.DBConfig
	RedisEndpoint string
	TraceConfig   *TraceConfig
}

func NewConfig(env string, dbSecretName string, redisEndpoint string) (*Config, error) {
	var cfg *Config
	switch env {
	case "prd":
		cfg = &Config{
			DBConfig: &config.DBConfig{
				SecretsManagerDBConfig: &config.SecretsManagerDBConfig{
					SecretID: dbSecretName,
				},
			},
			RedisEndpoint: redisEndpoint,
			TraceConfig: &TraceConfig{
				EnableTracing: true,
			},
		}
	case "cloud9":
		cfg = &Config{
			DBConfig: &config.DBConfig{
				SecretsManagerDBConfig: &config.SecretsManagerDBConfig{
					SecretID: dbSecretName,
				},
			},
			RedisEndpoint: redisEndpoint,
			TraceConfig: &TraceConfig{
				EnableTracing: false,
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
			RedisEndpoint: "localhost:6379",
			TraceConfig: &TraceConfig{
				EnableTracing: false,
			},
		}
	default:
		return nil, errcode.New(fmt.Errorf("unknown Environment: %s", env))
	}

	return cfg, nil
}

type TraceConfig struct {
	EnableTracing bool `json:"enable_tracing"`
}
