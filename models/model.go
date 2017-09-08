package models

type LateReason struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}