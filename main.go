package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/eDyrr/weather-api/model"
	"github.com/gorilla/mux"
)

var tmpl *template.Template

func main() {
	tmpl, _ = template.ParseGlob("./templates/*.html")

	os.Setenv("API_KEY", "3598f70e46904e2f968144054240711")

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
		fmt.Print(r.FormValue("location"))
	})

	router.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("accessed /weather")
		location := r.FormValue("location")

		response, _ := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", os.Getenv("API_KEY"), location))

		defer response.Body.Close()

		data, _ := io.ReadAll(response.Body)

		var weather model.Weather

		json.Unmarshal(data, &weather)

		fmt.Printf("%+v\n", weather)
		tmpl.ExecuteTemplate(w, "result", &weather)
	}).Methods("POST")

	http.ListenAndServe(":3000", router)
}
