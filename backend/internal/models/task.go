package models

type Task struct {
	Id          int    `json:"id"`
	AlgorithmId int    `json:"-"`
	IsSolved    bool   `json:"is_solved"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}
