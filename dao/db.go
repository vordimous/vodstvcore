package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"vodstv/core"
	"vodstv/core/models"

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

	dbinfo := os.Getenv("DATABASE_URL")

	if dbinfo == "" {
		connectionName := mustGetenv("CLOUDSQL_CONNECTION_NAME")
		user := mustGetenv("CLOUDSQL_USER")
		password := os.Getenv("CLOUDSQL_PASSWORD")

		dbinfo = fmt.Sprintf("%s:%s@cloudsql(%s)/vodstv", user, password, connectionName)
	}

	fmt.Println("DATABASE_URL: " + dbinfo)

	var err error
	//set the global variable
	db, err = gorm.Open("postgres", dbinfo)
	core.CheckErr(err, "db connect failed")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	err = db.DB().Ping()
	core.CheckErr(err, "db ping failed")

	db.LogMode(true)
}

// DbMigration ...
func DbMigration() {
	GetDB().AutoMigrate(&models.Vod{})
	GetDB().AutoMigrate(&models.Match{})
	GetDB().AutoMigrate(&models.Tag{})
	GetDB().AutoMigrate(&models.Watcher{})
	GetDB().AutoMigrate(&models.Feed{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
