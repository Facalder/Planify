package models

import (
	"github.com/Facalder/Planify/pkg"
	"time"
)

type AdminModel struct {
	Id        int
	FullName  string
	ShortName string
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}

type TabAdmin [pkg.NMAX]AdminModel

var (
	NAdmin int = 0
	Admins TabAdmin
)
