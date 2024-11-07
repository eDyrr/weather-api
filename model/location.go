package model

type Location struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
	Time    string `json:"localtime"`
}
