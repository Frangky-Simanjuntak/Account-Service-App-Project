package topup

import (
	"be13/sql/account/entities"
	"database/sql"
	"fmt"
	"log"
)

var nohp = entities.Users{}
var nohp1 = entities.Topup{}

func InsertToTopup(db *sql.DB, datadilogin entities.Topup, in entities.Users) {
	fmt.Println(in.User_id)

	fmt.Println("masukan nominal topup")
	fmt.Scanln(&nohp1.Jumlah)
	var query = "INSERT INTO topup(id, users_id, jumlah, created_at, updated_at) VALUES (?,?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}
	result, errExec := statement.Exec(nohp1.Topup_id, in.User_id, nohp1.Jumlah, nohp1.Created_at, nohp1.Updated_at)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("berhasil")
		} else {
			fmt.Println("gagal")
		}
	}

}

func TopUpAkun(db *sql.DB) {
	fmt.Println("masukan no handphone anda")
	fmt.Scanln(&nohp.No_Handphone)
	fmt.Println(nohp.No_Handphone)

	statm := db.QueryRow("SELECT id, nama, passwords, saldo FROM users WHERE no_handphone = ?", nohp.No_Handphone)
	var in entities.Users
	errs := statm.Scan(&in.User_id, &in.Nama, &in.Password, &in.Saldo)
	if errs != nil {
		log.Fatal("ERRORS line 47", errs.Error())
	}
	fmt.Println("Ini id  users ", in)

	var query1 string = "UPDATE users SET saldo = ? WHERE id = ?"
	statement, errPrepare := db.Prepare(query1)

	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}
	saldosekarang := in.Saldo + nohp1.Jumlah
	result, errExec := statement.Exec(&saldosekarang, &in.User_id)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("selamat topup berhasil ")
		} else {
			fmt.Println("maaf topup gagal")
		}
	}
	fmt.Println("dibawah ini hasil print jumlah topup")
	fmt.Println(nohp1.Jumlah)
	fmt.Println("dibawah ini hasil print jumlah saldo awal")
	fmt.Println(in.Saldo)
}
