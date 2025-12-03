package models

import "time"

type User struct {
	Id        uint
	Name      string
	Books     []Book
	Birthdate time.Time
}
