package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	pb "github.com/backendgolang/searchapp/pkg/grpc/proto"
	// "github.com/tokopedia/grpc/common/client"
)

var (
	defaultTimeout = 1 * time.Second
)

// SearchAppClientInterface is an interface for searchApp's grpc client
type SearchAppClientInterface interface {
	GetMovieByTitle(ctx context.Context, title string, page int64) (*pb.MovieResponse, error)
}

//Client shall act as grpc client for searchApp modules
type Client struct {
	SearchAppClient pb.SearchAppClient

	// conn is used for redial
	conn *grpc.ClientConn
}

//Options shall contains the configuration which will be used to get the client
type Options struct {
	Address string
}

// NewStandardClient will return client (with context if needed)
func NewStandardClient(ctx context.Context, o Options, opts ...grpc.DialOption) (*Client, error) {
	cli := new(Client)

	target := fmt.Sprintf("localhost:%s", o.Address)
	conn, err := grpc.Dial(target, opts)
	cli.conn = conn
	cli.SearchAppClient = pb.NewSearchAppClient(conn)

	return cli, err
}
