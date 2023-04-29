package models

type ThemeMenu struct {
	Id         int         `json:"theme_id"`
	Title      string      `json:"title"`
	Position   int         `json:"position"`
	Algorithms []Algorithm `json:"algorithms"`
}
