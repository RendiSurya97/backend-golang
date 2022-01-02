package http

import (
	movieController "github.com/backendgolang/searchapp/internal/movie/controller"
)

var (
	movieCtrl movieController.Controller
)

// Init is function for initialization
func Init() {
	if movieCtrl == nil {
		movieCtrl = movieController.New()
	}
}
