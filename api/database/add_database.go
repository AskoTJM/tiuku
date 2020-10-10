package database

/*
// add_database.go
// Code to add data to database
*/

// Joining Student user to segment
// status: works
func AddStudentToSegment(user string, segmentId uint) string {
	if Tiukudb == nil {
		ConnectToDB()
	}

	segmentToJoin := GetSegmentDataById(segmentId)
	joiningStudent := GetStudentUser(user)

	if err := Tiukudb.Table(EnrollmentSegmentList).Create(&SchoolSegmentsSession{
		ID:                      0,
		SegmentID:               segmentToJoin.ID,
		AnonID:                  joiningStudent.AnonID,
		StudentSegmentsSessions: joiningStudent.AnonID + "_sessions",
		Privacy:                 "",
	}).Error; err != nil {
		response := "Error joining Segment. <database/update_database->UpdataParticipationToSegment>"
		return response
	}
	response := "Participated to Segment"
	AddSegmentToStudentsSegments(joiningStudent, segmentToJoin)
	return response
}

// Add joined Segment to Students Segment List
func AddSegmentToStudentsSegments(joiningStudent StudentUser, segmentToJoin Segment) string {
	if Tiukudb == nil {
		ConnectToDB()
	}

	tableToEdit := joiningStudent.AnonID + "_segments"
	if Tiukudb.HasTable(tableToEdit) {
		if err := Tiukudb.Table(tableToEdit).Create(&StudentSegment{
			ID:                     0,
			SegmentID:              segmentToJoin.ID,
			StudentSegmentSessions: joiningStudent.AnonID + "_sessions",
			Archived:               false,
		}).Error; err != nil {
			response := "Error joining Segment. <database/update_database->AddSegmentToStudentsSegment>"
			return response
		}
	} else {
		response := "Error joining. Student user doesn't have segments table. <database/update_database->AddSegmentToStudentsSegment>"
		return response
	}

	response := "Participated to Segment"
	return response
}
