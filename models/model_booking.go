package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID         int       `json:"user_id"`
	RoomID         int       `json:"room_id"`
	CheckIn        time.Time `json:"check_in" gorm:"not null"`
	CheckOut       time.Time `json:"check_out" gorm:"not null"`
	NumberOfGuest  int       `json:"number_of_guest" gorm:"not null"`
	BookingStatus  int       `json:"booking_status" gorm:"not null"`
	AdditionalItem string    `json:"additional_item" gorm:"type:text"`
	ApprovalID     int       `json:"approval_id,omitempty"`
}

func (b *Booking) TableName() string {
	return "booking"
}
