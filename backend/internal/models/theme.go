package models

type Theme struct {
	Id       int    `json:"theme_id"`
	Title    string `json:"title"`
	Position int    `json:"position"`
}
