package config

import (
	"github.com/SiriusServiceDesk/auth-service/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type HttpServer struct {
	Address string `yaml:"address" env-default:"0.0.0.0:8000"`
}

type GRPCServer struct {
	Address string `yaml:"address" env-default:"0.0.0.0:3000"`
}

type AuthService struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env         string      `yaml:"env" env-required:"true"`
	HttpServer  HttpServer  `yaml:"http_server" env-required:"true"`
	GRPCServer  GRPCServer  `yaml:"grpc_server" env-required:"true"`
	AuthService AuthService `yaml:"auth_service"`
}

func GetConfig() *Config {

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist " + path)
	}

	var config Config

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		logger.Fatal("cannot read config file %s: %s", path, err)
	}

	return &config
}
