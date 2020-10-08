package database

import (
	"log"
	"time"
)

/*
// update_database.go
// Description:
*/

// Stop Active Session
// ToDo : Needs way to check if that are isn't value in EndTime already
func StopActiveSession(student string, editSession uint) string {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var response string
	var tempSession StudentSegmentSession
	studentNow := GetStudentUser(student)
	tableToEdit := studentNow.AnonID + "_sessions"
	//tempSession = GetSession(studentNow, editSession)
	// = GetOpenSession(editSession)
	num := Tiukudb.Table(tableToEdit).Where("end_time != ?", "").Last(&tempSession)
	log.Println(num)

	err := Tiukudb.Table(tableToEdit).Where("id = ?", editSession).Updates(StudentSegmentSession{EndTime: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)})
	// Gets error but works, weird
	if err != nil {
		log.Println("Error updating data for session. database/update_database->StopActiveSession")
		log.Println(err)
		return "Error"
	}

	response = "Session stopped"
	return response
	//}
}

// Update or replace existing course data
// status: work in progress
func UpdateCourse(updateCourse Course) {
	if Tiukudb == nil {
		ConnectToDB()
	}
}
