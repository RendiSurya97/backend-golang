package grpc

import (
	"fmt"
	"log"
	"net"

	pbSearchApp "github.com/backendgolang/searchapp/pkg/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/tokopedia/grace.v1"

	movieController "github.com/backendgolang/searchapp/internal/movie/controller"
)

var (
	movieCtrl movieController.Controller
)

type (
	//Handler will contains grpc servers setting
	Handler struct {
		s       *grpc.Server
		address string
	}

	//Option will contains all config to be used for grpc packages
	Option struct {
		Address string
	}

	server struct {
	}
)

// Init is function for initialization
func Init() {
	if movieCtrl == nil {
		movieCtrl = movieController.New()
	}
}

// New must be called to init grpc package
func New(o Option) *Handler {
	grpcServer := Handler{}
	s := grpc.NewServer()
	pbSearchApp.RegisterSearchAppServer(s, &server{})
	reflection.Register(s)

	grpcServer.s = s
	grpcServer.address = o.Address

	return &grpcServer
}

//Start shall be called after calling new, this will be used to execute grpc settings and get the handler ready to serve
func (h *Handler) Start() (err error) {
	var listener net.Listener

	//init used package
	Init()

	listener, err = grace.Listen(h.address)
	if err != nil {
		log.Printf("[GRPC] failed to listen grpc server,err: %v\n", err)
		return
	}

	fmt.Println("[GRPC] starting grpc service on ", h.address)

	err = h.s.Serve(listener)
	if err != nil {
		log.Printf("[GRPC] failed to serve grpc server,err: %v\n", err)
		return
	}
	return
}
