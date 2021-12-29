package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDatabase - returns a pointer to a database object
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	dbUsername := "postgres" //os.Getenv("DB_USERNAME")
	dbPassword := "postgres" //os.Getenv("DB_PASSWORD")
	dbHost := "172.17.0.2"   //os.Getenv("DB_HOST") //172.17.0.2
	dbTable := "comments"    //os.Getenv("DB_TABLE")
	dbPort := "5432"         //os.Getenv("DB_PORT") //5432

	// connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)

	connectString := fmt.Sprintf("host=%s port=%s dbname=%s password=%s user=%s sslmode=disable", dbHost, dbPort, dbTable, dbPassword, dbUsername)

	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil
}
