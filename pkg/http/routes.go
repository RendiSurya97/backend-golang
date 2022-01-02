package http

import (
	"github.com/julienschmidt/httprouter"
)

// AssignRoutes is for assigning route to the handler
func AssignRoutes(router *httprouter.Router) {
	router.GET("/ping", PingHandler)
	router.GET("/movies/:title", GetByTitleHandler)
}
