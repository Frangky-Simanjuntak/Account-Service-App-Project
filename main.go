package main

import (
	"be13/sql/account/config"
	"be13/sql/account/controllers/Updated"
	"be13/sql/account/controllers/anotherusers"
	"be13/sql/account/controllers/deletee"
	"be13/sql/account/controllers/login"
	"be13/sql/account/controllers/register"
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
		fmt.Println("\nMasukkan Pilihan Anda  ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			{
				newUser := entities.Users{}
				fmt.Println("masukkan nama")
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
					fmt.Println("Hallo menu utama 2")
					isRun2 := true
					for isRun2 {
						fmt.Print("\nSUB MENU:\n 1. Read Account \n 2. Update Account \n 3. Delete Account \n 4. Top-Up \n 5. Transfer \n 6. History Top-Up \n 7. History Transfer \n 8. Melihat Profil user lain \n 9. Keluar\n")
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

							}
						case 5:
							{
								transfer.Transfer(dbConnection)
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
								fmt.Println("Masukkan No Telfon user lain :")
								fmt.Scanln(&cari.No_Handphone)

								row, err := anotherusers.Search(dbConnection, cari)
								if err != nil {
									log.Fatal("\nERROR ON LOGIN DATA", err.Error())
								} else {
									fmt.Printf("\nNama : %s \nSaldo : %d\n", row.Nama, row.Saldo)
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
			}
		case 0:
			{
				isRun = false
			}
		}
	}
}
