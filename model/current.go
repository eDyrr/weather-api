package model

type Current struct {
	LastUpdated   string    `json:"last_updated"`
	TempC         float32   `json:"temp_c"`
	TempF         float32   `json:"temp_f"`
	IsDay         bool      `json:"is_day"`
	Condition     Condition `json:"condition"`
	WindSpeed     float32   `json:"wind_kph"`
	WindDirection string    `json:"wind_dir"`
	Humidity      float32   `json:"humidity"`
	Cloud         float32   `json:"cloud"`
}
