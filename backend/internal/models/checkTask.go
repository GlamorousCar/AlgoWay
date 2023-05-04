package models

type CheckTaskRequest struct {
	TaskID     uint64 `json:"task_id"`
	Lang       string `json:"lang"`
	SourceCode string `json:"source_code"`
}
