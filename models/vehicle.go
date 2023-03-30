package models

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"
)

var plateRegex = regexp.MustCompile(`(?m)^([A-Z]{3}\d{3}|[A-Z]{2}\d{4}|[A-Z]{2}\d{3}[A-Z]|[A-Z]\d{4,5})$`)

type Vehicle struct {
	Model
	Name  string `gorm:"not null;" json:"name"`
	Plate string `gorm:"not null;unique" json:"plate"`
}

func (v *Vehicle) BeforeSave(tx *gorm.DB) error {
	mErr := v.Model.BeforeSave(tx)
	if mErr != nil {
		return mErr
	}
	if len(v.Name) == 0 {
		return fmt.Errorf("vehicle no name provided")
	}
	if !plateRegex.MatchString(v.Plate) {
		return fmt.Errorf("vehicle plate must be a valid colombian plate")
	}
	return nil
}
