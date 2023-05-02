package models

type Task struct {
	Id          int    `json:"id"`
	AlgorithmId int    `json:"-"`
	IsSolved    bool   `json:"is_solved"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}

type CheckTaskRequest struct {
	TaskID uint64 `json:"task_id"`
	Lang   string `json:"lang"`
	Code   string `json:"code"`
}
