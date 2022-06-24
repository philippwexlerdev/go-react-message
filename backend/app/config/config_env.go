package config

import "github.com/caarlos0/env/v6"

// AppConfig presents app conf
type AppConfig struct {
	Port string `env:"PORT" envDefault:"8005"`

	DBHost   string `env:"DB_HOST" envDefault:"localhost"`
	DBPort   string `env:"DB_PORT" envDefault:"5450"`
	DBUser   string `env:"DB_USER" envDefault:"postgres"`
	DBPass   string `env:"DB_PASS" envDefault:"postgres"`
	DBName   string `env:"DB_NAME" envDefault:"chaintraced"`
	DBSchema string `env:"DB_SCHEMA" envDefault:"public"`
}

var config AppConfig

func LoadConfig() {
	_ = env.Parse(&config)
}

func GetConfig() AppConfig {
	return config
}
