package models

type Personalidade struct {
	ID       int    `json:""`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade
