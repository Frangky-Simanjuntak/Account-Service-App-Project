package login

import (
	"be13/sql/account/entities"
	"database/sql"
	"log"
)

func LoginUser(db *sql.DB, users entities.Users) (entities.Users, error) {
	statm := db.QueryRow("SELECT id, nama, passwords, no_handphone ,saldo FROM users WHERE no_handphone = ? AND passwords = ?", users.No_Handphone, users.Password) //PREPARE MENYIAPKAN KODE YG AKAN DI EKSEKUSI DI SQL

	var row entities.Users
	errs := statm.Scan(&row.User_id, &row.Nama, &row.Password, &row.No_Handphone, &row.Saldo)
	if errs != nil {
		log.Fatal("Maaf No Telfon atau Password salah ", errs.Error())
	}
	return row, nil
}
