package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tmpl *template.Template

func main() {

	tmpl, _ = template.ParseGlob("./templates/*.html")

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	router.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		log.Println("req sent to /weather")
	}).Methods("POST")

	http.ListenAndServe(":3000", router)
}
