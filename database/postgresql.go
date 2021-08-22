package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Dbcon     *gorm.DB
	Errdb     error
	dbuser    string
	dbpass    string
	dbname    string
	host      string
	dbport    int
	dbdebug   bool
	dbtype    string
	sslmode   string
	dbtimeout string
)

func init() {

	dbuser = os.Getenv("user")
	dbpass = os.Getenv("password")
	dbname = os.Getenv("dbname")
	host = os.Getenv("host")
	dbport = 5432
	if DbOpen() != nil {
		log.Println("Failed Open db Postgres")
	}
}

// DbOpenDatabase ...
func DbOpen() error {
	args := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s ", host, dbport, dbuser, dbpass, dbname)
	Dbcon, Errdb = gorm.Open(postgres.Open(args), &gorm.Config{})
	if Errdb != nil {
		log.Fatalln("open db Err ", Errdb)
		return Errdb
	}

	return nil
}

// GetDbCon ...
func GetDbCon() *gorm.DB {

	return Dbcon
}
