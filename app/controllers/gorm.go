package controllers

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"github.com/revel/revel/modules/db/app"
	"test/app/models"
)

//==============================================================================
// DB
//================================================
var (
	DB gorm.DB
)

//==============================================================================
func InitDB() {
	revel.INFO.Println("Initializing DB..")

	db.Init()

	driver, _ := revel.Config.String("db.driver")
	spec, _ := revel.Config.String("db.spec")

	var err error
	DB, err = gorm.Open(driver, spec)
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	DB.LogMode(true)

	DB.AutoMigrate(models.Page{})

	revel.INFO.Println("DB initialized")
}

//==============================================================================
