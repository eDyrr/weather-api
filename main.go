package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/eDyrr/weather-api/model"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

var tmpl *template.Template

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	ctx := context.Background()

	tmpl, _ = template.ParseGlob("./templates/*.html")

	os.Setenv("API_KEY", "3598f70e46904e2f968144054240711")

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
		fmt.Print(r.FormValue("location"))
	})

	router.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {

		location := r.FormValue("location")

		var weather model.Weather

		err := client.HGetAll(ctx, fmt.Sprintf("location:%s", location)).Scan(&weather)

		if err != nil {
			panic(err)
		}

		if weather.Location.Name != "" {
			tmpl.ExecuteTemplate(w, "result", &weather)
			return
		}

		response, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", os.Getenv("API_KEY"), location))
		if err != nil {
			http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
			return
		}

		defer response.Body.Close()

		data, _ := io.ReadAll(response.Body)

		// var weather model.Weather

		if err := json.Unmarshal(data, &weather); err != nil {
			http.Error(w, "Failed to parse weather data", http.StatusInternalServerError)
			return
		}

		client.Set(ctx, location, weather, time.Second*20)

		fmt.Printf("%+v\n", weather)
		tmpl.ExecuteTemplate(w, "result", &weather)
	}).Methods("POST")

	http.ListenAndServe(":3000", router)
}
