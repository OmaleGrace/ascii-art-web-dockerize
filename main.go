package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageData struct {
	Result string
}

type ErrorData struct{
	Message string
	Status int
}

func main() {
	serve:= http.Dir("static/")
	get := http.StripPrefix("/static/", http.FileServer(serve))
	http.Handle("/static/", get)
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", ArtHandler)
	fmt.Println("Server Running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
