package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/nhatvu148/go_reportmaker/pkg/setting"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
