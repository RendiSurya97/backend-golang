package ext_module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/backendgolang/searchapp/util/config"
)

type (
	IOmdb interface {
		GetMovieByTitle(title string, page int64) (result MovieResult, err error)
	}
	omdb struct{}

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

func (e *extModule) NewOMDB() IOmdb {
	return &omdb{}
}

func (o *omdb) GetMovieByTitle(title string, page int64) (result MovieResult, err error) {
	url := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%d", config.Get().URL.OmdbURL, config.Get().SecretKey["omdb"].Value, title, page)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[Controller] error get movies detail: ", err)
		return
	}
	defer resp.Body.Close()

	movieResultReponse := MovieResult{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &movieResultReponse)

	return movieResultReponse, nil
}
