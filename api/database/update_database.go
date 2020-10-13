package database

import (
	"time"
)

/*
// update_database.go
// Description:
*/

// Stop Active Session
// ToDo : Needs way to check if that are isn't value in EndTime already
func StopActiveSession(student string, editSession uint) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var response bool
	//var tempSession StudentSegmentSession
	studentNow := GetStudentUser(student)
	tableToEdit := studentNow.AnonID + "_sessions"
	//Tiukudb.Table(tableToEdit).Where("end_time != ?", "").Last(&tempSession)
	//log.Println(num)

	Tiukudb.Table(tableToEdit).Where("id = ?", editSession).Updates(StudentSegmentSession{EndTime: time.Now().Format(time.RFC3339), Updated: time.Now().Format(time.RFC3339)})

	response = true
	return response
	//}
}

// Replace Session data
// T0D0
func ReplaceSession(oldSession uint, newSession StudentSegmentSession) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var response bool
	//oldSession := Tiukudb.Table(tableToEdit)

	response = true
	return response
}

// Update or replace existing course data
// status: work in progress
func UpdateCourse(updateCourse Course) {
	if Tiukudb == nil {
		ConnectToDB()
	}
}
