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
// W0rks
func CalculateOverAllTime(calcSession []database.StudentSegmentSession) (string, bool) {
	var errorFlag bool = false
	var returnTime time.Duration
	var durationTime time.Duration
	var returnString string
	//log.Printf("CalculateOverAllTime starting. %v \n", calcSession)
	for i := range calcSession {
		startTimeString := calcSession[i].StartTime
		log.Printf("startTime is %v \n", startTimeString)
		endTimeString := calcSession[i].EndTime
		log.Printf("endTime is %v \n", endTimeString)
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
// W0rks
func GetTimeDifference(from string, to string) (time.Duration, bool) {
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
// W0rks
func ParseTimeFormat(timeIn string) (time.Time, bool) {
	var errorFlag bool = false
	layout := "2006-01-02T15:04:05Z"
	response, err := time.Parse(layout, timeIn)
	if err != nil {
		log.Println("Error: Could not parse time. <stats/calculations.go->GetTimeDifferenceParseTimeFormat. ")
		errorFlag = true
	}
	log.Printf("Response is: %v", response)
	return response, errorFlag
}
