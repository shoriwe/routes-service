package models

import uuid "github.com/satori/go.uuid"

type VehicleStatus int

const (
	EngineOn VehicleStatus = iota
	EngineOff
)

type Location struct {
	Model
	VehicleUUID uuid.UUID     `gorm:"name:vehicle_uuid;not null;" json:"vehicle_uuid"`
	Vehicle     Vehicle       `gorm:"foreignKey:vehicle_uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Latitude    float64       `gorm:"not null;" json:"latitude"`
	Longitude   float64       `gorm:"not null;" json:"longitude"`
	Speed       float64       `gorm:"not null;" json:"speed"`
	Status      VehicleStatus `gorm:"not null;" json:"status"`
}
