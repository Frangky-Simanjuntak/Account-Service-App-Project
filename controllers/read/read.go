package read

import (
	"be13/sql/account/entities"
	"database/sql"
	"fmt"
	"log"
)

func ReadAkun(db *sql.DB, readid entities.Users) {
	statm := db.QueryRow("SELECT nama, passwords, no_handphone, saldo FROM users WHERE id = ?", readid.User_id)
	var in entities.Users
	errs := statm.Scan(&in.Nama, &in.Password, &in.No_Handphone, &in.Saldo)
	if errs != nil {
		log.Fatal("ERRORS line 15", errs.Error())
	}
	fmt.Printf("nama: %s\n password: %d\n no Handphone: %s\n saldo: %d\n ",in.Nama, in.Password, in.No_Handphone,in.Saldo)
}
