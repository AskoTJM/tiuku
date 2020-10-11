package scripts

/*
// scripts.go
// Description: functions independent of database or http connections, mostly short snippets to make things easier.
*/
import (
	"fmt"
	"net/http"
	"strconv"
)

// Turn uint64 to string
func Uint64ToString(newUint uint64) string {
	//var u uint32 = newUint
	//var s = strconv.FormatUint(uint64(u), 10)
	var s = strconv.FormatUint(uint64(newUint), 10)
	return s
}

// Turn uint to string
func UintToString(newUint uint) string {
	s := fmt.Sprint(newUint)
	return s
}

// Turn string to uint
func StringToUint(newString string) uint {
	s1, _ := strconv.ParseUint(newString, 10, 32)
	s2 := uint(s1)
	return s2
}

// Turn string to uint32
func StringToUint32(newString string) uint32 {
	s1, _ := strconv.ParseUint(newString, 10, 32)
	s2 := uint32(s1)
	return s2
}

// Turn string to uint64
func StringToUint64(newString string) uint64 {
	s1, _ := strconv.ParseUint(newString, 10, 32)
	//s2 := uint32(s1)
	return s1
}

// Turn int to uint
func IntToUint(newInt int) uint {
	return uint(newInt)
}

// Turn uint to int
func UintToInt(newUint uint) int {
	return int(newUint)
}

func TestResponse(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// Check if content of request is JSON
func CheckJSONContent(w http.ResponseWriter, r *http.Request) string {
	if r.Header.Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNoContent)
		response := "Problems creating new Course, no body in request information. <database/create_database->CreateCourse> Error: No body information available."
		return response
	} else {
		//rbody, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		rbody := r.Header.Get("Content-Type")
		// Check if content type is correct one.
		if rbody != "application/json" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotAcceptable)
			response := "Error: Content-Type is not application/json."
			return response
		}

	}
	return "PASS"
}

func CleanJSON(jsonToCheck string) string {
	var response string

	return response
}
