package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//PingHandler is for checking with ping
func PingHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "pong")
}
