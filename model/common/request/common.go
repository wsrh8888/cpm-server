package request

import uuid "github.com/satori/go.uuid"

type GetByUuid struct {
	UUID uuid.UUID `json:"uuid" form:"uuid"`
}

type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}
