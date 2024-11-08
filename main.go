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
	})

	defer client.Close()

	ctx := context.Background()

	status, _ := client.Ping(ctx).Result()

	fmt.Println(status)

	tmpl, _ = template.ParseGlob("./templates/*.html")

	os.Setenv("API_KEY", "3598f70e46904e2f968144054240711")

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	router.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {

		var weather model.Weather

		// err := client.HGetAll(ctx, fmt.Sprintf("location:%s", location)).Scan(&weather)

		jsonWeather, err := client.Get(ctx, r.FormValue("location")).Result()

		if jsonWeather != "" {
			fmt.Print("from redis")
			fmt.Printf("%+v\n", jsonWeather)
		}

		if err == redis.Nil {
			fmt.Print("from API")
			response, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", os.Getenv("API_KEY"), r.FormValue("location")))
			if err != nil {
				http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
			}

			defer response.Body.Close()

			jsonWeather, _ := io.ReadAll(response.Body)

			if err := json.Unmarshal(jsonWeather, &weather); err != nil {
				// http.Error(w, "Failed to parse weather data", http.StatusInternalServerError)
			}
			fmt.Printf("%+v\n", weather)
			// tmpl.ExecuteTemplate(w, "result", &weather)

			client.Set(ctx, r.FormValue("location"), jsonWeather, time.Second*20)
		} else {
			json.Unmarshal([]byte(jsonWeather), &weather)
		}

		tmpl.ExecuteTemplate(w, "result", &weather)
	}).Methods("POST")

	http.ListenAndServe(":3000", router)
}
