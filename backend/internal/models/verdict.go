package models

type Verdict struct {
	ID    int    `json:"verdict_id"`
	Abbr  string `json:"abbr"`
	Title string `json:"title"`
}
