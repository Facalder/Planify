package models

import (
	"github.com/Facalder/Planify/pkg"
	"time"
)

type ManagerModel struct {
	Id         int
	FullName   string
	ShortName  string
	UserName   string
	Password   string
	Bio        string
	Department string
	CreatedAt  time.Time
	UpdateAt   time.Time
	DeletedAt  time.Time
}

type TabManager [pkg.NMAX]ManagerModel

var (
	NManager int = 0
	Managers TabManager
)
