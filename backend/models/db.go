package models

import (
	"fmt"

	"github.com/yuhori/recipe-viewer/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB() error {
	user := configs.GetString("MySQL_USER")
	password := configs.GetString("MySQL_PASSWORD")
	endpoint := configs.GetString("MySQL_ENDPOINT")
	database := configs.GetString("MySQL_DATABASE")
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, endpoint, database)
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}

	autoMigration()
	return nil
}

func autoMigration() {
	Db.AutoMigrate(&Recipe{})
}
