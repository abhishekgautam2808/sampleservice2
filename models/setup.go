package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "ginuser"
	password = "postgres"
	dbname   = "userdb"
)

// To reference the gorm and do operations with it.
var DB *gorm.DB

//ConnectDatabase  is used to connect to the database and to create the table
func ConnectDatabase() {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	database, err := gorm.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	database.AutoMigrate(&User{})
	DB = database
}
