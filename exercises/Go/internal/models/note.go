package models

type Note struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Contents string `json:"contents"`
}
