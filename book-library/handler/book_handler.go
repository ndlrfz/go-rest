package handler

import (
	"book-library/dto"
	"book-library/service"
	"book-library/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBooksAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	books, err := h.service.GetBooks(ctx)
	if err != nil {
		render.Render(w, r, utils.ErrNotFound)
		return
	}

	render.JSON(w, r, books)

}

func (h *BookHandler) GetBookHandler(w http.ResponseWriter, r *http.Request) {

	// 1. Coba ambil dari Query Param (?id=1)
	// idParam := r.URL.Query().Get("id")
	idParam := chi.URLParam(r, "id")

	// 2. JIKA kosong, coba ambil dari Path (asumsi route /book/{id})
	if idParam == "" {
		// Mengambil sisa path setelah "/book/"
		idParam = strings.TrimPrefix(r.URL.Path, "/book/")
	}

	// Debugging (bisa dihapus nanti)
	// fmt.Println("ID yang ditangkap:", idParam)

	fmt.Println("Raw ID:", idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// http.Error(w, "Invalid ID: format harus angka", http.StatusBadRequest)
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	book, err := h.service.GetByID(r.Context(), uint(id))
	if err != nil {
		// http.Error(w, "Book Not Found", http.StatusNotFound)
		render.Render(w, r, utils.ErrNotFound)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(book)
	render.JSON(w, r, book)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var req dto.BookReq

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := h.service.CreateBook(r.Context(), &req); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{"message": "Book created successfully"})
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {

	var reqBook dto.BookReq
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idParam)
	fmt.Println(id)
	fmt.Println(&reqBook)

	if err := render.DecodeJSON(r.Body, &reqBook); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := h.service.UpdateBook(r.Context(), &reqBook, uint(id)); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	if err := h.service.DeleteBook(r.Context(), uint(id)); err != nil {
		render.Render(w, r, utils.ErrNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "Book deleted successfully"})

}
