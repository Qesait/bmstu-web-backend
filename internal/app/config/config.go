package config

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceHost   string
	ServicePort   int
	MinioEndpoint string
	BucketName    string

	JWT JWTConfig
}

type JWTConfig struct {
	Token         string
	SigningMethod jwt.SigningMethod
	ExpiresIn     time.Duration
}

func NewConfig() (*Config, error) {
	var err error

	configName := "config"
	_ = godotenv.Load()
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.WatchConfig()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	if os.Getenv("JWT_TOKEN") != "" {
		cfg.JWT.Token = os.Getenv("JWT_TOKEN")
	} else {
		return nil, errors.New("JWT_TOKEN env variable not provided")
	}

	log.Info("config parsed")

	return cfg, nil
}
