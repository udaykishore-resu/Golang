package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	ServerPort    string         `mapstructure:"server.port"`
	Database      DatabaseConfig `mapstructure:"database"`
	JWTSecret     string         `mapstructure:"jwt.secret"`
	JWTExpiration int            `mapstructure:"jwt.expiration_hours"`
	AllowedUsers  []string       `mapstructure:"security.allowed_users"`
}

// DatabaseConfig holds the database configuration
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

// LoadConfig loads configuration from config.yaml
func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	config := Config{
		ServerPort: viper.GetString("server.port"),
		Database: DatabaseConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			Name:     viper.GetString("database.name"),
		},
		JWTSecret:     viper.GetString("jwt.secret"),
		JWTExpiration: viper.GetInt("jwt.expiration_hours"),
		AllowedUsers:  viper.GetStringSlice("security.allowed_users"),
	}

	return config, nil
}
