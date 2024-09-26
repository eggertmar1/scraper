package models

import "time"

type League struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Country   string    `db:"country"`
	CreatedAt time.Time `db:"created_at"`
}
