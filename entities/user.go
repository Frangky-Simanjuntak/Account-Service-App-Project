package entities

import "time"

type Users struct {
	User_id, Password, Saldo int
	Nama, No_Handphone       string
}
type Topup struct {
	Topup_id, User_id, Jumlah int
	Created_at, Updated_at    time.Time
}
