package database

import (
	"log"
	"net/http"
)

// Fix archive status of courses segments and categories
// Should work by working through courses and checking that their segments and categories have same archive status
// T0D0
func CheckArchiveConflicts() int {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var numberOfFixes int

	return numberOfFixes
}

// Test if required tables exist, should include all the necessary table or fail
// W0rks
func CheckIfRequiredTablesExist() bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	if !Tiukudb.HasTable(CourseTableToEdit) {
		return false
	}
	if !Tiukudb.HasTable(SegmentTableToEdit) {
		return false
	}
	if !Tiukudb.HasTable(StudentsTableToEdit) {
		return false
	}
	if !Tiukudb.HasTable(FacultyTableToEdit) {
		return false
	}
	if !Tiukudb.HasTable(CategoriesTableToEdit) {
		return false
	}
	return true
}

// Check if content of request is JSON
// W0rks
func CheckJSONContent(w http.ResponseWriter, r *http.Request) string {
	if r.Header.Get("Content-Type") == "" {
		// Removed to be able to use this code with empty values

		//w.WriteHeader(http.StatusNoContent)
		response := "EMPTY"
		return response
	} else {
		//rbody, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		rbody := r.Header.Get("Content-Type")
		// Check if content type is correct one.
		if rbody != "application/json" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotAcceptable)
			response := "TYPE_ERROR"
			return response
		}

	}
	return "PASS"
}

// Check if Category matches Segment, returns True if match False if not
// W0rks
func CheckIfCategoryMatchSegment(testCategory uint, testSegment uint) (bool, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var errorFlag bool
	var responseString string
	var tempCategory SegmentCategory
	if testCategory == 0 {
		errorFlag = false
		responseString = "Category not provided or incorrect one."
	} else if testCategory > 3 {
		res := Tiukudb.Table(CategoriesTableToEdit).Where("id = ?", testCategory).Where("segment_id = ?", testSegment).Find(&tempCategory).RowsAffected
		if res == 0 {
			errorFlag = false
			responseString = "Error. Incorrect Category for Segment."
		}
		if res == 1 {
			errorFlag = true
			responseString = "Category matches the Segment."
		}
	} else {
		errorFlag = true
		responseString = "Category is default one."
	}
	return errorFlag, responseString
}

// Check if Sessions Category matches the Segment, returns True if match False if not
// W0rks ,might need additional work
func CheckIfSessionMatchesCategory(tempSession StudentSegmentSession) (bool, string, SegmentCategory) {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var errorFlag bool
	var responseString string
	var tempCategory SegmentCategory
	//log.Printf("Category is %v and Segment is %v", tempSession.Category, tempSession.SegmentID)
	if tempSession.Category == 0 {
		errorFlag = false
		responseString = "Category not provided or incorrect one."
	} else if tempSession.Category > 3 {
		//log.Printf("Category over 4")
		result := Tiukudb.Table(CategoriesTableToEdit).Where("id = ?", tempSession.Category).Where("segment_id = ?", tempSession.SegmentID).Where("active = ?", true).Find(&tempCategory)

		res := result.RowsAffected
		//log.Printf("Category over 4 and %v rows matching found", res)
		if res == 0 {
			errorFlag = false
			responseString = "Error. Incorrect Category for Segment."
		}
		if res == 1 {
			errorFlag = true
			responseString = "Category matches the Segment."
		}
	} else {
		errorFlag = true
		responseString = "Category is default one."
	}
	return errorFlag, responseString, tempCategory
}

// Check if ResourceID exists in users table.Input: user and resource_id to check
// T35T
func CheckIfResourceIDExistsInSessionTable(user string, ruid uint) (uint, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	//var errorFlag bool
	var responseStatusCode uint
	var responseString string
	var tempStudent StudentUser
	var tempSession StudentSegmentSession
	if err := Tiukudb.Table(StudentsTableToEdit).Where("student_id = ?", user).Find(&tempStudent).Error; err != nil {
		log.Printf("Error retrieving user data. %v ", err)
		//errorFlag = false
		responseStatusCode = http.StatusInternalServerError
		responseString = "Error retrieving user data."
	} else {
		matches := Tiukudb.Table(tempStudent.AnonID+"_sessions").Where("resource_id = ?", ruid).Find(&tempSession).RowsAffected
		if matches == 0 {
			log.Printf("Error retrieving session data. Incorrect resource_id. ")
			//errorFlag = false
			responseStatusCode = http.StatusBadRequest
			responseString = "Error retrieving session data. Incorrect resource_id."
		} else {
			//errorFlag = true
			responseStatusCode = http.StatusAccepted
			responseString = "ResourceID exists."
		}
	}

	return responseStatusCode, responseString
}

// Check if Student user exists student users table, returns StatusOk if
// W1P
func CheckIfUserExists(StudentID string) (string, int) {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var responseCode int
	var responseString string
	if StudentID == "" {
		responseCode = http.StatusBadRequest
		responseString = "Error: Student User doesn't exist."
	} else {
		//tableToEdit := schoolShortName + "_StudentUsers"
		var tempStudent StudentUser
		//log.Printf("StudentId is %v", StudentID)
		result := Tiukudb.Table(StudentsTableToEdit).Where("student_id = ?", StudentID).Find(&tempStudent)
		if result.RowsAffected == 0 {
			responseCode = http.StatusBadRequest
			responseString = "Incorrect StudentID"
		} else if result.RowsAffected > 1 {
			log.Printf("Error, multiple users with same StudentID found in <database/check_database.go->CheckIfUserExists>")
			responseCode = http.StatusInternalServerError
			responseString = "Problems with the server."
		} else if result.RowsAffected == 1 {
			responseCode = http.StatusOK
			responseString = "Student ID found"
		} else {
			log.Printf("Error: In <database/check_database.go->CheckIfUserExists>")
		}
	}
	return responseString, responseCode
}

// Check if Student user doesn't exists in student users table, returns StatusOk if
// W0rks
func CheckIfUserIdAvailable(StudentID string) (string, int) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int
	var responseString string
	//tableToEdit := schoolShortName + "_StudentUsers"
	var tempStudent StudentUser

	result := Tiukudb.Table(StudentsTableToEdit).Where("student_id = ?", StudentID).Find(&tempStudent)
	if result.RowsAffected == 0 {
		responseCode = http.StatusOK
		responseString = "StudentID available."
	} else if result.RowsAffected < 1 {
		log.Printf("Error, multiple users with same StudentID found in <database/check_database.go->CheckIfUserExists>")
		responseCode = http.StatusInternalServerError
		responseString = "Problems with the server."
	} else {
		responseCode = http.StatusBadRequest
		responseString = "Student ID already exists."
	}

	return responseString, responseCode
}

// Check if Faculty User is in DB
// W0rks
func CheckIfFacultyUserExists(FacultyID string) (string, int) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int
	var responseString string
	//tableToEdit := schoolShortName + "_StudentUsers"
	var tempFaculty FacultyUser

	result := Tiukudb.Table(FacultyTableToEdit).Where("faculty_id = ?", FacultyID).Find(&tempFaculty)
	if result.RowsAffected == 0 {
		responseCode = http.StatusBadRequest
		responseString = "Incorrect FacultyID"
	} else if result.RowsAffected < 1 {
		log.Printf("Error, multiple users with same FacultyID found in <database/check_database.go->CheckIfFacultyUserExists>")
		responseCode = http.StatusInternalServerError
		responseString = "Problems with the server."
	} else {
		responseCode = http.StatusOK
		responseString = "FacultyID found"
	}

	return responseString, responseCode
}

// Check if Faculty ID is available
// W0rks
func CheckIfFacultyIdAvailable(FacultyID string) (string, int) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int
	var responseString string
	var tempFaculty FacultyUser

	result := Tiukudb.Table(FacultyTableToEdit).Where("faculty_id = ?", FacultyID).Find(&tempFaculty)
	if result.RowsAffected == 0 {
		responseCode = http.StatusOK
		responseString = "FacultyID available."
	} else if result.RowsAffected < 1 {
		log.Printf("Error, multiple users with same FacultyID found in <database/check_database.go->CheckIfFacultyIdAvailable>")
		responseCode = http.StatusInternalServerError
		responseString = "Problems with the server."
	} else {
		responseCode = http.StatusBadRequest
		responseString = "Faculty ID already exists."
	}

	return responseString, responseCode
}
