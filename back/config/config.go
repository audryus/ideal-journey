package config

import (
	"ideal-journey/clients/logger"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		HTTP      `yaml:"http"`
		Log       `yaml:"logger"`
		Cassandra `yaml:"cassandra"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	Cassandra struct {
		Host     string `env-required:"true" yaml:"host" env:"CASSANDRA_HOST"`
		Keyspace string `env-required:"true" yaml:"keyspace" env:"CASSANDRA_KEYSPACE"`
	}
)

var config *Config

func GetConfig() *Config {
	if config != nil {
		return config
	}
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		logger.Error("[Config] Err", err)
		panic(err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		logger.Error("[Config ENV] Err", err)
		panic(err)
	}
	config = cfg
	return config
}
