package deletee

import (
	"database/sql"
	"fmt"
	"log"
)

func DeleteAkun(db *sql.DB, no_Handphonee string) (sql.Result, error) {
	var query = "DELETE FROM users WHERE no_handphone = ?" //tanda tanya diambil dari baris 114 keatas
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr Qouery", errPrepare.Error())
	}
	result, errExec := statement.Exec(no_Handphonee) //jumlah dan urutan disesuaikan dngan tanda tanya komen mysql
	if errExec != nil {
		log.Fatal("erorr Exec ", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("DELETE AKUN BERHASIL")
		} else {
			fmt.Println("DELETE GAGAL !")
		}
	}
	return result, nil
}
