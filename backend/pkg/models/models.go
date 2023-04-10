package models

import "errors"

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Theme struct {
	Id       int    `json:"theme_id"`
	Title    string `json:"title"`
	Position int    `json:"position"`
}

type Algorithm struct {
	Id          int    `json:"algorithm_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Position    int    `json:"position"`
	ThemeId     int    `json:"theme_id"`
}

type ThemeMenu struct {
	Id         int         `json:"theme_id"`
	Title      string      `json:"title"`
	Position   int         `json:"position"`
	Algorithms []Algorithm `json:"algorithms"`
}

type AlgorithmTheory struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type Task struct {
	Id          int    `json:"id"`
	AlgorithmId int    `json:"-"`
	IsSolved    bool   `json:"is_solved"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}
