package models

type Folder struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Requests     []Request `json:"requests,omitempty"`
	CollectionId int       `json:"collection_id,omitempty"`
}
