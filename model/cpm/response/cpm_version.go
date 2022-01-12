package response

import "time"

type VersionList struct {
	Version   string    `json:"version"`
	Publisher string    `json:"publisher"`
	CreatedAt time.Time `json:"createdAt"`
}
