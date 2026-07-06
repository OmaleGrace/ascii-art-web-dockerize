package handlers

import (
	"html/template"
	"net/http"
)

type ErrorData struct{
	Message string
	Status int
}

func ErrorHandler(w http.ResponseWriter, message string, code int) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/error.html",
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)

	Error := ErrorData{
		Status:  code,
		Message: message,
	}

	tmpl.Execute(w, Error)
}
