package model

type Current struct {
	LastUpdated   string    `json:"last_updated" redis:"lastupdated"`
	TempC         float32   `json:"temp_c" redis:"tempc"`
	TempF         float32   `json:"temp_f" redis:"tempf"`
	IsDay         bool      `json:"is_day" redis:"isday"`
	Condition     Condition `json:"condition" redis:"condition"`
	WindSpeed     float32   `json:"wind_kph" redis:"winspeed"`
	WindDirection string    `json:"wind_dir" redis:"winddirection"`
	Humidity      float32   `json:"humidity" redis:"humidity"`
	Cloud         float32   `json:"cloud" redis:"cloud"`
}
