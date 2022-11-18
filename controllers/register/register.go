package register

import (
	"be13/sql/account/entities"
	"database/sql"

	//"database/sql"
	"fmt"
	"log"
)

func Add(db *sql.DB, newUser entities.Users) (sql.Result, error) {

	var query = "INSERT INTO Users(nama, passwords, no_handphone, saldo) VALUES (?,?,?,?)"
	statement, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}
	result, errExec := statement.Exec(newUser.Nama, newUser.Password, newUser.No_Handphone, 0)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Printf("Hallo %s Register Akun Berhasil :>\n", newUser.Nama)
			fmt.Println(" ")
		} else {
			fmt.Printf("Register Akun Gagal o_0 \n")
		}
	}
	return result, nil
}