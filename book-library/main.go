package main

import (
	"book-library/config"
	"book-library/handler"
	"book-library/repository"
	"book-library/service"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Health struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Health{
		Status: "ok",
		Msg:    "Running..",
	})
}

func main() {
	fmt.Println("Hello Golang")
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	repo := repository.NewBookRepository(db)
	svc := service.NewBookService(repo)
	hdl := handler.NewBookHandler(svc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	r.Get("/check", checkHealth)
	r.Get("/book/{id}", hdl.GetBookHandler)

	fmt.Println("Web server running on 0.0.0.0:8000")
	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	server.ListenAndServe()
}
