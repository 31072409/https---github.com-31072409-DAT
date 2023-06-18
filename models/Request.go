package models

import "gorm.io/datatypes"

type Request struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name,omitempty"`
	Data         datatypes.JSON `json:"body,omitempty"`
	Method       string         `json:"method,omitempty"`
	Url          string         `json:"url,omitempty"`
	FolderId     *int           `json:"folder_id,omitempty,omitempty"`
	CollectionId int            `json:"collection_id,omitempty"`
	Responses    []Response     `json:"repsonses,omitempty,omitempty"`
}
