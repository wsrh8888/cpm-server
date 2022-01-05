package model

import "time"

type Version struct {
	ID          int    `gorm:"primary_key;type:int(5);auto_increment"`
	ProjectId   string `gorm:"type:varchar(20);not null;comment:'项目ID'"`
	Version     string `gorm:"type:varchar(20);not null;comment:'版本号'"`
	Publisher   string `gorm:"type:varchar(110);not null;comment:'发布者'"`
	Description string `gorm:"type:varchar(20);not null;comment:'描述信息'"`
	Keywords    string `gorm:"type:varchar(20);not null;comment:'关键字'"`
	CreatedAt   time.Time
}
