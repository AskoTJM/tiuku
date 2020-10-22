package stats

import (
	"log"
	"time"

	"github.com/AskoTJM/tiuku/api/database"
)

//
// calculations.go
// Description: calculations code.
//

var ectsHours uint = 27

// CountSegmentHours: How many hours allocated from ECTS
// T35T use uint or int?
func CountSegmentTargetHours(ects uint) uint {
	responseUint := ects * ectsHours
	return responseUint
}

// Count overall time
// W1P

func CalculateOverAllTime(calcSession []database.StudentSegmentSession) (string, bool) {
	var errorFlag bool = false
	var returnTime time.Duration
	var durationTime time.Duration
	var returnString string

	for i := range calcSession {
		startTimeString := calcSession[i].StartTime
		endTimeString := calcSession[i].EndTime
		if endTimeString == database.StringForEmpy {
			break
		} else {
			durationTime, _ = GetTimeDifference(startTimeString, endTimeString)
			returnTime = returnTime + durationTime
		}

	}
	returnString = returnTime.String()
	return returnString, errorFlag
}

// Time Difference returns time.Duration and errorFlag
// W1P
func GetTimeDifference(from string, to string) (time.Duration, bool) {
	//log.Print("Time now is %v \n", time.Now())
	var errorFlag bool = false
	var response time.Duration
	fromT, err := ParseTimeFormat(from)
	toT, err2 := ParseTimeFormat(to)
	if err || err2 {
		errorFlag = true
	} else {
		if fromT.Before(toT) {
			response = toT.Sub(fromT)
			log.Println(response)
		} else {
			errorFlag = true
		}
	}
	return response, errorFlag
}

// Parse String to time.Time, returns time.Time and errorFlag
// W1P
func ParseTimeFormat(timeIn string) (time.Time, bool) {
	var errorFlag bool = false
	layout := "2006-01-02T15:04:05.000Z"
	response, err := time.Parse(layout, timeIn)
	if err != nil {
		log.Println("Error: Could not parse time. <scripts/calculations.go->GetTimeDifferenceParseTimeFormat. ")
		errorFlag = true
	}
	return response, errorFlag
}
