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
// W0rks
func ReplaceSession(user string, oldSession uint, newSession StudentSegmentSession) (bool, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseBool bool
	var responseString string
	studentData := GetStudentUser(user)
	tableToEdit := studentData.AnonID + "_sessions"
	var oldStudentSession StudentSegmentSession
	log.Printf("Session to replace is %v", oldSession)
	// Should give the last version because GORM sorts with primary key, ie highest primary key = latest addition
	if err := Tiukudb.Table(tableToEdit).Where("resource_id = ?", oldSession).Last(&oldStudentSession).Error; err != nil {
		log.Printf("Error in retrieving table %v in <database/update_database->ReplaceSession. %v \n", oldSession, err)
		responseString = "Error in retrieving table. Incorrect resource ID"
		responseBool = false
	} else {
		// Mark old one as Deleted
		log.Printf("oldStudentSession Version is now %v", oldStudentSession)
		oldStudentSession.Deleted = time.Now().Format(time.RFC3339)
		oldStudentSession.Updated = time.Now().Format(time.RFC3339)
		newSession.Version = oldStudentSession.Version + 1
		Tiukudb.Table(tableToEdit).Save(&oldStudentSession)

		if err := Tiukudb.Table(tableToEdit).Create(&newSession).Error; err != nil {
			log.Printf("Error in starting Session %v \n", err)
			responseString = "Error in creating replacing Session"
			responseBool = false
		} else {
			responseString = "Session data replaced."
			responseBool = true
		}
	}
	return responseBool, responseString
}

// Update or replace existing course data
// T0D0
func UpdateCourse(updateCourse Course) {
	if Tiukudb == nil {
		ConnectToDB()
	}
}
