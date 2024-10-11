package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config") // Name of the file (without extension)
	viper.SetConfigType("yaml")   // Type of the config file
	viper.AddConfigPath(".")      // Path to look for the config file (current directory)

	// Handle error if config file cannot be read
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	AppConfig = Config{
		DBHost:     viper.GetString("db.host"),
		DBPort:     viper.GetInt("db.port"),
		DBUser:     viper.GetString("db.user"),
		DBPassword: viper.GetString("db.password"),
		DBName:     viper.GetString("db.dbname"),
	}
}
