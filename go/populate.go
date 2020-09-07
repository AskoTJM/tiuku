package swagger

import "log"

func populateSchool() {
	if err := db.Create(&Schools{
		ID:       0,
		Finnish:  "Oulun Ammattikorkeakoulu",
		English:  "Oulu University of Applied Sciences",
		Campuses: []Campuses{},
	}).Error; err != nil {
		log.Panic("Problems populating table of Schools. <go/populate.go->populateSchool>")
	}

	if err := db.Create(&Campuses{
		ID:         0,
		Finnish:    "Linnanmaan Kampus",
		English:    "Campus Linnanmaa",
		Apartments: []Apartments{},
	}).Error; err != nil {
		log.Panic("Problems populating table of Campuses. <go/populate.go->populateSchool>")
	}

	if err := db.Create(&Apartments{
		ID:      0,
		Finnish: "Informaatioteknologia",
		English: "Information Technology",
		Degrees: []Degrees{},
	}).Error; err != nil {
		log.Panic("Problems populating table of Apartments. <go/populate.go->populateSchool>")
	}

	if err := db.Create(&Degrees{
		ID:      0,
		Finnish: "Insinööri (AMK), tieto- ja viestintätekniikka",
		English: "Bachelor of Engineering, Information Technology",
	}).Error; err != nil {
		log.Panic("Problems populating table of Degrees. <go/populate.go->populateSchool>")
	}
	/*
		db.Model(&Schools{
			ID:       0,
			Finnish:  "",
			English:  "",
			Campuses: []Campuses{},
		}).AddForeignKey()
	*/
}

func populateStudents() {

	if err := db.Create(&StudentUsers{
		ID:              0,
		StudentID:       "oppi1",
		AnonID:          "Anon1",
		StudentName:     "Oppilas 1",
		StudentSegments: StudentSegments{},
		StudentEmail:    "oppilas1@oppilaitos.fi",
		StudentClass:    "tit1",
	}).Error; err != nil {
		log.Panic("Problems populating table of table StudentUsers. <go/populate.go->populateStudents>")
	}

	if err := db.Create(&StudentUsers{
		ID:              0,
		StudentID:       "oppi2",
		AnonID:          "Anon2",
		StudentName:     "Oppilas 2",
		StudentSegments: StudentSegments{},
		StudentEmail:    "oppilas2@oppilaitos.fi",
		StudentClass:    "tit1",
	}).Error; err != nil {
		log.Panic("Problems populating table of table StudentUsers. <go/populate.go->populateStudents>")
	}

	if err := db.Create(&StudentUsers{
		ID:              0,
		StudentID:       "oppi3",
		AnonID:          "Anon3",
		StudentName:     "Oppilas 3",
		StudentSegments: StudentSegments{},
		StudentEmail:    "oppilas3@oppilaitos.fi",
		StudentClass:    "tit2",
	}).Error; err != nil {
		log.Panic("Problems populating table of table StudentUsers. <go/populate.go->populateStudents>")
	}
}
