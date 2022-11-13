package configuration

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Constants struct {
	HOST string
	PORT string
	DB   struct {
		DIALECT    string
		URL        string
		NAME       string
		USER       string
		PASSWORD   string
		HOST       string
		PORT       string
		PROPERTIES struct {
			MaxOpenConnections int
			MaxIdleConnections int
			ConnMaxIdleTime    int
		}
	}
}

func initViper() (*Constants, error) {
	ConfigName := os.Getenv("CONFIG_FILE_NAME")
	ConfigPath := os.Getenv("CONFIG_FILE_PATH")
	if ConfigName == "" {
		ConfigName = "promotions.config"
	}
	if ConfigPath == "" {
		ConfigPath = "./configuration"
	}

	viper.SetConfigName(ConfigName)
	viper.AddConfigPath(ConfigPath)
	err := viper.ReadInConfig()
	if err != nil {
		return &Constants{}, err
	}
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", "1321")
	viper.SetDefault("DB.DIALECT", "postgres")

	var constants Constants
	err = viper.Unmarshal(&constants)
	if err != nil {
		log.Panic("Error initializing configuration: ", err)
	}
	return &constants, err
}

func New() (*Constants, error) {
	constants, err := initViper()
	if err != nil {
		log.Panic("Error initializing configuration: ", err)
	}
	return constants, err
}
