package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	db_user := "root"
	db_pass := "root"
	db_host := "localhost"
	db_port := 3303
	db_name := "emp_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db_user, db_pass, db_host, db_port, db_name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB Not Found")
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect DB")
	}
	fmt.Println("DB connected ......")
	return db
}
