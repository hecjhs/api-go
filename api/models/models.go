package models

import (
	"fmt"
	"os"
	"path"
	"runtime"

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
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	fmt.Printf("lalallal %v\n", dir)
	// path, _ := filepath.Abs("")
	d, _ := os.Getwd()
	fmt.Printf("current dir %v \n", d)
	// fmt.Printf(path)
	db, err := gorm.Open("sqlite3", dir+"/fixtures/data.db")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Domain{})
	db.Create(&Domain{Domain: "alpha", Priority: 5, Weight: 5})
	db.Create(&Domain{Domain: "omega", Priority: 5, Weight: 1})
	db.Create(&Domain{Domain: "beta", Priority: 1, Weight: 5})
}
