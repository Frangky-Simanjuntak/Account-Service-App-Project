package main

import (
	"be13/sql/account/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection := config.ConnectToDB()

	defer dbConnection.Close() // menutup koneksi

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
				fmt.Println("ADD ACOUNT")
			}
		case 2:
			{
				
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
							fmt.Println("HALO SUB MENU 1")
						}
					case 2:
						{
							fmt.Println("HALO SUB MENU 2")
						}
					case 3:
						{
							fmt.Println("HALO SUB MENU 3")
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
							fmt.Println("HALO SUB MENU 8")
						}
					case 9:
						{
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
