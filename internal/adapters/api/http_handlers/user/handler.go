package user

import (
	"jwt_registration_api/internal/adapters/api/http_handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

const (
	registerURL = "/register"
	loginURL    = "/login"
)

type handler struct {
	service Service
	logger  *logrus.Logger
}

func NewHandler(service Service, logger *logrus.Logger) http_handlers.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) RegisterRoute(router *httprouter.Router) {
	router.POST(registerURL, h.RegisterUser)
	router.POST(loginURL, h.LoginUser)
}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("reg log")
	w.Write([]byte("Ну работает"))
}

func (h *handler) LoginUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("login log")
}
