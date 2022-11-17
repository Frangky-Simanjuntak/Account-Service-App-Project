package transfer

import (
	"be13/sql/account/entities"
	"database/sql"
	"fmt"
)

func Transfer(db *sql.DB) {

	fmt.Println("==============")
	fmt.Println("FITUR TRANSFER")
	fmt.Println("==============")

	// membaca data saldo pengirim dan penerima
	pengirims := entities.Users{}
	penerima := entities.Users{}

	fmt.Println("Masukan Nomor Akun Anda :")
	fmt.Scanln(&pengirims.No_Handphone)
	fmt.Println("Masukan Nomor Penerima :")
	fmt.Scanln(&penerima.No_Handphone)

	//NYIMPEN DATA PENGIRIM
	statm, err := db.Prepare("SELECT id, saldo FROM users WHERE no_handphone = ?")
	if err != nil {
		fmt.Println("GAGAL line 26", err.Error())
	}
	var in entities.Users
	errs := statm.QueryRow(pengirims.No_Handphone).Scan(&in.User_id, &in.Saldo)
	if errs != nil {
		fmt.Println("ERRORS line 31", errs.Error())
	}

	//LANJUT KE TRANSFER KE NOMOR PENERIMA

	var transfer int
	fmt.Println("Masukan Nominal Transfer :")
	fmt.Scanln(&transfer)

	if in.Saldo < transfer {
		fmt.Println("Saldo Anda Tidak Mencukupi")
	}

	// QUERY PENERIMA
	stat, err := db.Prepare("SELECT id, saldo from users where no_handphone = ?")
	if err != nil {
		fmt.Println("Penerima Tidak Ditemukan line 46", err.Error())
	}
	var terima entities.Users
	er := stat.QueryRow(penerima.No_Handphone).Scan(&terima.User_id, &terima.Saldo)
	if er != nil {
		fmt.Println("Penerima Tidak Ditemukan line 51", er.Error())
	}

	//NGISI UNTUK PENERIMA
	stat, err = db.Prepare("INSERT INTO transfers (jumlah, pengirim_id, penerima_id) values (?, ?, ?) ")
	if err != nil {
		fmt.Println("Gagal Mengirim", err.Error())
	}
	_, error := stat.Exec(transfer, in.User_id, terima.User_id)
	if error != nil {
		fmt.Println("Gagal Ke Menu Transfer History", error.Error())
	}

	stat, err = db.Prepare("UPDATE users SET saldo = ? where no_handphone = ?")
	if err != nil {
		fmt.Println("Gagal update", err.Error())
	}

	// saldo pengirim dikurang jumlah transfer
	// saldo penerima dari transfer
	saldo := in.Saldo - transfer
	saldo2 := terima.Saldo + transfer

	_, error = stat.Exec(saldo, pengirims.No_Handphone)
	if error != nil {
		fmt.Println("Gagal", error.Error())
	}
	_, error = stat.Exec(saldo2, penerima.No_Handphone)
	if error != nil {
		fmt.Println("Gagal", error.Error())
	}

	fmt.Println("Sukses Melakukan Transfer Sebesar : Rp.", transfer)

}
