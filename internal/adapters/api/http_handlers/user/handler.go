package user

import (
	"jwt_registration_api/internal/adapters/api/http_handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	registerURL = "/register"
	loginURL    = "/login"
)

type handler struct {
	service Service
}

func NewHandler(service Service) http_handlers.Handler {
	return &handler{service: service}
}

func (h *handler) RegisterRoute(router *httprouter.Router) {
	router.POST(registerURL, h.RegisterUser)
	router.POST(loginURL, h.LoginUser)
}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (h *handler) LoginUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
