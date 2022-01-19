package response

import (
	"cpm/model/cpm"
	"time"

	uuid "github.com/satori/go.uuid"
)

type ProjectList struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	AuthorName   string    `json:"authorName"`
	LanguageName string    `json:"languageName"`
	TypeName     string    `json:"typeName"`
}
type VersionInfo struct {
	CreatedAt     time.Time `json:"createdAt"`
	Description   string    `json:"description"`
	Keywords      string    `json:"keywords"`
	Version       string    `json:"version"`
	PublisherName string    `json:"publisherName"`
}

type CpmProjectAllInfo struct {
	CpmProject    ProjectList     `json:"cpmProject"`
	CpmVersionNew VersionInfo     `json:"cpmVersionNew"`
	CpmImport     []cpm.CpmImport `json:"cpmImport"`
}
