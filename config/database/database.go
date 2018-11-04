package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// initialize environment variables
var (
	DBNAME             = getEnv("DBNAME", "vglib")
	DBUSER             = getEnv("DBUSER", "vglibdev")
	DBPASS             = getEnv("DBPASS", "abc123vglib")
	dbConnectionString = "user=" + DBUSER + " password=" + DBPASS + " dbname=" + DBNAME + " sslmode=disable"
	dialect            = "postgres"
)

// this pointer variable is initialized with the database connection and can then be called
// for a dependency injection in other packages
var GormConn *gorm.DB

func init() {
	GormConn = openDatabaseConnection(dialect, dbConnectionString)
}

// private functions for connecting to the database
func openDatabaseConnection(dialect string, dbConnectionString string) *gorm.DB {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(dialect, dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// MigrateDependencyTables - These functions should be used in their respective order
// and should only be ran to create or recreate the tables in the database.
// func MigrateDependencyTables() {
// 	GormConn.DropTableIfExists(
// 		&businesslogic.Account{}, &businesslogic.Developer{},
// 		&businesslogic.People{}, &businesslogic.Publisher{})
// 	GormConn.AutoMigrate(
// 		&businesslogic.Account{}, &businesslogic.Developer{},
// 		&businesslogic.People{}, &businesslogic.Publisher{})
// }

// // MigrateTables - as above
// func MigrateTables() {
// 	GormConn.DropTableIfExists(
// 		&businesslogic.Game{}, &businesslogic.Character{},
// 		&businesslogic.History{}, &businesslogic.Search{})
// 	GormConn.AutoMigrate(
// 		&businesslogic.Game{}, &businesslogic.Character{},
// 		&businesslogic.History{}, &businesslogic.Search{})
// }
