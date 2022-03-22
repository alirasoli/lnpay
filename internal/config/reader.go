package config

import (
	"errors"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (config *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if path != "" {
		viper.AddConfigPath(path)
	}
	viper.AddConfigPath("/etc/lnpay/")
	viper.AddConfigPath("$HOME/.local/share/lnpay/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		} else {
			return nil, errors.New("error in config file")
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, err
}
