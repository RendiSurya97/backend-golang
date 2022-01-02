package main

import (
	"fmt"
	"log"

	appGrpc "github.com/backendgolang/searchapp/pkg/grpc"
	"github.com/backendgolang/searchapp/util/config"
)

func main() {
	//Init Config
	err := config.InitConfig()
	if err != nil {
		msg := fmt.Sprintf("error when init config app: %v", err)
		log.Fatalln(msg)
	}

	grpcHandler := appGrpc.New(appGrpc.Option{
		Address: fmt.Sprintf(":%s", config.Get().Port.GRPC),
	})

	err = grpcHandler.Start()
	if err != nil {
		log.Fatalln("Fail to start grpc handler,err: ", err.Error())
	}
}
