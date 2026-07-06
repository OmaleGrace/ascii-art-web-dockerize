package main

import (
	"ascii-art-web-dockerize/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	serve := http.Dir("static/")
	get := http.StripPrefix("/static/", http.FileServer(serve))
	http.Handle("/static/", get)
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.ArtHandler)
	fmt.Println("Server Running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
