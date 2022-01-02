package main

import (
	"fmt"
	"log"
	"net/http"

	appHttp "github.com/backendgolang/searchapp/pkg/http"

	"github.com/backendgolang/searchapp/util/config"
	"github.com/julienschmidt/httprouter"
)

func main() {
	//Init Config
	err := config.InitConfig()
	if err != nil {
		msg := fmt.Sprintf("error when init config app: %v", err)
		log.Fatalln(msg)
	}

	//init http
	appHttp.Init()

	// routing
	router := httprouter.New()
	appHttp.AssignRoutes(router)

	//get port from config file
	port := fmt.Sprintf(":%s", config.Get().Port.HTTP)
	fmt.Println("[HTTP] starting http service on ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
