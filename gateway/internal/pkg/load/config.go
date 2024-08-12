package load

import (
	"github.com/spf13/viper"
	"log"
)

type ServiceConfig struct {
	Host string
	Port int
}

type Config struct {
	ServerHost     string
	ServerPort     int
	CarWashService ServiceConfig
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		ServerHost: viper.GetString("server.host"),
		ServerPort: viper.GetInt("server.port"),
		CarWashService: ServiceConfig{
			Host: viper.GetString("services.carwash_service.host"),
			Port: viper.GetInt("services.carwash_service.port"),
		},
	}

	log.Printf("Configuration loaded: %+v", cfg)
	return cfg, nil
}
