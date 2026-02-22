package main

import "github.com/sidz111/go-gin-emp-crud/db"

func main() {
	db := db.ConnectDb()

	defer db.Close()

}
