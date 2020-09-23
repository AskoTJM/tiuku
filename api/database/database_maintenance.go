package database

// Contains scripts for maintenance of data

// Desc: Check if Student user exists
// Status: Works, maybe with slight changes could be used for all row counting?
func CheckIfUserExists(StudentID string) int64 {
	if db == nil {
		ConnectToDB()
	}

	tableToEdit := schoolShortName + "_StudentUsers"
	var tempStudent StudentUser

	result := db.Table(tableToEdit).Where("student_id = ?", StudentID).Find(&tempStudent)

	return result.RowsAffected
}

/*
func CheckIfAnonIdExists(anonid string) int {

}
*/
