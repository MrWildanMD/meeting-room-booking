package models

import "gorm.io/gorm"

type Users struct {
    gorm.Model
    Name      string    `gorm:"type:varchar(100);not null" json:"name"`
    PrivyID   string    `gorm:"type:varchar(100);not null" json:"privy_id"`
    TypeUser  int       `gorm:"type:integer;not null" json:"type_user"`
    Email     string    `gorm:"type:varchar(100)" json:"email"`
    TeleID    string    `gorm:"type:varchar(100)" json:"tele_id"`
    Bookings  []Booking `gorm:"foreignkey:UserID" json:"bookings"`
    Notifications []Notification `gorm:"foreignkey:UserID" json:"notifications"`
}

func (u *Users) TableName() string {
	return "users"
}
