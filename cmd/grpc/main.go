package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/backendgolang/searchapp/util/config"
)

func main() {
	//Init Config
	err := config.InitConfig()
	if err != nil {
		msg := fmt.Sprintf("error when init config app: %v", err)
		log.Fatalln(msg)
	}

	//get port from config file
	port := fmt.Sprintf(":%s", config.Get().Port.GRPC)
	fmt.Println("[GRPC] starting http service on ", port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
