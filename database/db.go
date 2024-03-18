package database

import (
	_ "github.com/go-sql-driver/mysql"
	"go-server/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() {
	var err error
	//db, err = sql.Open("mysql", "user:password@tcp(localhost:8036)/db")
	dsn := "user:password@tcp(localhost:8036)/db?charset=utf8mb4&parseTime=True&loc=Local"
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
