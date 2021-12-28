package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name    string `gorm:"type:varchar(20);not null"`
	Version string `gorm:"varchar(110);not null"`
	Main    string `gorm:"varchar(110);not null"`
	Author  string `gorm:"varchar(110);not null"`
}
