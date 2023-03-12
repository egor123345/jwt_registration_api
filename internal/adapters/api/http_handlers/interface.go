package http_handlers

import (
	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	RegisterRoute(router *httprouter.Router)
}
