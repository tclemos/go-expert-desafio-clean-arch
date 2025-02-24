package config

import (
	"github.com/spf13/viper"
)

type GRPCConfig struct {
	Host string `mapstructure:"GRPC_HOST"`
	Port int    `mapstructure:"GRPC_PORT"`
}

type RESTConfig struct {
	Host     string `mapstructure:"REST_HOST"`
	Port     int    `mapstructure:"REST_PORT"`
	GRPCHost string `mapstructure:"GRPC_HOST"`
	GRPCPort int    `mapstructure:"GRPC_PORT"`
}

type DBConfig struct {
	Path string `mapstructure:"DB_PATH"`
}

func LoadConfig[T any](path string, config *T) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(config); err != nil {
		return err
	}

	return nil
}
