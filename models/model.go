package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Model struct {
	UUID      uuid.UUID  `gorm:"primaryKey" json:"uuid"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (m *Model) BeforeSave(tx *gorm.DB) error {
	if m.UUID == uuid.Nil {
		m.UUID = uuid.NewV4()
	}
	return nil
}
