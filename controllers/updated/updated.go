package Updated

import (
	"be13/sql/account/entities"
	"database/sql"
	"fmt"
	"log"
)

func UpdatedAkun(db *sql.DB, datadilogin entities.Users) {
	fmt.Println("di bawah ini hasil print datadilogin")
	fmt.Println(datadilogin)
	fmt.Println("=======================")
	fmt.Print("selamat datang di updated :\n1. ganti nama\n2. ganti password\n3. ganti no handphone \n")
	fmt.Println("masukan pilihan anda :")
	var pilihupdated int
	fmt.Scanln(&pilihupdated)
	switch pilihupdated {
	case 1:
		{
			nama := entities.Users{}
			fmt.Println("masukan nama baru")
			fmt.Scanln(&nama.Nama)
			var query1 string = "UPDATE users SET nama = ? WHERE id = ?"
			statement, errPrepare := db.Prepare(query1)

			if errPrepare != nil {
				log.Fatal("erorr prepare insert", errPrepare.Error())

			}
			result, errExec := statement.Exec(&nama.Nama, &datadilogin.User_id)
			if errExec != nil {
				log.Fatal("erorr Exec insert", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("selamat berhasil ganti nama ")
				} else {
					fmt.Println("maaf ganti nama gagal")
				}
			}

		}
	case 2:
		{
			untukpassword := entities.Users{}
			nama := entities.Users{}
			fmt.Println("masukan password lama")
			fmt.Scanln(&untukpassword.Password)
			fmt.Println("masukan password baru")
			fmt.Scanln(&nama.Password)
			var query1 string = "UPDATE users SET passwords = ? WHERE id = ?"
			statement, errPrepare := db.Prepare(query1)

			if errPrepare != nil {
				log.Fatal("erorr prepare insert", errPrepare.Error())

			}
			result, errExec := statement.Exec(&nama.Password, &datadilogin.User_id)
			if errExec != nil {
				log.Fatal("erorr Exec insert", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("selamat berhasil ganti password ")
				} else {
					fmt.Println("maaf ganti password gagal")
				}
			}

		}
	case 3:
		{
			nama := entities.Users{}
			fmt.Println("masukan no handphone baru")
			fmt.Scanln(&nama.No_Handphone)
			var query1 string = "UPDATE users SET no_handphone = ? WHERE id = ?"
			statement, errPrepare := db.Prepare(query1)

			if errPrepare != nil {
				log.Fatal("erorr prepare insert", errPrepare.Error())

			}
			result, errExec := statement.Exec(&nama.No_Handphone, &datadilogin.User_id)
			if errExec != nil {
				log.Fatal("erorr Exec insert", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("selamat berhasil ganti no handphone ")
				} else {
					fmt.Println("maaf ganti no handphone gagal")
				}
			}
		}

	}
}
