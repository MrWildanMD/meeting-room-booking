package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name          string         `json:"name" gorm:"not null"`
	PrivyID       string         `json:"privy_id" binding:"required" gorm:"not null"`
	TypeUser      int            `json:"type_user,omitempty" gorm:"not null; default:0"`
	Email         string         `json:"email" binding:"required"`
	TeleID        string         `json:"tele_id,omitempty"`
	Bookings      []Booking      `json:"bookings,omitempty" gorm:"foreignKey:UserID"`
	Notifications []Notification `json:"notifications,omitempty" gorm:"foreignKey:UserID"`
}

func (u *Users) TableName() string {
	return "users"
}
