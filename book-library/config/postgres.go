package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Cannot load .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")

	connString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		host, port, dbname, user, password,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Cannot connect to the database server: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot ping the dtabase server: %v", err)
	}

	return db, nil
}

// package config
//
// import (
// 	"database/sql"
// 	"fmt"
// 	"os"
// 	"strconv"
//
// 	"github.com/joho/godotenv"
// 	_ "github.com/lib/pq"
// )
//
// func ConnectDB() (*sql.DB, error) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println("Cannot find .env file")
// 	}
//
// 	host := os.Getenv("DB_HOST")
// 	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
// 	user := os.Getenv("DB_USER")
// 	dbname := os.Getenv("DB_NAME")
// 	pass := os.Getenv("DB_PASS")
//
// 	connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
// 		host, port, user, dbname, pass,
// 	)
//
// 	db, err := sql.Open("postgres", connString)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if err := db.Ping(); err != nil {
// 		return nil, err
// 	}
//
// 	return db, nil
// }
