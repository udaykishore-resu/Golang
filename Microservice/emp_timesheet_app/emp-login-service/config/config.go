package config

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Schema   string `mapstructure:"schema"`
}

type Config struct {
	Database      DatabaseConfig `mapstructure:"database"`
	Serverport    string         `mapstructure:"server.port"`
	JWTSecret     string         `mapstructure:"jwt.secret"`
	JWTExpiration string         `mapstructure:"jwt.expiration_hours"`
	AllowedUsers  []string       `mapstructure:"security.allowed_users"`
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file %s", err)
	}

	config := Config{
		Database: DatabaseConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			Username: viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
			Schema:   viper.GetString("database.schema"),
		},
		Serverport:    viper.GetString("server.port"),
		JWTSecret:     viper.GetString("jwt.secret"),
		JWTExpiration: viper.GetString("jwt.expiration_hours"),
		AllowedUsers:  viper.GetStringSlice("security.allowed_users"),
	}

	return config, nil
}
