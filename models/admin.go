package models

import "time"

type AdminModel struct {
	Id        int64
	FullName  string
	ShortName string
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}
