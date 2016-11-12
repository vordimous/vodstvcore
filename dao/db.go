package dao

import (
	"database/sql"
	"esvodsCore/models"
	"esvodsCore/util"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //import postgres
)

//DB ...
type DB struct {
	*sql.DB
}

const (
	//DbHost ...
	DbHost = "localdocker"
	//DbUser ...
	DbUser = "esvodsapi"
	//DbPassword ...
	DbPassword = "esvodsapi"
	//DbName ...
	DbName = "esvods"
)

var db *gorm.DB

//Init ...
func Init() {

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbHost, DbName)

	var err error
	//set the global variable
	db, err = gorm.Open("postgres", dbinfo)
	util.CheckErr(err, "db connect failed")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	err = db.DB().Ping()
	util.CheckErr(err, "db ping failed")

	db.LogMode(true)
}

// DbMigration ...
func DbMigration() {
	GetDB().AutoMigrate(&models.Vod{})
	GetDB().AutoMigrate(&models.Match{})
	GetDB().AutoMigrate(&models.Tag{})
	GetDB().AutoMigrate(&models.Watcher{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}
