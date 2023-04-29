package models

type Algorithm struct {
	Id          int    `json:"algorithm_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Position    int    `json:"position"`
	ThemeId     int    `json:"theme_id"`
}
