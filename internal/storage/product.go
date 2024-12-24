package storage

import (
	"time"
)

type Product struct {
	Id       int64      `db:"id"`
	Sku      string     `db:"sku"`
	Name     string     `db:"name"`
	Category string     `db:"category"`
	Price    int        `db:"price"`
	Created  time.Time  `db:"created_at"`
	Deleted  *time.Time `db:"deleted"`
}