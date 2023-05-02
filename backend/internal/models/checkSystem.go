package models

type Verdict struct {
	Abbr  string
	Title string
}
type CheckSystemIncomingData struct {
	TaskId     int    `json:"task_id"`
	CodeLang   string `json:"code_language"`
	SourceCode string `json:"source_code"`
}
