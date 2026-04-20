package model

type Book struct {
	ID     uint     `json:"id"`
	Title  string   `json:"title"`
	Genre  []string `json:"genre"`
	Author string   `json:"author"`
}
