package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	PrivyID string `json:"privy_id" gorm:"not null"`
	TypeUser int `json:"type_user" gorm:"not null; default:0"`
	Email string `json:"email"`
	TeleID string `json:"tele_id"`
	Bookings []Booking `json:"bookings,omitempty" gorm:"foreignKey:UserID"`
	Notifications []Notification `json:"notifications,omitempty" gorm:"foreignKey:UserID"`
}

func (u *Users) TableName() string {
	return "users"
}
