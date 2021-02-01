package models

import (
	"gorm.io/gorm"
)

// ColorModel db model
type ColorModel struct {
	gorm.Model
	BinID int
	Label string
	Hex   string
	Lab   map[string]interface{}
	RGB   map[string]interface{}
	HSB   map[string]interface{}
}
