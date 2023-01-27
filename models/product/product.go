package product

import "time"

type Product struct {
	ID         int64     `json:"id"`
	Nama       string    `json:"nama" validate:"required"`
	KodeProduk string    `json:"kodeProduk"`
	Kuantitas  int64     `json:"kuantitas"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
}
