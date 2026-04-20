package handler

import (
	"book-library/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID - HANDLER", http.StatusBadRequest)
		return
	}

	book, err := h.service.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(book)
}
