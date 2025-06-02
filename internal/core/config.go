package core

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	PostgresUri    string `env:"POSTGRES_URI"`
	JwtPrivateKey  string `env:"JWT_PRIVATE_KEY"`
	JwtPublicKey   string `env:"JWT_PUBLIC_KEY"`
	JobListBaseUrl string `env:"JOB_LIST_BASE_URL" env-default:"https://dev6.dansmultipro.com"`
}

func LoadConfig(logger *zap.Logger) (*Config, error) {
	var config Config
	if err := godotenv.Load(); err != nil {
		logger.Warn("no .env found, using default envvar...", zap.Error(err))
	}

	if err := cleanenv.ReadEnv(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
