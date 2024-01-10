package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	ApplicationPort int `mapstructure:"APPLICATION_PORT"`

	MySQLHost     string `mapstructure:"MYSQL_HOST"`
	MySQLPort     int    `mapstructure:"MYSQL_PORT"`
	MySQLUser     string `mapstructure:"MYSQL_USER"`
	MySQLPassword string `mapstructure:"MYSQL_PASSWORD"`
	MySQLDatabase string `mapstructure:"MYSQL_DATABASE"`
}

// Load reads configuration from environment variables or file.
func Load(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)

	return config, err
}
