package database

import (
	"time"
)

/*
// delete_database.go
// Removing stuff from database
//
*/

// Joining Student user to segment
// status: Works
func DeleteStudentFromSegment(user string, segRes uint) string {
	if Tiukudb == nil {
		ConnectToDB()
	}

	joiningStudent := GetStudentUser(user)
	segmentToJoin := GetSegmentDataById(segRes)
	if err := Tiukudb.Table(EnrollmentSegmentList).Where("anon_id = ? AND segment_id = ?", joiningStudent.AnonID, segmentToJoin.ID).Delete(&SchoolSegmentsSession{}).Error; err != nil {
		response := "Error removing Student from Segment. <database/update_database->RemoveStudentFromSegment>"
		return response
	}
	response := "Removed student from Segment"
	return response
}

func DeleteSessionFromStudent(studentId string, sessionID uint) string {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var response string
	student := GetStudentUser(studentId)
	tableToEdit := student.AnonID + "_sessions"

	Tiukudb.Table(tableToEdit).Where("id = ?", sessionID).Updates(StudentSegmentSession{Deleted: time.Now().Format(time.RFC3339)})
	response = "Deleted Session"
	return response
}
