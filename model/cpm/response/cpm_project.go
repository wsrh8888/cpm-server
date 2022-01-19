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

type CpmProjectAllInfo struct {
	CpmProject ProjectList
	// CpmVersion cpm.CpmVersion
	// CpmType    cpm.CpmType
}
