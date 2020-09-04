package swagger

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
type User struct {
	ID       int
	Username string
}
*/
type UserModel struct {
	ID      int `gorm:"primary_key"`
	Name    string
	Address string
}

func connectToDB() {

	db, err := gorm.Open("mysql", "testiuser:testipassword@tcp(localhost:3306)/tiukuDB?charset=utf8")
	if err != nil {
		log.Printf("Connection failed to Open")
	}
	log.Printf("Connection Established")

	db.CreateTable(&User{})

	//dsn := "testiuser:testipassword@tcp(db:3306)/tiukuDB?charset=utf8&parseTime=True"
	//db, err := gorm.Open(mysql, "testiuser:testipassword@/tiukuDB?charset=utf8&parseTime=True")
	//gorm.Open("mysql", "testiuser:testipassword@tcp(127.0.0.1:3306)/tiukuDB?charset=utf8&parseTime=True")
	//db, err := sql.Open("mysql", "testiuser:testipassword@/tiukuDB")
	//dsn := "testiuser:testipassword@tcp(127.0.0.1:3306)/tiukuDB?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}

	//db.Debug().DropTableIfExists(&UserModel{})
	//Drops table if already exists
	//db.Debug().AutoMigrate(&UserModel{})
	//Auto create table based on Model

	//fmt.Printf("...yay...")
	//log.Printf("Connection Established")

	//test := segment{Name: "Vilperi", Kukkuu: "No totta kai", createdAt: time.Now()}
	//db.Create(&test)
	//return db
}

func testCreate() {
	//test := segment{Name: "Vilperi", Kukkuu: "No totta kai", createdAt: time.Now()}
	//db.Create(&test)
}
