package model

import "time"

type Project struct {
	ID        int    `gorm:"primary_key;type:int(4);auto_increment"`
	Name      string `gorm:"type:varchar(20);not null;comment:'项目名'"`
	Author    string `gorm:"type:varchar(110);not null;comment:'创建者'"`
	Type      int    `gorm:"type:enum('1', '2', '9');default:'9';comment:'组件库1/页面2,其他9'"`
	Language  int    `gorm:"type:enum('vue2', 'vue3', 'react', 'js');default:'js'"`
	CreatedAt time.Time
}
