package database

import (
	"go-server/entities"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	//db, err = sql.Open("mysql", "user:password@tcp(localhost:8036)/db")
	dsn := "user:password@tcp(database:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&entities.Game{})

	log.Println("Connected to the MySQL database")
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
