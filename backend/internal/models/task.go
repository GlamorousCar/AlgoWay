package models

type Task struct {
	Id          int    `json:"id"`
	AlgorithmId int    `json:"-"`
	IsSolved    string `json:"is_solved"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}
