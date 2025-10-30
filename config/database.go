package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error

	connStr := "host=localhost user=postgres password=admin dbname=bioskop_db sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database tidak bisa diakses:", err)
	}

	fmt.Println("Berhasil konek ke database PostgreSQL!")
}
