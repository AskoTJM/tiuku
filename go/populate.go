package swagger

func populateSchool() {
	db.Create(&Schools{
		ID:       0,
		Finnish:  "Oulun Ammattikorkeakoulu",
		English:  "Oulu University of Applied Sciences",
		Campuses: []Campuses{},
	})

	db.Create(&Campuses{
		ID:         0,
		Finnish:    "Linnanmaan Kampus",
		English:    "Campus Linnanmaa",
		Apartments: []Apartments{},
	})

	db.Create(&Apartments{
		ID:      0,
		Finnish: "Informaatioteknologia",
		English: "Information Technology",
		Degrees: []Degrees{},
	})

	db.Create(&Degrees{
		ID:      0,
		Finnish: "Insinööri (AMK), tieto- ja viestintätekniikka",
		English: "Bachelor of Engineering, Information Technology",
	})
}
