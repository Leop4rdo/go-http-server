package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open(
		sqlite.Open("test.db"),
		&gorm.Config {},
	)

	if err != nil {
		panic("Failed to connectToDatabase!")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	Database = database
}
