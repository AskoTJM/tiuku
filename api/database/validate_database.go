package database

import (
	"net/http"
)

// Validate that incoming Session has required data
// W0rks might need fine tuning.
func ValidateNewSessionStruct(newSession StudentSegmentSession) (bool, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseBool bool = true
	var responseString string
	var tempCategory SegmentCategory

	// Minimum needed to pass
	// Check for SegmentID
	if newSession.SegmentID == (StudentSegmentSession{}.SegmentID) {
		responseBool = false
		responseString = "Error: SegmentID required. \n"
	}
	// Check for Category
	var test bool
	var tempString string
	// First check if the Category and Segment match
	test, tempString, tempCategory = CheckIfSessionMatchesCategory(newSession)
	if !test {
		responseBool = false
		responseString = responseString + tempString
	} else {
		// If Category is mandatory to comment
		if tempCategory.MandatoryToComment {

			if newSession.Comment == (StudentSegmentSession{}.Comment) {
				if responseBool {
					responseBool = false
				}
				responseString = responseString + "Error: Comment required. \n"
			}
		}
		// Student name has to be visible.
		if tempCategory.MandatoryToTrack {
			if newSession.Privacy {
				if responseBool {
					responseBool = false
				}
				responseString = responseString + "Error: This category requires name to be visible. \n"
			}
		}
	}
	if responseString == "" {
		responseString = "Course Valid. \n"
	}
	return responseBool, responseString
}

// Check if New Category has required minimum of data
// W0rks
func ValidateNewCategory(newCategory SegmentCategory) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int = http.StatusOK
	var responseString string
	// Check if segment exists

	// Mandatory data checks
	if newCategory.SegmentID == (SegmentCategory{}.SegmentID) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = "Error: Missing SegmentID. \n"
	}
	if newCategory.MainCategory == (SegmentCategory{}.MainCategory) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Maincategory. \n"
	}
	if newCategory.SubCategory == "" {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Subcategory name. \n"
	}
	if responseString == "" {
		responseString = "New Category Valid. \n"
	}
	return responseCode, responseString
}

// Check if New Course has required minimum of data CourseCode, CourseName, Degree,
// W0rks
func ValidateNewCourse(newCourse Course) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int = http.StatusOK
	var responseString string
	// Check if segment exists

	// Data checks
	if newCourse.CourseCode == (Course{}.CourseCode) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = "Error: Missing CourseCode. \n"
	}
	if newCourse.Degree == (Course{}.Degree) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Degree. \n"
	}
	if newCourse.CourseName == "" {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Course name. \n"
	}
	if responseString == "" {
		responseString = "Course Valid. \n"
	}
	return responseCode, responseString
}

// Check if New Segment has required minimum of data
// W0rks
func ValidateNewSegment(newSegment Segment) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int = http.StatusOK
	var responseString string
	// Check if segment exists

	// Data checks
	if newSegment.CourseID == (Segment{}.CourseID) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = "Error: Missing CourseID. \n"
	}
	if newSegment.TeacherID == (Segment{}.TeacherID) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing TeachedID. \n"
	}
	if newSegment.Scope == (Segment{}.Scope) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Scope. \n"
	}
	if newSegment.ExpectedAttendance == (Segment{}.ExpectedAttendance) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing ExpectedAttendance. \n"
	}

	if newSegment.SegmentName == "" {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Segment name. \n"
	}
	if responseString == "" {
		responseString = "Segment Valid. \n"
	}
	return responseCode, responseString
}

// Check if New {StudentUser} has valid data
// T35T
func ValidateNewStudentUser(newStudent StudentUser) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int
	var responseString string
	// Check if already exists
	res := CheckIfUserExists(newStudent.StudentID)
	if res != 0 {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = "Error: StudentID already exists. \n"
	} else {
		if newStudent.StudentID == (StudentUser{}.StudentID) {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = "Error: Missing StudentID. \n"
		}
		if newStudent.StudentEmail == "" {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = responseString + "Error: Missing StudentEmail. \n"
		}

		if newStudent.StudentName == "" {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = responseString + "Error: Missing Student name. \n"
		}
		if responseString == "" {
			responseCode = http.StatusOK
			responseString = "Segment Valid."
		}
	}
	return responseCode, responseString
}

// Check if New {FacultyUser} has valid data
// T35T
func ValidateNewFacultyUser(newFaculty FacultyUser) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int
	var responseString string
	// Check if already exists
	res := CheckIfFacultyUserExists(newFaculty.FacultyID)
	if res != 0 {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = "Error: Faculty user already exists. \n"
	} else {
		if newFaculty.FacultyID == (FacultyUser{}.FacultyID) {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = "Error: Missing FacultyID. \n"
		}
		if newFaculty.FacultyEmail == "" {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = responseString + "Error: Missing Faculty Email. \n"
		}

		if newFaculty.FacultyName == "" {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = responseString + "Error: Missing Faculty name. \n"
		}
		if responseString == "" {
			responseCode = http.StatusOK
			responseString = "Segment Valid. \n"
		}
	}
	return responseCode, responseString
}
