package models

import (
	"crypto/rand"
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type APIKey struct {
	Model
	VehicleUUID uuid.UUID `gorm:"not null;" json:"vehicleUUID"`
	Vehicle     Vehicle   `gorm:"foreignKey:VehicleUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Key         string    `gorm:"not null;" json:"apiKey"`
}

func (ak *APIKey) BeforeSave(tx *gorm.DB) error {
	mErr := ak.Model.BeforeSave(tx)
	if mErr != nil {
		return mErr
	}
	if len(ak.Key) == 0 {
		var chunk [64]byte
		_, err := rand.Read(chunk[:])
		if err != nil {
			panic(err)
		}
		ak.Key = hex.EncodeToString(chunk[:])
	}
	return nil
}
