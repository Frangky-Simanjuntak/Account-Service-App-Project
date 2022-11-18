package anotherusers

import (
	"be13/sql/account/entities"
	"database/sql"
	"log"
)

func Search(db *sql.DB, search entities.Users) (entities.Users, error) {
	result := db.QueryRow("SELECT nama, saldo FROM users WHERE no_handphone = ?", search.No_Handphone) // ditanya 1,  proses menjalankana query SQL

	var userrow entities.Users                              // penampung tiap baris data dari db                                                                                                     // membuat variabel penampung
	errScan := result.Scan(&userrow.Nama, &userrow.Saldo) //3 //melakukan scanning data dari masing" row dan menyimpannya kedalam variabel yang dibuat sebelumnya

	if errScan != nil {
		log.Fatal("ERROR ON SEARCH", errScan.Error())
	}
	return userrow, nil
}
