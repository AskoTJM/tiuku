package database

import "log"

// desc: Find course by courseCode
// Status: Unclear
func FindCourseTableByCode(courseCode string) Course {
	if db == nil {
		ConnectToDB()
	}
	var tempCourse Course
	db.Table(courseTableToEdit).Where("course_code = ?", courseCode).Find(&tempCourse).Row()
	return tempCourse
}

// desc: Find course by id
// Status: Unclear
func FindCourseTableById(id string) Course {
	if db == nil {
		ConnectToDB()
	}
	var tempCourse Course
	db.Table(courseTableToEdit).Where("id = ?", id).Find(&tempCourse).Row()
	return tempCourse
}

// desc: Find segment by id
// status:
func FindSegmentDataById(id string) Segment {
	if db == nil {
		ConnectToDB()
	}
	var tempSegment Segment
	db.Table(segmentTableToEdit).Where("id = ?", id).Find(&tempSegment).Row()
	return tempSegment
}

func FindSegmentTableByCourseId(courseID uint) []Segment {
	if db == nil {
		ConnectToDB()
	}
	var tempSegment []Segment
	//tempSegment := make([]Segment, 0)
	result := db.Table(segmentTableToEdit).Where("course_id = ?", courseID).Find(&tempSegment)
	if result != nil {
		log.Println(result)
	}
	returnSegment := make([]Segment, 0)
	result2, _ := result.Rows()
	var tempCourse2 Segment
	for result2.Next() {
		if err3 := result.ScanRows(result2, &tempCourse2); err3 != nil {
			log.Println(err3)
		}
		returnSegment = append(returnSegment, tempCourse2)
	}
	return returnSegment
}

// desc: Find categories belonging to segment
// comment: If using categories table for all segments
func FindCategoriesBySegmentId(segmentID uint) []SegmentCategory {
	if db == nil {
		ConnectToDB()
	}
	var tempSegment []SegmentCategory

	result := db.Table(categoriesTableToEdit).Where("course_id = ?", segmentID).Find(&tempSegment)
	if result != nil {
		log.Println(result)
	}
	returnSegment := make([]SegmentCategory, 0)
	result2, _ := result.Rows()
	var tempSegment2 SegmentCategory
	for result2.Next() {
		if err3 := result.ScanRows(result2, &tempSegment2); err3 != nil {
			log.Println(err3)
		}
		returnSegment = append(returnSegment, tempSegment2)
	}
	return returnSegment
}
