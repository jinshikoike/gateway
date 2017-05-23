package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connection connect to database server
func Connection(user, password, host, port, dbname string) (*gorm.DB, error) {
	connectionStr := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname

	db, err := gorm.Open("mysql", connectionStr)
	db.LogMode(true)
	if err != nil {
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}
