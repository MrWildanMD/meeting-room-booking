package models

import (
	"gorm.io/gorm"
)

type Room struct {
    gorm.Model
    Title       string    `gorm:"type:varchar(100);not null" json:"title"`
    Description string    `gorm:"type:varchar(100);not null" json:"description"`
    Location    string    `gorm:"type:varchar(100);not null" json:"location"`
    Type        int       `gorm:"type:integer;not null" json:"type"`
    Capacity    int       `gorm:"type:integer;not null" json:"capacity"`
    Status      int       `gorm:"type:integer;not null" json:"status"`
    Offices     []Office  `gorm:"foreignkey:RoomID" json:"offices"`
    Bookings    []Booking `gorm:"foreignkey:RoomID" json:"bookings"`
}

func (r *Room) TableName() string {
	return "room"
}