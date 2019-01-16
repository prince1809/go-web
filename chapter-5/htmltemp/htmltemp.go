package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type Note struct {
	Title       string
	Description string
	createdOn   time.Time
}

// store for the Notes collection
var noteStore = make(map[string]Note)

// variable to generate key for the collection
var id int = 0

// Entry point of the program
func main() {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current dir: ", dir)

	r := mux.NewRouter()
	fs := http.FileServer(http.Dir(dir))

	r.Handle("/", fs)
	r.Handle("/public", http.FileServer(http.Dir("./public")))

	server := &http.Server{
		Addr:    ":9090",
		Handler: r,
	}

	log.Println("Listening...")
	server.ListenAndServe()
	//http.ListenAndServe(":9090", fs)

}
