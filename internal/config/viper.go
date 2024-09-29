package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
	newViper := viper.New()

	newViper.SetConfigName("env")
	newViper.SetConfigType("json")
	newViper.AddConfigPath("./../")
	newViper.AddConfigPath("./")

	err := newViper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("error: failed to config viper: " + err.Error()))
	}

	return newViper
}
