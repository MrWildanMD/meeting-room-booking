package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
    gorm.Model
    UserID         int       `gorm:"type:integer;not null" json:"user_id"`
    RoomID         int       `gorm:"type:integer;not null" json:"room_id"`
    CheckIn        time.Time `gorm:"type:datetime;not null" json:"check_in"`
    CheckOut       time.Time `gorm:"type:datetime;not null" json:"check_out"`
    NumberOfGuests int       `gorm:"type:integer;not null" json:"number_of_guests"`
    BookingStatus  int       `gorm:"type:integer;not null" json:"booking_status"`
    AdditionalItem string    `gorm:"type:text" json:"additional_item"`
    ApprovalID     int       `json:"approval_id"`
    User           Users      `gorm:"foreignkey:UserID" json:"user"`
    Room           Room      `gorm:"foreignkey:RoomID" json:"room"`
    Notifications  []Notification `gorm:"foreignkey:BookingID" json:"notifications"`
}

func (b *Booking) TableName() string {
	return "booking"
}
