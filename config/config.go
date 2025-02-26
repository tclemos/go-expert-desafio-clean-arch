package config

import (
	"github.com/spf13/viper"
)

type GRPCConfig struct {
	Host string `mapstructure:"HOST"`
	Port int    `mapstructure:"PORT"`
}

type RESTConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     int    `mapstructure:"PORT"`
	GRPCHost string `mapstructure:"GRPC_HOST"`
	GRPCPort int    `mapstructure:"GRPC_PORT"`
}

type GraphQLConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     int    `mapstructure:"PORT"`
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
