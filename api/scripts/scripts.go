package scripts

/*
// scripts
// Description: functions independent of database or http connections, mostly to make somethings easier.
*/
import (
	"fmt"
	"net/http"
	"strconv"
)

// desc: Turn uint64 to string
func Uint64ToString(newUint uint64) string {
	//var u uint32 = newUint
	//var s = strconv.FormatUint(uint64(u), 10)
	var s = strconv.FormatUint(uint64(newUint), 10)
	return s
}

// desc: Turn uint to string
func UintToString(newUint uint) string {
	s := fmt.Sprint(newUint)
	return s
}

// desc: Turn string to uint
func StringToUint(newString string) uint {
	s1, _ := strconv.ParseUint(newString, 10, 32)
	s2 := uint(s1)
	return s2
}

// desc: Turn string to uint32
func StringToUint32(newString string) uint32 {
	s1, _ := strconv.ParseUint(newString, 10, 32)
	s2 := uint32(s1)
	return s2
}

// desc: Turn string to uint64
func StringToUint64(newString string) uint64 {
	s1, _ := strconv.ParseUint(newString, 10, 32)
	//s2 := uint32(s1)
	return s1
}

func TestResponse(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
