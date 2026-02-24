package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func handler(w http.ResponseWriter, r *http.Request) {
	var result string
	err := db.QueryRow("SELECT 'Hello from PostgreSQL!'").Scan(&result)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
	fmt.Fprintln(w, result)
}

func main() {
	connStr := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s sslmode=disable",
		mustEnv("POSTGRES_USER"),
		mustEnv("POSTGRES_PASSWORD"),
		mustEnv("POSTGRES_DB"),
	)

	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		log.Printf("Попытка %d: не удалось подключиться, жду...", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", handler)
	log.Println("Сервер запущен на :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func mustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Переменная окружения %s не задана", key)
	}
	return val
}
