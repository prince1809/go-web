package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	fs := http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
	mux.Handle("/public/", fs)
	http.ListenAndServe(":9090", mux)
}
