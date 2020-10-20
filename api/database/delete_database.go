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

// Remove Student user From segment participation
// W0rks errorFlag fixed
func DeleteStudentFromSegment(user string, segRes uint) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var errorFlag bool = false
	joiningStudent := GetStudentUserWithStudentID(user)
	segmentToJoin := GetSegmentDataById(segRes)
	// Remove from Schools Participation table
	if err := Tiukudb.Table(SchoolParticipationList).Where("anon_id = ? AND segment_id = ?", joiningStudent.AnonID, segmentToJoin.ID).Delete(&SchoolSegmentsSession{}).Error; err != nil {
		errorFlag = true
	}
	return errorFlag
}

// Remove Student user From all segments participation
// T35T	errorFlag fixed
func DeleteStudentFromAllSegments(user string) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var errorFlag bool = false
	joiningStudent := GetStudentUserWithStudentID(user)
	if err := Tiukudb.Table(SchoolParticipationList).Where("anon_id = ?", joiningStudent.AnonID).Delete(&SchoolSegmentsSession{}).Error; err != nil {
		log.Printf("Error. Removing student user from all segments in <database/delete_database->DeleteStudentFromAllSegments> %v \n", err)
		errorFlag = true
	}
	return errorFlag
}

// Remove Session from Student User, SoftDelete
// W0rks
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
// W1P errorFlag fixed,  Add Transfer of non-Archived Sessions and Remove from Segment
func DeleteStudentUser(studentId string) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	//var responseString string
	//var responseCode int
	var errorFlag bool = false
	var tempStudent StudentUser
	if err := Tiukudb.Table(StudentsTableToEdit).Where("student_id = ?", studentId).Find(&tempStudent).Error; err != nil {
		log.Printf("Error. Can't find Student data. <database/delete_database.go->DeleteStudentUser> %v \n", err)
		errorFlag = true
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
		// T0D0 Archive possible Sessions
		Tiukudb.DropTableIfExists(tempStudent.AnonID + "_sessions")
		// Should remove user from active segments participation
		Tiukudb.DropTableIfExists(tempStudent.AnonID + "_segments")
		//errorFlag = false
	}
	return errorFlag
}

// Delete User Session from User Session table, used after Archiving
// W1P errorFlag fixed move to delete_database.go
func DeleteSessionsFromUsersSessionsTable(user string, segId uint) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var errorFlag bool = false

	return errorFlag
}
