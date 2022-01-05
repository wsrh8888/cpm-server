package model

import "time"

type User struct {
	ID        int `gorm:"primary_key;type:int(4);auto_increment"`
	Email     string
	Password  string `gorm:"type:varchar(110);not null;comment:'创建者'"`
	CreatedAt time.Time
}

func (User) TableName() string {
	return "user"
}
