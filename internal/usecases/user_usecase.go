package usecases

import (
	"github.com/aguradev/react-go-auth/internal/repository"
	"github.com/sirupsen/logrus"
)

type UserUsecaseInterface interface {
	Login()
	Create()
}

type UserUsecase struct {
	UserRepo repository.UserRepoInterface
	Logrus   *logrus.Logger
}

func NewUserUsecase(userRepo repository.UserRepoInterface, logrus *logrus.Logger) UserUsecaseInterface {
	return &UserUsecase{
		UserRepo: userRepo,
		Logrus:   logrus,
	}
}

func (c *UserUsecase) Login() {
	c.Logrus.Infof("this is logger login")
}

func (c *UserUsecase) Create() {
	c.Logrus.Infof("this is logger create")
}
