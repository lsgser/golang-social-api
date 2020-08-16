package faculties

import (
	"errors"
	"log"

	CO "../config"
)

//Faculty struct
type Faculty struct {
	ID        int64  `json:"id,omitempty"`
	School    int64  `json:"school,omitempty"`  /*University name*/
	Faculty   string `json:"faculty,omitempty"` /*string type so this it the name of the faculty i.e Engineering*/
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

//NewFaculty :  returns a pointer struct of a Faculty type
func NewFaculty() *Faculty {
	return new(Faculty)
}

//GetFaculties : returns all faculties based on the school id integer that's provided
func GetFaculties(s int64) ([]Faculty, error) {

	faculties := make([]Faculty, 0)
	database, err := CO.GetDB()

	if err != nil {
		err = errors.New("DB connection error")
		return faculties, err
	}
	rows, err := database.Query("SELECT school_id, faculty FROM faculties WHERE school_id=?", s)
	//statement, errors := database.Prepare("SELECT faculty FROM faculties WHERE school_id=?")

	if err != nil {
		return faculties, err
	}
	defer rows.Close()

	for rows.Next() {
		faculty := Faculty{}
		rows.Scan(&faculty.ID, &faculty.School, &faculty.Faculty, &faculty.CreatedAt, &faculty.UpdatedAt)
		faculties = append(faculties, faculty)
	}

	log.Println("Faculties :", faculties)
	return faculties, nil
}

//GetFaculty :  returns the faculty data
func GetFaculty(f int64) (Faculty, error) {

	faculty := Faculty{}

	database, err := CO.GetDB()

	if err != nil {
		err := errors.New("DB connection error")
		return faculty, err
	}

	schoolidQuery, err := database.Prepare("SELECT faculty FROM faculties WHERE school_id=?")

	if err != nil {
		return faculty, err
	}

	defer schoolidQuery.Close()

	err = schoolidQuery.QueryRow(f).Scan(&faculty.Faculty)

	if err != nil {
		return faculty, err
	}

	// facultyQuery, err := database.Prepare("SELECT * FROM faculties WHERE id=?")

	// if err != nil {
	//    return facultyQuery, err
	// }

	// defer facultyQuery.Close()

	// err = facultyQuery.QueryRow(ID).Scan(&faculty.ID, &faculty.School, &faculty.Faculty, &faculty.CreatedAt, &faculty.UpdatedAt)

	// if err != nil {
	// 	return faculty, err
	// }

	return faculty, nil
}
