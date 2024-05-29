package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment  string `mapstructure:"ENVIRONMENT"`
	DBSource     string `mapstructure:"DB_SOURCE"`
	MigrationURL string `mapstructure:"MIGRATION_URL"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
