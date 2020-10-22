package faculty

/*
// post_api_faculty_v1.go
// Description: POST request functions for Faculty users
*/
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// Create new Course in course table
// W0rks
func PostCourses(w http.ResponseWriter, r *http.Request) {

	var response string

	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else if !resF.Admin {
		response = "Insuffient rights."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resJsonString, resJsonCode := database.CheckJSONContent(w, r)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if resJsonCode != http.StatusOK {
			w.WriteHeader(resJsonCode)
			response = resJsonString
		} else {
			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			var newCourse database.Course
			err := dec.Decode(&newCourse)
			if err != nil {
				log.Println(err)
			}
			resCode, resString := database.ValidateNewCourse(newCourse)
			if resCode != http.StatusOK {
				log.Printf("Response from Validitation test %v", resString)
				w.WriteHeader(resCode)
				response = resString
			} else {
				response = database.CreateCourse(newCourse, database.CourseTableToEdit)
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusCreated)
			}
		}
	}
	fmt.Fprintf(w, "%s", response)
}

// New Student User
// W0rks
func PostStudents(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else if !resF.Admin {
		response = "Insuffient rights."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resJsonString, resJsonCode := database.CheckJSONContent(w, r)
		if resJsonCode != http.StatusOK {
			w.WriteHeader(resJsonCode)
			response = resJsonString
		} else {
			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			var newStudent database.StudentUser
			err := dec.Decode(&newStudent)
			if err != nil {
				log.Println(err)
			}
			resCode, resString := database.ValidateNewStudentUser(newStudent)
			if resCode != http.StatusOK {
				log.Printf("Response from Validitation test %v", resString)
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(resCode)
				response = resString
			} else {
				resCode2, resString2 := database.CreateNewStudentUser(newStudent)
				if resCode2 != http.StatusOK {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(resCode)
					//response = resString2
				} else {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(http.StatusCreated)
					//response
				}
				response = resString2
			}
		}
	}
	fmt.Fprintf(w, "%s", response)
}

// New Faculty User
// W0rks
func PostFaculty(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else if !resF.Admin {
		response = "Insuffient rights."
		w.WriteHeader(http.StatusBadRequest)
	} else {

		resJsonString, resJsonCode := database.CheckJSONContent(w, r)
		if resJsonCode != http.StatusOK {
			w.WriteHeader(resJsonCode)
			response = resJsonString
		} else {
			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			var newFaculty database.FacultyUser
			err := dec.Decode(&newFaculty)
			if err != nil {
				log.Println(err)
			}
			resCode, resString := database.ValidateNewFacultyUser(newFaculty)
			if resCode != http.StatusOK {
				log.Printf("Response from Validitation test %v", resString)
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(resCode)
				response = resString
			} else {
				resCode2, resString2 := database.CreateNewFacultyUser(newFaculty)
				if resCode2 != http.StatusOK {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(resCode)
				} else {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(http.StatusCreated)
				}
				response = resString2
			}
		}
	}
	fmt.Fprintf(w, "%s", response)
}

// New segment for the course
// W0rks
func PostCoursesCourseSegments(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else if !resF.Admin {
		response = "Insuffient rights."
		w.WriteHeader(http.StatusBadRequest)
	} else {

		resJsonString, resJsonCode := database.CheckJSONContent(w, r)
		if resJsonCode != http.StatusOK {
			w.WriteHeader(resJsonCode)
			response = resJsonString
		} else {
			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			var newSegment database.Segment
			err := dec.Decode(&newSegment)
			if err != nil {
				log.Println(err)
			}
			vars := mux.Vars(r)
			courseCode := vars["course"]
			newSegment.CourseID = scripts.StringToUint(courseCode)
			resCode, resString := database.ValidateNewSegment(newSegment)
			if resCode != http.StatusOK {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(resCode)
				response = resString
			} else {
				response = database.CreateSegment(newSegment, database.SegmentTableToEdit)

				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusCreated)
			}
		}
	}
	fmt.Fprintf(w, "%s", response)
}

// Add New category for segment
// W0rks
func PostCoursesCourseSegmentsSegmentCategories(w http.ResponseWriter, r *http.Request) {

	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resJsonString, resJsonCode := database.CheckJSONContent(w, r)
		if resJsonCode != http.StatusOK {
			w.WriteHeader(resJsonCode)
			response = resJsonString
		} else {
			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			var newCategory database.SegmentCategory
			err := dec.Decode(&newCategory)
			if err != nil {
				log.Println(err)
			}
			vars := mux.Vars(r)
			segmentCode := vars["segment"]
			newCategory.SegmentID = scripts.StringToUint(segmentCode)
			resCode, resString := database.ValidateNewCategory(newCategory)
			if resCode != http.StatusOK {
				w.WriteHeader(resCode)
				response = resString
			} else {
				result := database.CreateCategory(newCategory, database.CategoriesTableToEdit)
				if result {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(http.StatusCreated)
					response = response + " Category created for Segment"
				} else {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(http.StatusInternalServerError)
					response = response + " Could not create Category for Segment"
				}
			}
		}
	}
	fmt.Fprintf(w, "%s", response)
}

// Add New School
// T0D0

// Add New Campus
// T0D0

// Add New Apartment
// T0D0

// Add New Degree
// T0D0
