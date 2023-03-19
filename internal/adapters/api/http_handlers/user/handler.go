package user

import (
	"encoding/json"
	"errors"
	"jwt_registration_api/internal/adapters/api/http_handlers"
	"jwt_registration_api/internal/adapters/api/http_handlers/dto"
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
	w.Header().Set("Content-Type", "application/json")

	regInput := dto.RegisterInput{}
	if err := json.NewDecoder(r.Body).Decode(&regInput); err != nil {
		err = errors.New("Incorrect input register user data: " + err.Error())
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	regPayload, err := h.service.Register(r.Context(), &regInput)
	if err != nil {
		err = errors.New("Can`t register user: " + err.Error())
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(regPayload); err != nil {
		h.logger.Error(err.Error())
	} else {
		h.logger.Info("Successful RegisterUser")
	}
}

func (h *handler) LoginUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loginInput := dto.LoginInput{}
	if err := json.NewDecoder(r.Body).Decode(&loginInput); err != nil {
		err = errors.New("Incorrect input login user data: " + err.Error())
		h.handleError(w, err, http.StatusBadRequest)
		return
	}

	loginPayload, err := h.service.Login(r.Context(), &loginInput)
	if err != nil {
		err = errors.New("Can`t login user: " + err.Error())
		h.handleError(w, err, http.StatusUnauthorized)
		return
	}

	if err = json.NewEncoder(w).Encode(loginPayload); err != nil {
		h.logger.Error(err.Error())
	} else {
		h.logger.Info("Successful login user")
	}
}

func (h *handler) handleError(w http.ResponseWriter, err error, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
	h.logger.Errorf("%s httpStatusCode: %d", err.Error(), httpStatusCode)
	if _, writeErr := w.Write([]byte(`{"error": "` + err.Error() + `"}`)); writeErr != nil {
		h.logger.Error(writeErr.Error())
	}
}
