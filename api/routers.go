/*
 * Routing information for tiuku API
 */
package tiuku

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/AskoTJM/tiuku/api/faculty"
	students "github.com/AskoTJM/tiuku/api/students"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	// For testing purposes using Header to give commands
	gotcmd := HeaderTests(w, r)

	fmt.Fprintf(w, "Welcome to tiuku API %s \nDone %s", now, gotcmd)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	// Student v1 RAW Routes
	// STUDENT DELETEs
	Route{
		"DeleteCoursesCourseSegmentsSegment",
		strings.ToUpper("Delete"),
		"/students/v1/courses/{course}/segments/{segment}",
		students.DeleteCoursesCourseSegmentsSegment,
	},

	// DELETE {session} from {segment}
	// Should use soft delete
	Route{
		"DeleteSegmentsSegmentSessionsSession",
		strings.ToUpper("Delete"),
		"/students/v1/segments/{segment}/sessions/{session}",
		students.DeleteSegmentsSegmentSessionsSession,
	},

	// STUDENT GETs

	Route{
		"GetSessions",
		strings.ToUpper("Get"),
		"/students/v1/sessions",
		students.GetSessions,
	},

	Route{
		"GetSessionsLast",
		strings.ToUpper("Get"),
		"/students/v1/sessions/last",
		students.GetSessionsLast,
	},

	Route{
		"GetSchools",
		strings.ToUpper("Get"),
		"/students/v1/schools",
		students.GetSchools,
	},

	Route{
		"GetCampuses",
		strings.ToUpper("Get"),
		"/students/v1/campuses",
		students.GetCampuses,
	},

	Route{
		"GetApartments",
		strings.ToUpper("Get"),
		"/students/v1/apartments",
		students.GetApartments,
	},

	Route{
		"GetDegrees",
		strings.ToUpper("Get"),
		"/students/v1/degrees",
		students.GetDegrees,
	},

	Route{
		"GetCourses",
		strings.ToUpper("Get"),
		"/students/v1/courses",
		students.GetCourses,
	},

	Route{
		"GetCoursesCourse",
		strings.ToUpper("Get"),
		"/students/v1/courses/{course}",
		students.GetCoursesCourse,
	},

	Route{
		"GetCoursesCourseSegments",
		strings.ToUpper("Get"),
		"/students/v1/courses/{course}/segments",
		students.GetCoursesCourseSegments,
	},

	Route{
		"GetCoursesCourseSegmentsSegment",
		strings.ToUpper("Get"),
		"/students/v1/courses/{course}/segments/{segment}",
		students.GetCoursesCourseSegmentsSegment,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentCategories",
		strings.ToUpper("Get"),
		"/students/v1/courses/{course}/segments/{segment}/categories",
		students.GetCoursesCourseSegmentsSegmentCategories,
	},

	Route{
		"GetSegmentsSegmentSession",
		strings.ToUpper("Get"),
		"/students/v1/segments/{segment}/sessions",
		students.GetSegmentsSegmentSessions,
	},

	Route{
		"GetSegmentsSegmentSession",
		strings.ToUpper("Get"),
		"/students/v1/segments/{segment}/sessions/{session}",
		students.GetSegmentsSegmentSessionsSession,
	},

	Route{
		"GetSegmentsSegmentSettingsSetting",
		strings.ToUpper("Get"),
		"/students/v1/segments/{segment}/settings/{setting}",
		students.GetSegmentsSegmentSettingsSetting,
	},

	Route{
		"GetUserSegmentSettings",
		strings.ToUpper("Get"),
		"/students/v1/segment/{segment}/settings",
		students.GetUserSegmentsSettings,
	},

	Route{
		"GetUserSegments",
		strings.ToUpper("Get"),
		"/students/v1/segments",
		students.GetUserSegments,
	},

	Route{
		"GetUserSegmentsResourceID",
		strings.ToUpper("Get"),
		"/students/v1/segments/{segment}",
		students.GetUserSegmentsResourceID,
	},

	// PATCHs STUDENT

	Route{
		"PatchSegmentSegmentSettings",
		strings.ToUpper("Patch"),
		"/students/v1/segment/{segment}/settings",
		students.PatchSegmentSegmentSettings,
	},

	Route{
		"PatchSegmentsSegmentSessionsSession",
		strings.ToUpper("Patch"),
		"/students/v1/segments/{segment}/sessions/{session}",
		students.PatchSegmentsSegmentSessionsSession,
	},

	Route{
		"PatchSegmentsSegmentSettingsSetting",
		strings.ToUpper("Patch"),
		"/students/v1/segments/{segment}/settings/{setting}",
		students.PatchSegmentsSegmentSettingsSetting,
	},

	// STUDENT POSTs
	Route{
		"PostCoursesCourseSegmentsSegment",
		strings.ToUpper("Post"),
		"/students/v1/courses/{course}/segments/{segment}",
		students.PostCoursesCourseSegmentsSegment,
	},

	Route{
		"PostSegmentsSegmentSessions",
		strings.ToUpper("Post"),
		"/students/v1/segments/{segment}/sessions",
		students.PostSegmentsSegmentSessions,
	},

	Route{
		"PostUserSegments",
		strings.ToUpper("Post"),
		"/students/v1/segments",
		students.PostUserSegments,
	},
	// STUDENTS PUTs
	Route{
		"PutSegmentsSegmentSession",
		strings.ToUpper("Put"),
		"/students/v1/segments/{segment}/sessions/{session}",
		students.PutSegmentsSegmentSessionsSession,
	},

	// FACULTY v1 RAW Routes

	// FACULTY DELETEs
	Route{
		"DeleteStudentsStudent",
		strings.ToUpper("Delete"),
		"/faculty/v1/students/{student}",
		faculty.DeleteStudentsStudent,
	},
	// FACULTY GETs
	Route{
		"GetStudents",
		strings.ToUpper("Get"),
		"/faculty/v1/students",
		faculty.GetStudents,
	},

	Route{
		"GetStudentsStudent",
		strings.ToUpper("Get"),
		"/faculty/v1/students/{student}",
		faculty.GetStudentsStudent,
	},

	Route{
		"GetFaculty",
		strings.ToUpper("Get"),
		"/faculty/v1/faculty",
		faculty.GetFaculty,
	},

	Route{
		"GetSchools",
		strings.ToUpper("Get"),
		"/faculty/v1/schools",
		faculty.GetSchools,
	},

	Route{
		"GetCampuses",
		strings.ToUpper("Get"),
		"/faculty/v1/campuses",
		faculty.GetCampuses,
	},

	Route{
		"GetApartments",
		strings.ToUpper("Get"),
		"/faculty/v1/apartments",
		faculty.GetApartments,
	},

	Route{
		"GetDegrees",
		strings.ToUpper("Get"),
		"/faculty/v1/degrees",
		faculty.GetDegrees,
	},

	Route{
		"GetUserSegments",
		strings.ToUpper("Get"),
		"/faculty/v1/segments",
		faculty.GetUserSegments,
	},

	Route{
		"GetCourses",
		strings.ToUpper("Get"),
		"/faculty/v1/courses",
		faculty.GetCourses,
	},

	Route{
		"GetCoursesCourse",
		strings.ToUpper("Get"),
		"/faculty/v1/courses/{course}",
		faculty.GetCoursesCourse,
	},

	Route{
		"GetCoursesCourseSegments",
		strings.ToUpper("Get"),
		"/faculty/v1/courses/{course}/segments",
		faculty.GetCoursesCourseSegments,
	},

	Route{
		"GetCoursesCourseSegmentsSegment",
		strings.ToUpper("Get"),
		"/faculty/v1/courses/{course}/segments/{segment}",
		faculty.GetCoursesCourseSegmentsSegment,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentStudents",
		strings.ToUpper("Get"),
		"/faculty/v1/courses/{course}/segments/{segment}/students",
		faculty.GetCoursesCourseSegmentsSegmentStudents,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentSessions",
		strings.ToUpper("Get"),
		"/faculty/v1/courses/{course}/segments/{segment}/sessions",
		faculty.GetCoursesCourseSegmentsSegmentSessions,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentCategoriesCategory",
		strings.ToUpper("Get"),
		"/faculty/v1/courses/{course}/segments/{segment}/categories/{category}",
		faculty.GetCoursesCourseSegmentsSegmentCategoriesCategory,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentCategoriesCategorySettings",
		strings.ToUpper("Get"),
		"/faculty/v1/courses/{course}/segments/{segment}/categories/{category}/settings",
		faculty.GetCoursesCourseSegmentsSegmentCategoriesCategorySettings,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentCategoriesCategorySettingsSetting",
		strings.ToUpper("Get"),
		"/faculty/v1/courses/{course}/segments/{segment}/categories/{category}/settings/{setting}",
		faculty.GetCoursesCourseSegmentsSegmentCategoriesCategorySettingsSetting,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentSettings",
		strings.ToUpper("Get"),
		"/faculty/v1/courses/{course}/segments/{segment}/categories",
		faculty.GetCoursesCourseSegmentsSegmentSettings,
	},
	// FACULTY PATCHs

	Route{
		"PatchStudentsStudent",
		strings.ToUpper("Patch"),
		"/faculty/v1/students/{student}",
		faculty.PatchStudentsStudent,
	},

	Route{
		"PatchFacultyFaculty",
		strings.ToUpper("Patch"),
		"/faculty/v1/faculty/{faculty}",
		faculty.PatchFacultyFaculty,
	},

	Route{
		"PatchCoursesCourse",
		strings.ToUpper("Patch"),
		"/faculty/v1/courses/{course}",
		faculty.PatchCoursesCourse,
	},

	// FACULTY POSTs
	Route{
		"PostStudents",
		strings.ToUpper("Post"),
		"/faculty/v1/students",
		faculty.PostStudents,
	},

	Route{
		"PostFaculty",
		strings.ToUpper("Post"),
		"/faculty/v1/faculty",
		faculty.PostFaculty,
	},

	Route{
		"PostCourses",
		strings.ToUpper("Post"),
		"/faculty/v1/courses",
		faculty.PostCourses,
	},

	Route{
		"PostCoursesCourseSegments",
		strings.ToUpper("Post"),
		"/faculty/v1/courses/{course}/segments",
		faculty.PostCoursesCourseSegments,
	},

	Route{
		"PostCoursesCourseSegmentsSegmentCategories",
		strings.ToUpper("Post"),
		"/faculty/v1/courses/{course}/segments/{segment}/categories",
		faculty.PostCoursesCourseSegmentsSegmentCategories,
	},
}
