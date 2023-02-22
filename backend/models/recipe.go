package models

import (
	"net/url"

	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model

	Name string `json:"name"`
	Link string `json:"link"`
}

func GetAll() ([]Recipe, error) {
	var rs []Recipe
	result := Db.Find(&rs)
	if result.Error != nil {
		return []Recipe{}, result.Error
	}
	return rs, nil
}

func GetOne() (Recipe, error) {
	var r Recipe
	result := Db.First(&r)
	if result.Error != nil {
		return Recipe{}, result.Error
	}
	return r, nil
}

func Search(params url.Values) ([]Recipe, error) {
	query := Db.Model(&Recipe{})
	for k, v := range params {
		switch k {
		case "id":
			query = query.Where("id = ?", v)
		case "name":
			query = query.Where("id = ?", v)
		case "link":
			query = query.Where("id = ?", v)
		}
	}
	var rs []Recipe
	result := Db.Find(&rs)
	if result.Error != nil {
		return []Recipe{}, result.Error
	}
	return rs, nil
}

func Create(r *Recipe) error {
	result := Db.First(r)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Delete(id string) error {
	result := Db.Delete(&Recipe{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
