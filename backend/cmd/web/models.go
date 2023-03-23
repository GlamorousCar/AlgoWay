package main

type theme struct {
	Id       int    `json:"theme_id"`
	Title    string `json:"title"`
	Position int    `json:"position"`
}

type algorithm struct {
	Id          int    `json:"algorithm_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Position    int    `json:"position"`
	ThemeId     int    `json:"theme_id"`
}

type themeMenu struct {
	Id         int         `json:"theme_id"`
	Title      string      `json:"title"`
	Position   int         `json:"position"`
	Algorithms []algorithm `json:"algorithms"`
}

type algorithmTheory struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
