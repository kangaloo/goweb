package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kangaloo/goweb/config"
	"log"
)

var (
	db     *gorm.DB
	dbType = "mysql"
)

func SetDB(database *gorm.DB) {
	db = database
}

func ConnectToDB() *gorm.DB {
	log.Printf("connecting to databases ...\n")
	log.Println(config.GetMysqlDSN())

	db, err := gorm.Open(dbType, config.GetMysqlDSN())
	if err != nil {
		panic("Failed to connect to databases " + err.Error())
	}

	db.SingularTable(true)
	return db
}
