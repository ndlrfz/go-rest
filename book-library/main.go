package main

import (
	"book-library/config"
	"book-library/handler"
	"book-library/repository"
	"book-library/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	msg := "Health check Golang simple"
	json.NewEncoder(w).Encode(msg)
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

	mux := http.NewServeMux()

	mux.HandleFunc("GET /check", checkHealth)
	mux.HandleFunc("GET /book", hdl.GetBookHandler)

	fmt.Println("Web server running on 0.0.0.0:8000")
	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	server.ListenAndServe()
}
