package main

import (
	"github.com/aguradev/react-go-auth/internal/config"
)

func main() {
	viperConfig := config.InitViper()
	db := config.InitDatabase(viperConfig)
	logrus := config.InitLogger(viperConfig)
	router := config.CreateNewServer()

	bootstrapConfig := config.BootstrapConfig{
		Router:   router,
		DB:       db,
		Config:   viperConfig,
		Log:      logrus,
		PortAddr: viperConfig.GetString("server.port"),
	}

	if err := bootstrapConfig.Bootstrap(); err != nil {
		logrus.Fatal(err)
	}
}
