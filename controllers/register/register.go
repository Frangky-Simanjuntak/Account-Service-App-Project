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
			fmt.Println("berhasil")
		} else {
			fmt.Println("gagal")
		}
	}
	return result, nil
}

// func Add(db *sql.DB, ) {
// newUser := entities.Users{}
// fmt.Println("masukkan nama")
// fmt.Scanln(&newUser.Nama)
// fmt.Println("masukkan password")
// fmt.Scanln(&newUser.Password)
// fmt.Println("masukkan no handphone")
// fmt.Scanln(&newUser.No_Handphone)

// var query = "INSERT INTO users(nama, passwords, no_handphone) VALUES (?,?,?)"
// statement, errPrepare := db.Prepare(query)
// if errPrepare != nil {
// 	log.Fatal("error prepare insert", errPrepare.Error())
// }

// result, errExec := statement.Exec(newUser.Nama, newUser.Password, newUser.No_Handphone)
// if errExec != nil {
// 	log.Fatal("error exec insert", errExec.Error())
// } else {
// 	row, _ := result.RowsAffected()
// 	if row > 0 {
// 		fmt.Println("Insert berhasil")
// 	} else {
// 		fmt.Println("Insert gagal")
// 	}
// 	}

// }
