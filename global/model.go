package global

import (
	"time"
)

type CPM_MODEL struct {
	ID        uint      `gorm:"primary_key"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
