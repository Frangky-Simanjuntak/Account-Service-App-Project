package main

import (
	"be13/sql/account/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection := config.ConnectToDB()

	defer dbConnection.Close() // menutup koneksi
}
