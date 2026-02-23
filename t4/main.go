package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=db user=user password=password dbname=db sslmode=disable"

	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
		}
		if err == nil {
			break
		}
		log.Printf("Попытка %d: не удалось подключиться", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
	defer db.Close()

	var result string
	err = db.QueryRow("SELECT 'PostgreSQL'").Scan(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
