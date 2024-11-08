package model

type Weather struct {
	Location Location `json:"location" redis:"location"`
	Current  Current  `json:"current" redis:"current"`
}
