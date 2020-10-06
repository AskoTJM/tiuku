package main

/*
 * tiuku API
 *
 * Access the Tiuku system.
 *
 * API version: 1.0
 * Contact: asko.mattila@gmail.com
 * Routing Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

import (
	"log"
	"net/http"

	sw "github.com/AskoTJM/tiuku/api"
	"github.com/AskoTJM/tiuku/api/database"
)

func main() {

	log.Printf("Server started")
	router := sw.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
	// Check to see if we have required tables.
	// ToDo: Decide how to react, create them or just complain.
	resCheck := database.CheckIfRequiredTablesExist()
	if !resCheck {
		log.Printf("No required tables found, make sure to create them before use")
	}
	//database.ConnectToDB()
	/*
		Tiukudb, err := gorm.Open("mysql", "apiaccess:apipass@tcp(db:3306)/tiukuDB?charset=utf8mb4")
		if err != nil {
			log.Printf("Problem with connecting to database. <database/database.go->connectToDB>")
			log.Println(err)
		}
		defer Tiukudb.Close()
	*/

}
