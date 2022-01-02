package grpc

import (
	"context"
	"errors"

	"github.com/backendgolang/searchapp/internal/movie/controller"
	pb "github.com/backendgolang/searchapp/pkg/grpc/proto"
)

func (s *server) GetMovieByTitle(ctx context.Context, request *pb.MovieRequest) (*pb.MovieResponse, error) {
	movieResult, err := movieCtrl.GetByTitle(request.GetTitle(), request.GetPage())
	if err != nil {
		//write error message
		return &pb.MovieResponse{}, errors.New("error get movie")
	}

	return parseResponse(movieResult), nil
}

func parseResponse(res controller.MovieResult) *pb.MovieResponse {
	movieRes := &pb.MovieResponse{
		Response:     res.Response,
		TotalResults: res.TotalResults,
	}

	for _, movieDetail := range res.MovieDetails {
		movieRes.Search = append(movieRes.Search, &pb.MovieDetail{
			Title:  movieDetail.Title,
			Year:   movieDetail.Year,
			ImdbID: movieDetail.ImdbID,
			Type:   movieDetail.Type,
			Poster: movieDetail.Poster,
		})
	}

	return movieRes
}
