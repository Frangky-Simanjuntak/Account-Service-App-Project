package transfer

import (
	"be13/sql/account/entities"
	"database/sql"
	"fmt"
	"log"
)

func Transfer(db *sql.DB) {

	

	// membaca data saldo pengirim dan penerima
	pengirims := entities.Users{}
	penerima := entities.Users{}

	fmt.Print(" Masukan Nomor Akun Anda : ")
	fmt.Scanln(&pengirims.No_Handphone)
	fmt.Print(" Masukan Nomor Penerima : ")
	fmt.Scanln(&penerima.No_Handphone)

	//NYIMPEN DATA PENGIRIM
	statm, err := db.Prepare("SELECT id, saldo FROM users WHERE no_handphone = ?")
	if err != nil {
		log.Fatal("GAGAL line 26", err.Error())
	}
	var in entities.Users
	errs := statm.QueryRow(pengirims.No_Handphone).Scan(&in.User_id, &in.Saldo)
	if errs != nil {
		log.Fatal("ERRORS line 33", errs.Error())
	}

	//LANJUT KE TRANSFER KE NOMOR PENERIMA

	var transfer int
	fmt.Print(" Masukan Nominal Transfer :")
	fmt.Scanln(&transfer)

	if in.Saldo < transfer {
		fmt.Println("Saldo Anda Tidak Mencukupi")
	}

	// QUERY PENERIMA
	stat, err := db.Prepare("SELECT id, saldo from users where no_handphone = ?")
	if err != nil {
		log.Fatal("Penerima Tidak Ditemukan line 46", err.Error())
	}
	var terima entities.Users
	er := stat.QueryRow(penerima.No_Handphone).Scan(&terima.User_id, &terima.Saldo)
	if er != nil {
		log.Fatal("Penerima Tidak Ditemukan line 51", er.Error())
	}

	//NGISI UNTUK PENERIMA
	stat, err = db.Prepare("INSERT INTO transfers (jumlah, pengirim_id, penerima_id) values (?, ?, ?) ")
	if err != nil {
		log.Fatal("Gagal Mengirim", err.Error())
	}
	_, error := stat.Exec(transfer, in.User_id, terima.User_id)
	if error != nil {
		log.Fatal("Gagal Ke Menu Transfer History", error.Error())
	}

	stat, err = db.Prepare("UPDATE users SET saldo = ? where no_handphone = ?")
	if err != nil {
		log.Fatal("Gagal update", err.Error())
	}

	// saldo pengirim dikurang jumlah transfer
	// saldo penerima dari transfer
	saldo := in.Saldo - transfer
	saldo2 := terima.Saldo + transfer

	_, error = stat.Exec(saldo, pengirims.No_Handphone)
	if error != nil {
		log.Fatal("Gagal", error.Error())
	}
	_, error = stat.Exec(saldo2, penerima.No_Handphone)
	if error != nil {
		log.Fatal("Gagal", error.Error())
	}

	fmt.Println("Sukses Melakukan Transfer Sebesar : Rp.",transfer)

}
