package entities

import "time"

type Transfer struct {
	Transfer_id, Jumlah, Penerima_id, Pengirim_id int
	Created_at, Updated_at                        time.Time
}
