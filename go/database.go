package swagger

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectToDB() {
	dsn := "testiuser:testipassword@tcp(db:3306)/tiukuDBharset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error with Gorm connecting mySql!")
	}
	//return db
}
