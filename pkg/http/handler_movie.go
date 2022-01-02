package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//get movie by title
func GetByTitleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	title := p.ByName("title")
	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)

	movieResult, err := movieCtrl.GetByTitle(title, page)
	if err != nil {
		//write error message
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(movieResult)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
