package Updated

import (
	"be13/sql/account/entities"
	"database/sql"
	"fmt"
	"log"
)

func UpdatedAkun(db *sql.DB, datadilogin entities.Users) {
	fmt.Println("=======================")
	fmt.Print("selamat datang di updated :\n 1. ganti nama\n 2. ganti password\n 3. ganti no handphone \n")
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
					fmt.Printf("UPDATE NAMA BERHASIIL :>\n")
				} else {
					fmt.Printf("UPDATE NAMA GAGAL :<\n")
				}
			}
		}
	case 2:
		{
			nama := entities.Users{}
			fmt.Print("masukan password baru : ")
			fmt.Scanln(&nama.Password)
			fmt.Print("masukan id anda : ")
			fmt.Scanln(&nama.User_id)
			
			var query = "UPDATE users SET password = ? WHERE id = ?"
			statement, errPrepare := db.Prepare(query)
			if errPrepare != nil {
				log.Fatal("error saat prepare", errPrepare.Error())
			}
			result, errExec := statement.Exec(nama.Password, datadilogin.User_id) //jumlah dan urutan disesuaikan dngan tanda tanya komen mysql
			if errExec != nil {
				log.Fatal("error di statement", errExec.Error())
			} else {
				row, errrow := result.RowsAffected()

				if row > 0 {
					fmt.Println("UPDATE PASSWORD BERHASIL :> ")
				} else {
					fmt.Println("UPDATE PASSWORD GAGAL :< ", errrow.Error())
				}
			}
		}
	case 3:
		{
			nama2 := entities.Users{}
			fmt.Println("Masukan no telfon yang baru")
			fmt.Scanln(&nama2.Nama)
	
			var query2 string = "UPDATE users SET no_handphone = ? WHERE id = ?"
			statement, errPrepare := db.Prepare(query2)

			if errPrepare != nil {
				log.Fatal("erorr prepare insert", errPrepare.Error())

			}
			result, errExec := statement.Exec(&nama2.No_Handphone, &datadilogin.User_id)
			if errExec != nil {
				log.Fatal("erorr Exec insert", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("UPDATE NO TELFON BERHASIIL :>")
				} else {
					fmt.Println("UPDATE NO TELFON GAGAL :<")
				}
			}

		}

	}
}
