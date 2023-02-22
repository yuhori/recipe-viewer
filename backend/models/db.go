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
	user := configs.GetString("MYSQL_USER")
	password := configs.GetString("MYSQL_PASSWORD")
	endpoint := configs.GetString("MYSQL_ENDPOINT")
	database := configs.GetString("MYSQL_DATABASE")
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
