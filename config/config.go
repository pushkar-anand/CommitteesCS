package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// AppConfig stores all the config read from .env
type AppConfig struct {
	PORT        int
	Environment string
	Production  bool
	JWTSecret   string

	DBConfig *DBConfig
}

// DBConfig defines database config keys
type DBConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     int
}

var appConfig *AppConfig
var configOnce sync.Once

// GetAppConfig read from the .env file and returns instance of AppConfig
func GetAppConfig() *AppConfig {
	configOnce.Do(func() {
		initViper()
		appConfig = readConfig()
	})

	return appConfig
}

func initViper() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error initializing config.")
		panic(err)
	}
}

func readConfig() *AppConfig {
	env := viper.GetString("ENV")

	dbConfig := &DBConfig{
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetInt("DB_PORT"),
	}

	config := &AppConfig{
		PORT:        viper.GetInt("PORT"),
		Production:  strings.EqualFold(env, "PROD"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
		DBConfig:    dbConfig,
		Environment: env,
	}

	validateAppConfig(config)

	return config
}

func validateAppConfig(appConfig *AppConfig) {
	errStr := ""

	if appConfig.PORT == 0 {
		errStr += "server listening port not set" + "\n"
	}

	if appConfig.JWTSecret == "" {
		errStr += "JWT_SECRET not set" + "\n"
	}

	if errStr != "" {
		panic(errStr)
	}
}
