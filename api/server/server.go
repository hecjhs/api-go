package server

import (
	"os"
	"path/filepath"

	"github.com/hecjhs/api-go/api/models"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

var path, _ = filepath.Abs("")
var db, _ = gorm.Open("sqlite3", path+"/api/fixtures/data.db")

// SetUp server
func SetUp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	models.DB_init()
	return app
}

// RunServer should start server
func RunServer(app *iris.Application) {
	app.Run(
		iris.Addr(os.Getenv("PORT")),
	)
}
