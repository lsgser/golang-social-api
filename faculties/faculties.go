package faculties

import (
	CO "../config"
)

//Faculty struct
type Faculty struct {
	ID        int64  `json:"id,omitempty"`
	SchoolID  int64  `json:"school_id,omitempty"`
	Faculties string `json:"faculty,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

//GetAllFaculties : returns all faculties based on the school id integer that's provided
func GetAllFaculties(SchoolID int64) (Faculty, error) {

	faculty := Faculty{}
	database, errors := CO.GetDB()

	if errors != nil {
		return faculty, errors
	}
	//rows, errors := database.Query("SELECT school_id, faculty FROM faculties WHERE school_id=?", schoolID)
	statement, errors := database.Prepare("SELECT school_id,faculty FROM faculties WHERE school_id=?")

	if errors != nil {
		return faculty, errors
	}
	// defer rows.Close()
	defer statement.Close()
	errors = statement.QueryRow(faculty).Scan(&faculty.SchoolID, &faculty.Faculties)
	// for rows.Next() {
	// 	errors := rows.Scan(&faculty.SchoolID, &faculty.Faculties)
	// 	if errors != nil {
	// 		return faculty, errors
	// 	}
	// }
	// errors = rows.Err()
	if errors != nil {
		return faculty, errors
	}

	return faculty, nil
}
