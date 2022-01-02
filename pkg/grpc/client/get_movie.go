package client

import (
	"context"

	pb "github.com/backendgolang/searchapp/pkg/grpc/proto"
)

func (c *Client) GetBundleByBundleIDs(ctx context.Context, title string, page int64) (*pb.MovieResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	convertedRequest := &pb.MovieRequest{
		Title: title,
		Page:  page,
	}

	// set timeout if no timeout exist
	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SearchAppClient.GetMovieByTitle(ctx, convertedRequest)
}
