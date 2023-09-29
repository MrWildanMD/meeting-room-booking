package models

import "gorm.io/gorm"

type Office struct {
	gorm.Model
	RoomID      int    `gorm:"type:integer;not null" json:"room_id"`
	Name        string `gorm:"type:varchar;not null" json:"name"`
	Description string `gorm:"type:varchar;not null" json:"description"`
	Room        Room   `gorm:"foreignkey:RoomID" json:"room"`
}
