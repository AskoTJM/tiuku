package database

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
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
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error removing Student user from Participation list. <database/delete_database.go->DeleteStudentFromSegment> %v \n", err)
			errorFlag = true
		}
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
	var tempSchoolSegment SchoolSegmentsSession
	joiningStudent := GetStudentUserWithStudentID(user)
	log.Println(joiningStudent)
	if err := Tiukudb.Table(SchoolParticipationList).Where("anon_id = ?", joiningStudent.AnonID).Delete(&tempSchoolSegment).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error. Removing student user from all segments in <database/delete_database->DeleteStudentFromAllSegments> %v \n", err)
			errorFlag = true
		}
	}
	return errorFlag
}

// Remove Session from Student User, SoftDelete
// W0rks
func DeleteSessionFromStudent(studentId string, sessionID uint) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var errorFlag bool = false
	student := GetStudentUserWithStudentID(studentId)
	tableToEdit := student.AnonID + "_sessions"

	if err := Tiukudb.Table(tableToEdit).Where("id = ?", sessionID).Updates(StudentSegmentSession{Deleted: time.Now().Format(time.RFC3339)}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error. Removing Session failed. <database/delete_database.go->DeleteSessionFromStudent> %v \n", err)
			errorFlag = true
		}
	}
	return errorFlag
}

// Delete {student}, remove identifying data
// W0rks errorFlag fixed,  Add Transfer of non-Archived Sessions and Remove from Segment
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
