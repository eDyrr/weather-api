package model

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}
