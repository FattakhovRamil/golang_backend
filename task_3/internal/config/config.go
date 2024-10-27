// internal/config/config.go
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	InputFile  string `mapstructure:"input_file"`
	URL        string `mapstructure:"url"`
	LogFile    string `mapstructure:"log_file"`
	OutputFile string `mapstructure:"output_file"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
