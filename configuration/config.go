package configuration

import (
	"log"

	"github.com/spf13/viper"
)

type Constants struct {
	HOST string
	PORT string
	DB   struct {
		DIALECT string
		URL     string
		NAME    string
	}
}

func initViper() (*Constants, error) {
	viper.SetConfigName("promotions.config")
	viper.AddConfigPath("./configuration")
	err := viper.ReadInConfig()
	if err != nil {
		return &Constants{}, err
	}
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", "1321")
	viper.SetDefault("DB.DIALECT", "postgres")

	for _, k := range viper.AllKeys() {
		log.Println(k, ": ", viper.GetStringSlice(k))
	}

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
