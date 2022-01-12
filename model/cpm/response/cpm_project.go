package response

import (
	uuid "github.com/satori/go.uuid"
)

type ProjectList struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	AuthorName   string    `json:"authorName"`
	LanguageName string    `json:"languageName"`
	TypeName     string    `json:"typeName"`
}
