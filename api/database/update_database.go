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
	studentNow := GetStudentUserWithStudentID(student)
	tableToEdit := studentNow.AnonID + "_sessions"
	//Tiukudb.Table(tableToEdit).Where("end_time != ?", "").Last(&tempSession)
	//log.Println(num)

	Tiukudb.Table(tableToEdit).Where("id = ?", editSession).Updates(StudentSegmentSession{EndTime: time.Now().Format(time.RFC3339), Updated: time.Now().Format(time.RFC3339)})

	response = true
	return response
	//}
}

// Replace Session data
// W0rks errorFlag fixed
func ReplaceSession(user string, oldSession uint, newSession StudentSegmentSession) (string, bool) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var errorFlag bool = false
	var responseString string
	studentData := GetStudentUserWithStudentID(user)
	tableToEdit := studentData.AnonID + "_sessions"
	var oldStudentSession StudentSegmentSession
	//log.Printf("Session to replace is %v", oldSession)
	// Should give the last version because GORM sorts with primary key, ie highest primary key = latest addition
	if err := Tiukudb.Table(tableToEdit).Where("resource_id = ?", oldSession).Last(&oldStudentSession).Error; err != nil {
		log.Printf("Error in retrieving table %v in <database/update_database->ReplaceSession. %v \n", oldSession, err)
		responseString = "Error in retrieving table. Incorrect resource ID"
		errorFlag = true
	} else {
		// Mark old one as Deleted
		//log.Printf("oldStudentSession Version is now %v", oldStudentSession)
		oldStudentSession.Deleted = time.Now().Format(time.RFC3339)
		oldStudentSession.Updated = time.Now().Format(time.RFC3339)
		newSession.Version = oldStudentSession.Version + 1
		Tiukudb.Table(tableToEdit).Save(&oldStudentSession)

		if err := Tiukudb.Table(tableToEdit).Create(&newSession).Error; err != nil {
			log.Printf("Error in starting Session %v \n", err)
			responseString = "Error in creating replacing Session"
			errorFlag = true
		} else {
			responseString = "Session data replaced."
		}
	}
	return responseString, errorFlag
}

// Update or replace existing course data
// W0rks ?
func UpdateCourse(updateCourse Course) (string, bool) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var errorFlag bool = false
	var responseString string
	// Get previous state or Archived
	prevArchiveState := GetCourseTableById(updateCourse.ID)
	// For now, don't allow to unArchive. So once true, always to true
	if prevArchiveState.Archived {
		updateCourse.Archived = true
	}
	var tempCourse Course
	if err := Tiukudb.Table(CourseTableToEdit).Model(&tempCourse).Where("id = ?", updateCourse.ID).Updates(Course{
		Degree:          updateCourse.Degree,
		CourseCode:      updateCourse.CourseCode,
		CourseName:      updateCourse.CourseName,
		CourseStartDate: updateCourse.CourseStartDate,
		CourseEndDate:   updateCourse.CourseEndDate,
		Archived:        updateCourse.Archived,
	}).Error; err != nil {
		log.Printf("Error: Problem updating course data. <database/update_database.go->UpdateCourse> %v \n", err)
		responseString = "Problem updating course data."
		errorFlag = true
	}
	// Need to skip this if it was already Archived.
	// Also skip if errorFlag is up
	if !errorFlag {
		if tempCourse.Archived && !prevArchiveState.Archived {
			log.Printf("Archiving course... %v", updateCourse.ID)
			err := ArchiveCourse(updateCourse.ID)
			if err {
				log.Printf("Error: Failed to Archive course. <database/update_database.go->UpdateCourse.")
				errorFlag = true
				responseString = "Error: Failed to Archive course."
			} else {
				responseString = "Course Archived successfully."
			}
		} else {
			log.Printf("Course already Archived.")
			responseString = "Course already Archived"
		}
	}
	return responseString, errorFlag
}

// Update Student User data
// W0rks
func UpdateStudentUser(newStudent StudentUser) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var errorFlag bool = false
	var tempStudent StudentUser

	if err := Tiukudb.Table(StudentsTableToEdit).Model(&tempStudent).Where("id = ?", newStudent.ID).Updates(StudentUser{
		StudentID:       newStudent.StudentID,
		StudentName:     newStudent.StudentName,
		StudentSegments: "",
		StudentEmail:    newStudent.StudentEmail,
		StudentClass:    newStudent.StudentClass,
	}).Error; err != nil {
		log.Printf("Error: Problem updating student user data. <database/update_database.go->UpdateStudentUser> %v \n", err)
		errorFlag = true
	}
	return errorFlag
}

// Update Faculty User data
// T35T
func UpdateFacultyUser(newFaculty FacultyUser) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var errorFlag bool = false
	var tempFaculty FacultyUser

	if err := Tiukudb.Table(FacultyTableToEdit).Model(&tempFaculty).Where("id = ?", newFaculty.ID).Updates(FacultyUser{
		FacultyID:    newFaculty.FacultyID,
		FacultyName:  newFaculty.FacultyName,
		FacultyEmail: newFaculty.FacultyEmail,
		Apartment:    newFaculty.Apartment,
		Active:       newFaculty.Active,
		Teacher:      newFaculty.Teacher,
		Admin:        newFaculty.Admin,
	}).Error; err != nil {
		log.Printf("Error: Problem updating faculty user data. <database/update_database.go->UpdateFacultyUser> %v \n", err)
		errorFlag = true
	}
	return errorFlag
}
