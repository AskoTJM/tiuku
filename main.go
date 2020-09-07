package main

/*
 * StudentAPI
 *
 * API for Students to acccess the Tiuku system.
 *
 * API version: 1.0
 * Contact: asko.mattila@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	
	sw "./go"
)

/*
type Env struct {
	db *gorm.DB
}

*/

func main() {

	log.Printf("Server started")
	router := sw.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
