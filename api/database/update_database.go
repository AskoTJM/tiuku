package database

/*
// update_database.go
// Description:
*/

// Update or replace existing course data
// status: work in progress
func UpdateCourse(updateCourse Course) {
	if Tiukudb == nil {
		ConnectToDB()
	}
}

func UpdateParticipationToSegment(joiningStudent StudentUser, segmentToJoin Segment) string {
	if Tiukudb == nil {
		ConnectToDB()
	}
	//var freshBlood SchoolSegmentsSession

	if err := Tiukudb.Table(participationTableToEdit).Create(&SchoolSegmentsSession{
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
	return response
}
