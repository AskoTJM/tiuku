package database

/*
// get_database.go
// Description: Code for retrieving data from database
*/
import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Desc: Get Students data
// Status: Works, but needs more. Return value and obfuscing of AnonID if used outside
func GetStudentUser(StudentID string) StudentUser {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var tempStudent StudentUser

	result := Tiukudb.Table(StudentsTableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)
	if result == nil {
		log.Println(result)
	}
	if debugMode {
		log.Printf("tempStudent has value of: %s", tempStudent.AnonID)
	}
	return tempStudent
}

// Get segments of student user
// status:
func GetUserSegments(student StudentUser, params string) []StudentSegment {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var tempSegment []StudentSegment
	myAnonID := student.AnonID
	tableToEdit := myAnonID + "_segments"
	var result *gorm.DB //:= Tiukudb.Table(tableToEdit)
	if params == "no" {
		result = Tiukudb.Table(tableToEdit).Where("archived = ?", false).Find(&tempSegment)
		if result != nil {
			log.Println(result.Error)
		}
	} else if params == "yes" {
		result = Tiukudb.Table(tableToEdit).Find(&tempSegment)
		if result != nil {
			log.Println(result.Error)
		}
	} else if params == "only" {
		result = Tiukudb.Table(tableToEdit).Where("archived = ?", true).Find(&tempSegment)
		if result != nil {
			log.Println(result.Error)
		}
	} else {
		log.Printf("Error: Invalid parameters.")
	}

	returnSegments := make([]StudentSegment, 0)
	result2, _ := result.Rows()

	var tempSegments2 StudentSegment
	for result2.Next() {

		if err3 := result.ScanRows(result2, &tempSegments2); err3 != nil {
			log.Println(err3)
		}
		returnSegments = append(returnSegments, tempSegments2)
	}

	return returnSegments
}

// Get Courses, default only active. with "archived=yes" all courses, "archived=only" to get only archived ones.
// Status: Works
func GetCourses(choice string) []Course {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var tempCourses []Course
	var result *gorm.DB

	if choice == "no" {
		result = Tiukudb.Table(CourseTableToEdit).Where("archived = ?", false).Find(&tempCourses)
		if result != nil {
			log.Println(result.Error)
		}
	} else if choice == "yes" {
		result = Tiukudb.Table(CourseTableToEdit).Find(&tempCourses)
		if result != nil {
			log.Println(result.Error)
		}
	} else if choice == "only" {
		result = Tiukudb.Table(CourseTableToEdit).Where("archived = ?", true).Find(&tempCourses)
		if result != nil {
			log.Println(result)
		}
	} else {
		log.Println("Error: Invalid parameters.")
	}

	returnCourses := make([]Course, 0)
	result2, _ := result.Rows()

	var tempCourse2 Course
	for result2.Next() {

		if err3 := result.ScanRows(result2, &tempCourse2); err3 != nil {
			log.Println(err3)
		}
		returnCourses = append(returnCourses, tempCourse2)
	}

	return returnCourses
}

// FacultyUserSpesifics

// Desc: Get Faculty User
// Status: Works
func GetFacultyUser(FacultyID string) FacultyUser {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var tempFaculty FacultyUser

	result := Tiukudb.Table(FacultyTableToEdit).Where("faculty_id = ?", FacultyID).First(&tempFaculty)
	if result == nil {
		log.Println(result)
	}

	return tempFaculty
}

// Get segments of faculty user, active and with parameters, archived=yes and archived=only
// status:
// comment: New version as added Archived bool, makes search simpler, doesn't require reading Courses tables.
func GetFacultyUserSegments(r *http.Request) []Segment {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var tempSegment []Segment
	var result *gorm.DB //db.Table(segmentTableToEdit)
	user := r.Header.Get("X-User")
	// Get teachers ID number
	teacher := GetFacultyUser(user)

	paramTest := r.URL.Query()
	filter, params := paramTest["archived"]
	if !params || len(filter) == 0 {
		result = Tiukudb.Table(SegmentTableToEdit).Where("teacher_id = ?", teacher.ID).Where("archived = ?", false).Find(&tempSegment)
		if result != nil {
			log.Println(result.Error)
		}
	} else if paramTest.Get("archived") == "yes" {
		result = Tiukudb.Table(SegmentTableToEdit).Where("teacher_id = ?", teacher.ID).Find(&tempSegment)
		if result != nil {
			log.Println(result.Error)
		}
	} else if paramTest.Get("archived") == "only" {
		result = Tiukudb.Table(SegmentTableToEdit).Where("teacher_id = ?", teacher.ID).Where("archived = ?", true).Find(&tempSegment)
		if result != nil {
			log.Println(result.Error)
		}
	} else {
		fmt.Println("Error: Invalid parameters.")
	}

	returnSegments := make([]Segment, 0)
	result2, _ := result.Rows()

	var tempSegments2 Segment
	for result2.Next() {
		//Read row to tempSegments2
		if err3 := result.ScanRows(result2, &tempSegments2); err3 != nil {
			log.Println(err3)
		}
		returnSegments = append(returnSegments, tempSegments2)
	}

	return returnSegments
}

// Find course by courseCode
// Status: Unclear
func GetCourseTableByCode(courseCode string) Course {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var tempCourse Course
	Tiukudb.Table(CourseTableToEdit).Where("course_code = ?", courseCode).Find(&tempCourse).Row()
	return tempCourse
}

// Find course by id
// Status: Unclear
func GetCourseTableById(id uint) Course {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var tempCourse Course
	Tiukudb.Table(CourseTableToEdit).Where("id = ?", id).Find(&tempCourse).Row()
	return tempCourse
}

// Find segment by id
// status:
func GetSegmentDataById(id uint) Segment {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var tempSegment Segment
	Tiukudb.Table(SegmentTableToEdit).Where("id = ?", id).Find(&tempSegment).Row()
	return tempSegment
}

func GetSegmentTableByCourseId(courseID uint) []Segment {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var tempSegment []Segment
	//tempSegment := make([]Segment, 0)
	result := Tiukudb.Table(SegmentTableToEdit).Where("course_id = ?", courseID).Find(&tempSegment)
	if result != nil {
		log.Println(result)
	}
	returnSegment := make([]Segment, 0)
	result2, _ := result.Rows()
	var tempCourse2 Segment
	for result2.Next() {
		if err3 := result.ScanRows(result2, &tempCourse2); err3 != nil {
			log.Println(err3)
		}
		returnSegment = append(returnSegment, tempCourse2)
	}
	return returnSegment
}

// Get Category by ID number
// comment:
func GetCategoryById(categoryID uint) SegmentCategory {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var tempCategory SegmentCategory
	result := Tiukudb.Table(CategoriesTableToEdit).Where("id = ?", categoryID).Find(&tempCategory)
	if result != nil {
		log.Println(result)
	}
	return tempCategory
}

// Find ALL categories belonging to segment, true include 0 defaults / false only segments own categories.
// comment: If using categories table for all segments
func GetCategoriesBySegmentId(segmentID uint, includeZero bool, includeInActive bool) []SegmentCategory {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var tempSegment []SegmentCategory
	var result *gorm.DB

	if includeInActive {
		if includeZero {
			result = Tiukudb.Table(CategoriesTableToEdit).Where("segment_id = ?", segmentID).Or("segment_id = ?", 0).Find(&tempSegment)
			if debugMode {
				log.Printf("IncludeZero and IncludeInActive")
			}
		} else {
			result = Tiukudb.Table(CategoriesTableToEdit).Where("segment_id = ?", segmentID).Find(&tempSegment)
			if debugMode {
				log.Printf("No IncludeZero and IncludeInActive")
			}
		}
		//if result != nil {
		//	log.Println(result)
		//}
	} else {
		if includeZero {
			// With Segment_Id 0 value should always be active as it is mandatory category for all segments i.e. no need to check it
			result = Tiukudb.Table(CategoriesTableToEdit).Where("active = ? AND segment_id = ?", true, segmentID).Or("segment_id = ?", 0).Find(&tempSegment)

			if debugMode {
				log.Printf("IncludeZero and No IncludeInActive")
			}
		} else {
			result = Tiukudb.Table(CategoriesTableToEdit).Where("segment_id = ?", segmentID).Find(&tempSegment)
			if debugMode {
				log.Printf("No IncludeZero and IncludeInActive")
			}
		}
		//if result != nil {
		//	log.Println(result)
		//}
	}

	returnSegment := make([]SegmentCategory, 0)
	result2, _ := result.Rows()
	var tempSegment2 SegmentCategory
	for result2.Next() {
		if err3 := result.ScanRows(result2, &tempSegment2); err3 != nil {
			log.Println(err3)
		}
		returnSegment = append(returnSegment, tempSegment2)
	}
	return returnSegment
}

// Get degree with ID number, 0 returns all degrees
func GetDegree(degreeID uint) []Degree {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var result *gorm.DB
	var tempDegree []Degree
	if degreeID == 0 {
		result = Tiukudb.Table(DegreeTableToEdit).Find(&tempDegree)
	} else {
		result = Tiukudb.Table(DegreeTableToEdit).Where("id = ?", degreeID).Find(&tempDegree)
	}
	returnDegree := make([]Degree, 0)
	result2, _ := result.Rows()

	var tempDegree2 Degree
	for result2.Next() {

		if err3 := result.ScanRows(result2, &tempDegree2); err3 != nil {
			log.Println(err3)
		}
		returnDegree = append(returnDegree, tempDegree2)
	}

	return returnDegree
}

// Get Apartment info with ID number, 0 returns all apartments
func GetApartment(apartmentID uint) []Apartment {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var result *gorm.DB
	var tempApartment []Apartment
	if apartmentID == 0 {
		result = Tiukudb.Table(ApartmentTableToEdit).Find(&tempApartment)
	} else {
		result = Tiukudb.Table(ApartmentTableToEdit).Where("id = ?", apartmentID).Find(&tempApartment)
	}
	returnApartment := make([]Apartment, 0)
	result2, _ := result.Rows()

	var tempApartment2 Apartment
	for result2.Next() {

		if err3 := result.ScanRows(result2, &tempApartment2); err3 != nil {
			log.Println(err3)
		}
		returnApartment = append(returnApartment, tempApartment2)
	}

	return returnApartment
}

// Get Campus info with ID number, 0 returns all campuses
func GetCampus(campusID uint) []Campus {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var result *gorm.DB
	var tempCampus []Campus
	if campusID == 0 {
		result = Tiukudb.Table(CampusTableToEdit).Find(&tempCampus)
	} else {
		result = Tiukudb.Table(CampusTableToEdit).Where("id = ?", campusID).Find(&tempCampus)
	}
	returnCampus := make([]Campus, 0)
	result2, _ := result.Rows()

	var tempCampus2 Campus
	for result2.Next() {

		if err3 := result.ScanRows(result2, &tempCampus2); err3 != nil {
			log.Println(err3)
		}
		returnCampus = append(returnCampus, tempCampus2)
	}

	return returnCampus
}

// Get School info with ID number, 0 returns all schools
func GetSchool(schoolID uint) []School {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var result *gorm.DB
	var tempSchool []School
	if schoolID == 0 {
		result = Tiukudb.Table(SchoolsTableToEdit).Find(&tempSchool)
	} else {
		result = Tiukudb.Table(SchoolsTableToEdit).Where("id = ?", schoolID).Find(&tempSchool)
	}
	returnSchool := make([]School, 0)
	result2, _ := result.Rows()

	var tempSchool2 School
	for result2.Next() {

		if err3 := result.ScanRows(result2, &tempSchool2); err3 != nil {
			log.Println(err3)
		}
		returnSchool = append(returnSchool, tempSchool2)
	}

	return returnSchool
}

// Desc: GetAnonId with StudentID
// HOX! AnonID SHOULD NOT LEAVE OUTSIDE OF THE API
// Status: Done but REMOVED, better to use GetStudentUser
/*
func GetAnonId(StudentID string) string {
	if db == nil {
		ConnectToDB()
	}

	var tempStudent StudentUser

	result := db.Table(studentsTableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)
	if result == nil {
		log.Println(result)
	}
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	tempJSON := gjson.Get(s, "Value.AnonID")
	return tempJSON.String()
}
*/

// desc: Get name of the student with StudentID
// Status: Removed,not in use, just use GetStudentUser
/*
func GetStudentName(StudentID string) string {
	if db == nil {
		ConnectToDB()
	}

	var tempStudent StudentUser

	result := db.Table(courseTableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)
	if result == nil {
		log.Println(result)
	}
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	tempJSON := gjson.Get(s, "Value.StudentName")
	return tempJSON.String()
}
*/

// desc: Get segments of faculty user, active and with parameters, archived=yes and archived=only
// status: Works, but replaced by new version as Archived added possibility to archive invidual segments
// left if we happen to need to go back in this.
// HOX! Weird issues with filtering archived/non-archived results.
/*
func GetFacultyUserSegments(r *http.Request) []Segment {
	if db == nil {
		ConnectToDB()
	}
	var tempSegment []Segment

	user := r.Header.Get("X-User")
	// Get teachers ID number
	teacher := GetFacultyUser(user)
	result := db.Table(segmentTableToEdit).Where("teacher_id = ?", teacher.ID).Find(&tempSegment)
	paramTest := r.URL.Query()
	filter, params := paramTest["archived"]

	returnSegments := make([]Segment, 0)
	result2, _ := result.Rows()

	var tempSegments2 Segment
	for result2.Next() {

		//Read row to tempSegments2
		if err3 := result.ScanRows(result2, &tempSegments2); err3 != nil {
			log.Println(err3)
		}
		byID := tempSegments2.CourseID
		var tempCourse Course
		if err4 := db.Table(courseTableToEdit).First(&tempCourse, byID); err4 != nil {
			log.Println(err4)
		}
		log.Println("tempSegments2 value: ")
		log.Println(byID)
		if !params || len(filter) == 0 {
			if !tempCourse.Archived {
				returnSegments = append(returnSegments, tempSegments2)
			}
			if result != nil {
				log.Println(result.Error)
			}
		} else if paramTest.Get("archived") == "yes" {
			returnSegments = append(returnSegments, tempSegments2)
			if result != nil {
				log.Println(result.Error)
			}
		} else if paramTest.Get("archived") == "only" {
			if tempCourse.Archived {
				returnSegments = append(returnSegments, tempSegments2)
			}
			if result != nil {
				log.Println(result)
			}
		} else {
			fmt.Println("Error: Invalid parameters.")
		}
	}
	return returnSegments
}*/
