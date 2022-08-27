package notebook_db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(database_path string) *gorm.DB {
	/* Setup the Database */

	// Open database
	db, err := gorm.Open(sqlite.Open(database_path), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	// Generate table structure
	db.AutoMigrate(&User{}) // user table
	db.AutoMigrate(&Note{}) // notes table

	return db
}

