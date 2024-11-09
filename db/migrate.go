package db

import (
	"mailcast/helpers"
	"mailcast/models"
)

func Migrate() {
	// Migrate the schema
	err := Db.AutoMigrate(&models.Subscriber{})
	helpers.FailOnError(err, "Error connecting to the db")
}
