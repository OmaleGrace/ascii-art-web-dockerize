package handlers

import (
	"ascii-art-web-dockerize/ascii"
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
}

func ArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, "Wrong Method", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		ErrorHandler(w, "Bad Request", http.StatusBadRequest)
		return
	}

	file := "../banner/" + banner + ".txt"
	Lban, err := ascii.LoadBanner(file)
	if err != nil {
		ErrorHandler(w, "Error Loading Banner", http.StatusInternalServerError)
		return
	}

	genArt := ascii.GenerateArt(text, Lban)
	data := PageData{
		Result: genArt,
	}

	tmpl, err := template.ParseFiles(
		"../templates/base.html",
		"../templates/output.html",
	)

	if err != nil {
		ErrorHandler(w, "Error Parsing Template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
