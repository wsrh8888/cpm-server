package response

import (
	uuid "github.com/satori/go.uuid"
)

type ProjectList struct {
	UUID         uuid.UUID `json:"uuid"`
	Name         string    `json:"name"`
	AuthorName   string    `json:"authorName"`
	LanguageName string    `json:"languageName"`
	TypeName     string    `json:"typeName"`
}
