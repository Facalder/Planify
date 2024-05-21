package models

import (
	"github.com/Facalder/Planify/pkg"
	"time"
)

type EmployeeModel struct {
	Id         int
	FullName   string
	ShortName  string
	UserName   string
	Password   string
	Bio        string
	Department string
	Manager    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type TabEmployee [pkg.NMAX]EmployeeModel

var (
	NEmployee int = 0
	Employees TabEmployee
)
