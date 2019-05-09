package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	db     *gorm.DB
	dbType = "mysql"
	dsn    = "root:123456@tcp(127.0.0.1:3306)/goweb?charset=utf8&parseTime=True&loc=Local"
)

func SetDB(database *gorm.DB) {
	db = database
}

// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

func ConnectToDB() *gorm.DB {
	log.Printf("connecting to databases ...\n")
	db, err := gorm.Open(dbType, dsn)
	if err != nil {
		panic("Failed to connect to databases " + err.Error())
	}

	db.SingularTable(true)
	return db
}
