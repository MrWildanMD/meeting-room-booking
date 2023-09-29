package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID    int     `gorm:"type:integer;not null" json:"user_id"`
	BookingID int     `gorm:"type:integer" json:"booking_id"`
	Content   string  `gorm:"type:varchar(250);not null" json:"content"`
	Status    int     `gorm:"type:integer;not null" json:"status"`
	User      Users   `gorm:"foreignkey:UserID" json:"user"`
	Booking   Booking `gorm:"foreignkey:BookingID" json:"booking"`
}

func (n *Notification) TableName() string {
	return "notification"
}
