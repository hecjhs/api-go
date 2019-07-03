package models

import (
	"fmt"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Domain struct {
	gorm.Model
	Domain   string `gorm:"type:varchar(60)"`
	Priority int    `gorm:"type:int"`
	Weight   int    `gorm:"type:int"`
}

//Init Function Initialize database
func DB_init() {
	path, _ := filepath.Abs("")
	db, err := gorm.Open("sqlite3", path+"/api/fixtures/data.db")
	defer db.Close()
	if err != nil {
		fmt.Print(err)
		panic("Failed to Connect to Database")
	}
	db.AutoMigrate(&Domain{})
	db.Create(&Domain{Domain: "alpha", Priority: 5, Weight: 5})
	db.Create(&Domain{Domain: "omega", Priority: 5, Weight: 1})
	db.Create(&Domain{Domain: "alpha", Priority: 1, Weight: 5})
}
