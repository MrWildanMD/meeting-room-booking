package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID int `json:"user_id"`
	BookingID int `json:"booking_id,omitempty"`
	Content string `json:"content" gorm:"not null"`
	Status int `json:"status" gorm:"not null"`
	Timestamp time.Time `json:"timestamp" gorm:"not null"`
}

func (n *Notification) TableName() string {
	return "notification"
}
