package models

type CheckSystemRequestBody struct {
	TaskId     int    `json:"task_id"`
	CodeLang   string `json:"code_language"`
	SourceCode string `json:"source_code"`
}
