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

	// assign routes
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	//get port from config file
	port := fmt.Sprintf(":%s", config.Get().Port.HTTP)
	fmt.Println("[HTTP] starting http service on ", port)
	http.ListenAndServe(port, nil)
}
