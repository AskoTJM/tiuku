package database

import (
	"log"
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

	joiningStudent := GetStudentUserWithStudentID(user)
	segmentToJoin := GetSegmentDataById(segRes)
	if err := Tiukudb.Table(SchoolParticipationList).Where("anon_id = ? AND segment_id = ?", joiningStudent.AnonID, segmentToJoin.ID).Delete(&SchoolSegmentsSession{}).Error; err != nil {
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
	student := GetStudentUserWithStudentID(studentId)
	tableToEdit := student.AnonID + "_sessions"

	Tiukudb.Table(tableToEdit).Where("id = ?", sessionID).Updates(StudentSegmentSession{Deleted: time.Now().Format(time.RFC3339)})
	response = "Deleted Session"
	return response
}

// Delete {student}, remove identifying data
// W1P
func DeleteStudentUser(studentId string) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	//var responseString string
	//var responseCode int
	var responseBool bool
	var tempStudent StudentUser
	if err := Tiukudb.Table(StudentsTableToEdit).Where("student_id = ?", studentId).Find(&tempStudent).Error; err != nil {
		log.Printf("Error. Can't find Student data. <database/delete_database.go->DeleteStudentUser> %v \n", err)
		responseBool = false
	} else {
		// Remove identifying data.
		tempStudent.StudentID = ""
		tempStudent.StudentEmail = ""
		tempStudent.StudentName = ""
		tempStudent.StudentClass = ""
		Tiukudb.Save(tempStudent)
		// Archive Sessions

		// Remove personal tables
		Tiukudb.DropTableIfExists(tempStudent.AnonID + "_sessions_archived")
		// Should Archive sessions to Archive_Data before
		Tiukudb.DropTableIfExists(tempStudent.AnonID + "_sessions")
		// Should remove user from active segments participation
		Tiukudb.DropTableIfExists(tempStudent.AnonID + "_segments")
		responseBool = true
	}
	return responseBool
}
