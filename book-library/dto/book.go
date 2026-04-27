package dto

type BookReq struct {
	Title  string   `json:"title" validate:"required,omitempty"`
	Genre  []string `json:"genre" validate:"required,min=1,dive,min=3"`
	Author string   `json:"author" validate:"required"`
}
