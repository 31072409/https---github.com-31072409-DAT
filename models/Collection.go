package models

type Collection struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Folders   []Folder   `json:"folders,omitempty"`
	Request   []Request  `json:"requests,omitempty"`
	Responses []Response `json:"repsonses,omitempty"`
}
