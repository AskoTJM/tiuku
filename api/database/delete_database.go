package database

/*
// delete_database.go
// Removing stuff from database
//
*/

// Joining Student user to segment
// status: Works
func RemoveStudentFromSegment(joiningStudent StudentUser, segmentToJoin Segment) string {
	if Tiukudb == nil {
		ConnectToDB()
	}

	if err := Tiukudb.Table(enrollmentSegmentList).Where("anon_id = ? AND segment_id = ?", joiningStudent.AnonID, segmentToJoin.ID).Delete(&SchoolSegmentsSession{}).Error; err != nil {
		response := "Error removing Student from Segment. <database/update_database->RemoveStudentFromSegment>"
		return response
	}
	response := "Removed student from Segment"
	return response
}
