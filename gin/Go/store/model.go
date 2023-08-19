package store

import "time"

type User struct {
	ID        uint
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
