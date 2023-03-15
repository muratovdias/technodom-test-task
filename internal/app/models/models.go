package models

type Link struct {
	ID          int    `json:"id"`
	ActiveLink  string `json:"active_link"`
	HistoryLink string `json:"history_link"`
}
