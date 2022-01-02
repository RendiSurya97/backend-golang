package controller

import (
	"log"
	"sync"
)

func (c *ctrl) GetByTitle(title string, page int64) (result MovieResult, err error) {

	var wg sync.WaitGroup
	wg.Add(2)

	//get data from omdb
	go func() {
		defer wg.Done()

		omdbRes, omdbErr := omdbExt.GetMovieByTitle(title, page)
		if omdbErr != nil {
			err = omdbErr
			return
		}

		result = MovieResult{
			Response:     omdbRes.Response,
			TotalResults: omdbRes.TotalResults,
		}

		for _, movieDetail := range omdbRes.MovieDetails {
			result.MovieDetails = append(result.MovieDetails, MovieDetail{
				Title:  movieDetail.Title,
				Year:   movieDetail.Year,
				ImdbID: movieDetail.ImdbID,
				Type:   movieDetail.Type,
				Poster: movieDetail.Poster,
			})
		}
	}()

	//save search title to db
	go func() {
		defer wg.Done()

		mErr := movieCore.SaveLog(title)
		if mErr != nil {
			log.Println("[Controller][SaveLog] failed save to DB: ", err.Error())
		}
	}()

	wg.Wait()

	return result, err
}
