package models

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model

	Name string `json:"name"`
	Link string `json:"link"`
}
