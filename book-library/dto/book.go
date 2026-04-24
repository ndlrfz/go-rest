package dto

type BookReq struct {
	Title  string   `json:"title"`
	Genre  []string `json:"genre"`
	Author string   `json:"author"`
}
