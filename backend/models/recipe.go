package models

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model

	Name string `json:"name"`
	Link string `json:"link"`
}

func GetAll() ([]Recipe, error) {
	var r []Recipe
	result := Db.Find(&r)
	if result.Error != nil {
		return []Recipe{}, result.Error
	}
	return r, nil
}
