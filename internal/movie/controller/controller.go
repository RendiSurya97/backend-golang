package controller

import (
	"log"

	mCore "github.com/backendgolang/searchapp/internal/movie/core"
	"github.com/backendgolang/searchapp/internal/movie/ext_module"
)

type (
	Controller interface {
		GetByTitle(title string, page int64) (result MovieResult, err error)
	}
	ctrl struct{}

	MovieResult struct {
		MovieDetails []MovieDetail `json:"Search"`
		TotalResults string        `json:"totalResults"`
		Response     string        `json:"Response"`
	}

	MovieDetail struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	}
)

var (
	movieCore mCore.Movie
	omdbExt   ext_module.IOmdb
)

func New() Controller {
	err := prepare()
	if err != nil {
		log.Println("[Controller][Movie] New Controller - failed to prepare: ", err.Error())
	}

	return &ctrl{}
}

func prepare() (err error) {
	if movieCore == nil {
		movieCore = mCore.New()
	}

	if omdbExt == nil {
		omdbExt = ext_module.NewExtModule().NewOMDB()
	}

	return nil
}
