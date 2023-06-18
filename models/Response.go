package models

import "gorm.io/datatypes"

type Response struct {
	Id              int               `json:"id" gorm:"primaryKey"`
	Name            string            `json:"name,,omitempty"`
	Method          string            `json:"method,omitemptyhod,omitempty"`
	Url             string            `json:"url,omitempty"`
	Data            datatypes.JSON    `json:"body,omitempty"`
	OriginalRequest []OriginalRequest `json:"originalRequest,omitempty"`
	RequestId       int               `json:"request_id,omitempty"`
	CollectionId    int               `json:"collection_id,omitempty"`
}
type OriginalRequest struct {
	Method     string         `json:"method,omitempty"`
	Header     string         `json:"hearder,omitempty"`
	Data       datatypes.JSON `json:"body,omitempty"`
	Url        string         `json:"url,omitempty"`
	ResponseId int            `json:"response_id,omitempty"`
}
