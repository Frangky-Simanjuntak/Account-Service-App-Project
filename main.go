package main

import (
	"be13/sql/account/config"
	"be13/sql/account/controllers/Updated"
	"be13/sql/account/controllers/anotherusers"
	"be13/sql/account/controllers/deletee"
	"be13/sql/account/controllers/login"
	"be13/sql/account/controllers/register"
	"be13/sql/account/entities"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection := config.ConnectToDB()

	defer dbConnection.Close() // menutup koneksi
	in := entities.Users{}
	isRun := true
	for isRun {
		fmt.Println(" ")
		fmt.Print("MENU UTAMA: \n1. Add Account \n2. Login \n0. Keluar\n")
		fmt.Println("Masukkan Pilihan Anda  ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			{
				newUser := entities.Users{}
				fmt.Println("masukkan nama")
				fmt.Scanln(&newUser.Nama)
				fmt.Println("masukkan password")
				fmt.Scanln(&newUser.Password)
				fmt.Println("masukkan no handphone")
				fmt.Scanln(&newUser.No_Handphone)
				fmt.Println("ADD ACOUNT")
				register.Add(dbConnection, newUser)
			}
		case 2:
			{

				// in := entities.Users{}
				fmt.Println("Silahkan Masukkan Nomor Telepon Anda :")
				fmt.Scanln(&in.No_Handphone)
				fmt.Println("Silahkan Masukkan Password Anda :")
				fmt.Scanln(&in.Password)

				row, err := login.LoginUser(dbConnection, in)
				if err != nil {
					fmt.Println("ERROR ON LOGIN DATA", err.Error())
				} else {
					fmt.Println("Selamat Datang", row.Nama, "\nSaldo yang Anda Miliki adalah", row.Saldo)
				}

				fmt.Println(" ")
				fmt.Println("Hallo menu utama 2")
				isRun2 := true
				for isRun2 {
					fmt.Print("SUB MENU:\n1. Read Account \n2. Update Account \n3. Delete Account \n4. Top-Up \n5. Transfer \n6. History Top-Up \n7. History Transfer \n8. Melihat Profil user lain \n9. Keluar\n")
					fmt.Println("Masukkan Pilihan Anda ")
					var choice2 int
					fmt.Scanln(&choice2)
					switch choice2 {
					case 1:
						{

							fmt.Println("=====================")
							fmt.Println("FITUR MELIHAT AKUN")
							fmt.Printf(" id: %d\n nama: %s\n password: %d\n no telfon: %s\n saldo: %d\n", row.User_id, row.Nama, row.Password, row.No_Handphone, row.Saldo)
							fmt.Println("=====================")

						}
					case 2:
						{
							fmt.Println("HALO SUB MENU 2")
							Updated.UpdatedAkun(dbConnection, row)

						}
					case 3:
						{
							var No_Handphonee string
							fmt.Println("Masukkan no handphone anda:")
							fmt.Scanln(&No_Handphonee)
							deletee.DeleteAkun(dbConnection, No_Handphonee)

						}
					case 4:
						{
							fmt.Println("HALO SUB MENU 4")
						}
					case 5:
						{
							fmt.Println("HALO SUB MENU 5")
						}
					case 6:
						{
							fmt.Println("HALO SUB MENU 6")
						}
					case 7:
						{
							fmt.Println("HALO SUB MENU 7")
						}
					case 8:
						{
							fmt.Println("JINGOK USER LAIN")
							cari := entities.Users{}
							fmt.Println("Masukkan No telfon User yg ingin dilihat :")
							fmt.Scanln(&cari.No_Handphone)

							row, err := anotherusers.Search(dbConnection, cari)
							if err != nil {
								fmt.Println("Tidak di temukan", err)
							} else {
								fmt.Printf(" Nama : %s\n saldo : %d\n", row.Nama, row.Saldo)
							}

						}
					case 9:
						{
							fmt.Println("\n Terimakasih Telah Bertransaksi  ^_^ ")
							isRun2 = false
						}
					}
				}
			}
		case 0:
			{
				isRun = false
			}
		}
	}
}
