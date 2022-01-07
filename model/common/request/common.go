package request

import uuid "github.com/satori/go.uuid"

type GetByUuid struct {
	UUID uuid.UUID `json:"uuid" form:"uuid"`
}
