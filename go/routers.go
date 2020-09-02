/*
 * StudentAPI
 *
 * API for Students to acccess the Tiuku system.
 *
 * API version: 1.0
 * Contact: asko.mattila@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

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
	fmt.Fprintf(w, "Welcome to tiuku API.")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"DeleteCoursesCourseSegmentsSegment",
		strings.ToUpper("Delete"),
		"/courses/{course}/segments/{segment}",
		DeleteCoursesCourseSegmentsSegment,
	},

	Route{
		"DeleteSegmentsSegmentSession",
		strings.ToUpper("Delete"),
		"/segments/{segment}/{session}",
		DeleteSegmentsSegmentSession,
	},

	Route{
		"GetCourses",
		strings.ToUpper("Get"),
		"/courses",
		GetCourses,
	},

	Route{
		"GetCoursesCourseSegments",
		strings.ToUpper("Get"),
		"/courses/{course}/segments",
		GetCoursesCourseSegments,
	},

	Route{
		"GetCoursesCourseSegmentsSegment",
		strings.ToUpper("Get"),
		"/courses/{course}/segments/{segment}",
		GetCoursesCourseSegmentsSegment,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentCategories",
		strings.ToUpper("Get"),
		"/courses/{course}/segments/{segment}/categories",
		GetCoursesCourseSegmentsSegmentCategories,
	},

	Route{
		"GetSegmentsSegmentSession",
		strings.ToUpper("Get"),
		"/segments/{segment}/{session}",
		GetSegmentsSegmentSession,
	},

	Route{
		"GetSegmentsSegmentSettingsSetting",
		strings.ToUpper("Get"),
		"/segments/{segment}/settings/{setting}",
		GetSegmentsSegmentSettingsSetting,
	},

	Route{
		"GetUserSegmentSettings",
		strings.ToUpper("Get"),
		"/segment/{segment}/settings",
		GetUserSegmentSettings,
	},

	Route{
		"GetUserSegments",
		strings.ToUpper("Get"),
		"/segments",
		GetUserSegments,
	},

	Route{
		"GetUserSegmentsResourceID",
		strings.ToUpper("Get"),
		"/segments/{segment}",
		GetUserSegmentsResourceID,
	},

	Route{
		"PatchSegmentSegmentSettings",
		strings.ToUpper("Patch"),
		"/segment/{segment}/settings",
		PatchSegmentSegmentSettings,
	},

	Route{
		"PatchSegmentsSegmentSession",
		strings.ToUpper("Patch"),
		"/segments/{segment}/{session}",
		PatchSegmentsSegmentSession,
	},

	Route{
		"PatchSegmentsSegmentSettingsSetting",
		strings.ToUpper("Patch"),
		"/segments/{segment}/settings/{setting}",
		PatchSegmentsSegmentSettingsSetting,
	},

	Route{
		"PostCoursesCourseSegmentsSegment",
		strings.ToUpper("Post"),
		"/courses/{course}/segments/{segment}",
		PostCoursesCourseSegmentsSegment,
	},

	Route{
		"PostSegmentsSegment",
		strings.ToUpper("Post"),
		"/segments/{segment}",
		PostSegmentsSegment,
	},

	Route{
		"PostUserSegments",
		strings.ToUpper("Post"),
		"/segments",
		PostUserSegments,
	},

	Route{
		"PutSegmentsSegmentSession",
		strings.ToUpper("Put"),
		"/segments/{segment}/{session}",
		PutSegmentsSegmentSession,
	},
}
