package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func middlewareFirst(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("MiddlewareFirst - Before Handler")
	next(w, r)
	log.Println("MiddlewareFirst - After Handler")
}

func middlewareSecond(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("MiddlewareSecond - Before Handler")
	if r.URL.Path == "/message" {
		if r.URL.Query().Get("password") == "pass123" {
			log.Println("Authorized to the system")
			next(w, r)
		} else {
			log.Println("Failed to authorize to the system")
			return
		}
	} else {
		next(w, r)
	}
	log.Println("MiddlewareSecond - After Handler")
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index handler")
	fmt.Fprintf(w, "Welcome!")
}

func message(w http.ResponseWriter, r *http.Request) {
	log.Print("Executing message handler")
	fmt.Fprintf(w, "HTTP middleware is awesome")
}

func myMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// logic before executing the next handler
	next(w, r)
	// logic after running the next handler
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/message", message)
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(middlewareFirst))
	n.Use(negroni.HandlerFunc(middlewareSecond))
	n.UseHandler(router)
	negroni.NewStatic(http.Dir("public"))
	n.Run(":8080")
}
