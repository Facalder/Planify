package models

import (
	"github.com/Facalder/Planify/pkg"
	"time"
)

type task struct {
	Id          int
	Name        string
	Description string
	Priority    int // 1 (standar), 2 (low), 3 (med), 4 (high)
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type TabTask [pkg.NMAX]task

var (
	NTask int = 0
	Tasks TabTask
)
