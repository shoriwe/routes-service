package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type RouteStatus int

const (
	RouteInProgress RouteStatus = iota
	RouteCancelled
	RouteCompleted
)

type Route struct {
	Model
	VehicleUUID uuid.UUID   `gorm:"not null;" json:"vehicle_uuid"`
	Vehicle     Vehicle     `gorm:"foreignKey:VehicleUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Started     time.Time   `gorm:"default:CURRENT_TIMESTAMP;not null;" json:"started"`
	Completed   *time.Time  `json:"completed,omitempty"`
	Status      RouteStatus `gorm:"not null;" json:"status"`
	StartNode   int         `gorm:"not null;" json:"startNode"`
	EndNode     int         `gorm:"not null;" json:"endNode"`
}
