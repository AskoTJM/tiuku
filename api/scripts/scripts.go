package scripts

/*
// scripts.go
// Description: functions independent of database or http connections, mostly short snippets to make things easier.
*/
import (
	"fmt"
	"strconv"
)

// Turn uint64 to string
// W0rks
func Uint64ToString(newUint uint64) string {
	//var u uint32 = newUint
	//var s = strconv.FormatUint(uint64(u), 10)
	var s = strconv.FormatUint(uint64(newUint), 10)
	return s
}

// Turn uint to string
// W0rks
func UintToString(newUint uint) string {
	s := fmt.Sprint(newUint)
	return s
}

// Turn string to uint
// W0rks
func StringToUint(newString string) uint {
	s1, _ := strconv.ParseUint(newString, 10, 32)
	s2 := uint(s1)
	return s2
}

// Turn string to int
// W0rks
func StringToInt(newString string) int {
	s1, _ := strconv.Atoi(newString)
	return s1
}

// Turn string to uint32
// W0rks
func StringToUint32(newString string) uint32 {
	s1, _ := strconv.ParseUint(newString, 10, 32)
	s2 := uint32(s1)
	return s2
}

// Turn string to uint64
// W0rks
func StringToUint64(newString string) uint64 {
	s1, _ := strconv.ParseUint(newString, 10, 32)
	//s2 := uint32(s1)
	return s1
}

// Turn int to uint
// W0rks
func IntToUint(newInt int) uint {
	return uint(newInt)
}

// Turn uint to int
// W0rks
func UintToInt(newUint uint) int {
	return int(newUint)
}

// T0D0 not sure if needed, filtering on struct seems easier.
func CleanJSON(jsonToCheck string) string {
	var response string

	return response
}
