package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.Db
var err error

func InitDB() {
	user := config.GetString("MySQL_USER")
	password := config.GetString("MySQL_PASSWORD")
	endpoint := config.GetString("MySQL_ENDPOINT")
	database := config.GetString("MySQL_DATABASE")
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, endpoint, database)
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config())
	if err != nil {
		return err
	}

	autoMigration()
	return nil
}

func autoMigration() {
	Db.autoMigration(&Recipe{})
}
