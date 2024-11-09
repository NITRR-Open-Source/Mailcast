package db

import (
	"mailcast/helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func DBConn(dbUrl string) *gorm.DB {
	Db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	helpers.FailOnError(err, "Error connecting to the db")
	return Db
}
