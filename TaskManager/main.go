package main

import (
	"github.com/codegangsta/negroni"
	"github.com/prince1809/go-web/TaskManager/common"
	"github.com/prince1809/go-web/TaskManager/routers"
	"log"
	"net/http"
)

//Entry point of the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	//Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
