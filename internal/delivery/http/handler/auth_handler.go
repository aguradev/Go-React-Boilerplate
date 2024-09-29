package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aguradev/react-go-auth/internal/usecases"
	"github.com/sirupsen/logrus"
)

type AuthHandlerInterface interface {
	Authenticable(w http.ResponseWriter, r *http.Request)
	Registration()
}

type AuthHandler struct {
	UserUsecase usecases.UserUsecaseInterface
	Logrus      *logrus.Logger
}

func NewAuthHandler(userUsercase usecases.UserUsecaseInterface, logrus *logrus.Logger) AuthHandlerInterface {
	return &AuthHandler{
		UserUsecase: userUsercase,
		Logrus:      logrus,
	}
}

func (h *AuthHandler) Authenticable(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(map[string]interface{}{
		"message":     "Hello world",
		"status_code": 200,
	})

	if err != nil {
		h.Logrus.Fatal(err)
		return
	}

	h.Logrus.Info(http.StatusOK)

	w.Write([]byte(res))
}

func (h *AuthHandler) Registration() {

}
