package main

import (
	"be13/sql/account/config"
	"be13/sql/account/controllers/Updated"
	"be13/sql/account/controllers/anotherusers"
	"be13/sql/account/controllers/deletee"
	"be13/sql/account/controllers/login"
	"be13/sql/account/controllers/read"
	"be13/sql/account/controllers/register"
	"be13/sql/account/controllers/topup"
	"be13/sql/account/controllers/transfer"
	"be13/sql/account/entities"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	in1 := entities.Topup{}
	dbConnection := config.ConnectToDB()
	defer dbConnection.Close() // menutup koneksi
	isRun := true
	for isRun {
		fmt.Print("MENU UTAMA: \n 1. Add Account \n 2. Login \n 0. Keluar\n")
		fmt.Print("Masukkan Pilihan Anda : ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			{
				newUser := entities.Users{}
				fmt.Print("masukkan nama : ")
				fmt.Scanln(&newUser.Nama)
				fmt.Print("masukkan password anda : ")
				fmt.Scanln(&newUser.Password)
				fmt.Print("masukkan no handphone anda : ")
				fmt.Scanln(&newUser.No_Handphone)
				register.Add(dbConnection, newUser)
			}
		case 2:
			{
				//ROW DIBUTUH KAN DI CASE READ, UPDATE AKUN
				in := entities.Users{}
				fmt.Println(" ")
				fmt.Print("Silahkan Masukkan Nomor Telepon Anda : ")
				fmt.Scanln(&in.No_Handphone)
				fmt.Print("Silahkan Masukkan Password Anda : ")
				fmt.Scanln(&in.Password)

				row, err := login.LoginUser(dbConnection, in)
				if err != nil {
					log.Fatal("\nERROR ON LOGIN DATA", err.Error())
				} else {
					fmt.Printf("\nSelamat Datang %s \nSaldo yang Anda Miliki adalah %d\n", row.Nama, row.Saldo)

					fmt.Println(" ")

					isRun2 := true
					for isRun2 {
						fmt.Print("\nSUB MENU:\n 1. Read Account \n 2. Update Account \n 3. Delete Account \n 4. Top-Up \n 5. Transfer \n 6. History Top-Up \n 7. History Transfer \n 8. Melihat Profil user lain \n 9. Keluar\n")
						fmt.Print("Masukkan Pilihan Anda : ")
						var choice2 int
						fmt.Scanln(&choice2)
						switch choice2 {
						case 1:
							{

								fmt.Println("=====================")
								fmt.Println("FITUR MELIHAT AKUN")
								read.ReadAkun(dbConnection, row)
								fmt.Println("=====================")

							}
						case 2:
							{
								fmt.Println("HALO SUB MENU 2")
								Updated.UpdatedAkun(dbConnection, row)

							}
						case 3:
							{
								fmt.Println("=====================")
								fmt.Println("FITUR DELETE AKUN")
								var No_Handphonee string
								fmt.Println("Harap Masukkan no handphone anda:")
								fmt.Scanln(&No_Handphonee)
								deletee.DeleteAkun(dbConnection, No_Handphonee)
								fmt.Println("=====================")

							}
						case 4:
							{
								topup.InsertToTopup(dbConnection, in1, row)
								topup.TopUpAkun(dbConnection)
							}
						case 5:
							{
								fmt.Println("==============")
								fmt.Println("FITUR TRANSFER")
								transfer.Transfer(dbConnection)
								fmt.Println("==============")
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
								cari := entities.Users{}
								fmt.Print("Masukkan No Telfon user lain : ")
								fmt.Scanln(&cari.No_Handphone)

								row, err := anotherusers.Search(dbConnection, cari)
								if err != nil {
									log.Fatal("\nERROR ON LOGIN DATA", err.Error())
								} else {
									fmt.Printf("\nNama   : %s \nSaldo : %d\n", row.Nama, row.Saldo)
								}
							}
						case 9:
							{
								fmt.Printf("\n Terimakasih Telah Bertransaksi  ^_^")
								isRun2 = false
							}
						}
					}
				}
			}
		case 0:
			{
				fmt.Println("THANK YOU :>>")
				isRun = false
			}
		}
	}
}
