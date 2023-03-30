package models

import uuid "github.com/satori/go.uuid"

type APIKey struct {
	Model
	VehicleUUID uuid.UUID `gorm:"name:vehicle_uuid;not null;" json:"vehicle_uuid"`
	Vehicle     Vehicle   `gorm:"foreignKey:vehicle_uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Key         string    `gorm:"not null;" json:"apiKey"`
}
