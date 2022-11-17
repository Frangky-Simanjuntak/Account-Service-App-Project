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
			// fmt.Println("masukan id anda")
			// fmt.Scanln(&nama.User_id)
			fmt.Println("di bawah ini hasil print query di updated")
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
					fmt.Println("berhasil")
				} else {
					fmt.Println("gagal")
				}
			}
			// query := "UPDATE users SET nama = ? WHERE id = ?"

			// var nama2 entities.Users
			// err := statm.Scan(&nama.Nama, &nama2.User_id)
			// if err != nil {
			// 	fmt.Println(err.Error())
			// 	log.Fatal("Maaf updated gagal")

			// }
			// statement, errPrepare := db.Prepare(query)
			// if errPrepare != nil {
			// 	log.Fatal("error saat prepare", errPrepare.Error())
			// }
			// result, errExec := statement.Exec(nama.Nama, datadilogin.User_id) //jumlah dan urutan disesuaikan dngan tanda tanya komen mysql
			// if errExec != nil {
			// 	log.Fatal("error di statement", errExec.Error())
			// } else {
			// 	row, errrow := result.RowsAffected()

			// 	if row > 0 {
			// 		fmt.Println("selamat updated berhasil")
			// 	} else {
			// 		fmt.Println("maaf updated gagal", errrow.Error())
			// 	}
			// }
			// fmt.Println(nama)
		}
	case 2:
		{
			nama := entities.Users{}
			fmt.Println("masukan password baru")
			fmt.Scanln(&nama.Password)
			fmt.Println("masukan id anda")
			fmt.Scanln(&nama.User_id)
			fmt.Println("di bawah ini hasil print query di updated")
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
					fmt.Println("selamat updated berhasil")
				} else {
					fmt.Println("maaf updated gagal", errrow.Error())
				}
			}
		}
	case 3:
		{

		}

	}
}
