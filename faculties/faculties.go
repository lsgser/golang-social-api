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

//GetAllFaculties returns all faculties based on the school id integer that's provided
func GetAllFaculties(schoolID int64) (Faculty, error) {

	school := Faculty{}
	database, errors := CO.GetDB()

	if errors != nil {
		return school, errors
	}
	rows, errors := database.Query("SELECT school_id, faculty FROM faculties WHERE school_id=?", schoolID)

	if errors != nil {
		return school, errors
	}
	defer rows.Close()
	for rows.Next() {
		errors := rows.Scan(&school.SchoolID, &school.Faculties)

		if errors != nil {
			return school, errors
		}
	}
	errors = rows.Err()
	if errors != nil {
		return school, errors
	}
	return school, nil
}
