package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Title string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Location string `json:"location" gorm:"not null"`
	Type int `json:"type" gorm:"not null"`
	Capacity int `json:"capacity" gorm:"not null"`
	Status int `json:"status" gorm:"not null"`
}

func (r *Room) TableName() string {
	return "room"
}