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

	var err error

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		log.Printf("Попытка %d: не удалось подключиться", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		db.Close()
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
	defer db.Close()

	var result string
	err = db.QueryRow("SELECT 'PostgreSQL'").Scan(&result)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	fmt.Println(result)
}
