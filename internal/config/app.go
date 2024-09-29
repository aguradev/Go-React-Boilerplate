package config

import (
	"fmt"
	"net/http"

	"github.com/aguradev/react-go-auth/internal/delivery/http/handler"
	"github.com/aguradev/react-go-auth/internal/delivery/http/routes"
	"github.com/aguradev/react-go-auth/internal/repository"
	"github.com/aguradev/react-go-auth/internal/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	Router   *chi.Mux
	DB       *gorm.DB
	Config   *viper.Viper
	Log      *logrus.Logger
	PortAddr string
}

func (config *BootstrapConfig) Bootstrap() error {
	UserRepo := repository.NewUserRepo(config.DB)
	UserUsecase := usecases.NewUserUsecase(UserRepo, config.Log)
	AuthHandler := handler.NewAuthHandler(UserUsecase, config.Log)

	authRoute := routes.AuthRouteConfig{
		Router:      config.Router,
		AuthHandler: AuthHandler,
	}

	authRoute.AuthRoute()

	fmt.Printf("Listening on server http://localhost:%s \n", config.PortAddr)
	listenPort := fmt.Sprintf(":%s", config.PortAddr)
	return http.ListenAndServe(listenPort, config.Router)
}
