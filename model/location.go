package model

type Location struct {
	Name    string `json:"name" redis:"name"`
	Region  string `json:"region" redis:"region"`
	Country string `json:"country" redis:"country"`
	Time    string `json:"localtime" redis:"time"`
}
