package faculties

import (
	"encoding/json"
	"net/http"

	CO "../config"
	"github.com/julienschmidt/httprouter"
)

//ShowFaculties : Show all the Faculties for a specific id
func ShowFaculties(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	CO.AddSafeHeaders(&w)
	schoolID = p.ByName("id") // change this back to an integer with strconv ??
	faculties, errors := GetAllFaculties(schoolID)

	if errors != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + errors.Error() + `"}`))
		return
	}
	errors = json.NewEncoder(w).Encode(faculties)

	if errors != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}
}

//

func ShowFaculty(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	CO.AddSafeHeaders(&w)
	facultyName := p.ByName("s")
	schoolID = p.ByName("id") // change this back to an integer with strconv ??
	//schoolID = strconv.ParseInt(req.FormValue("s"),10,64) ??

	if schoolID == "" { // change the empty string to a value ??
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"school id was not provided"}`))
		return
	}

	if facultyName == "" {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Faculty name was not provided"}`))
		return
	}
	faculties, errors := GetFaculty(facultyName, schoolID)

	if errors != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + errors.Error() + `"}`))
		return
	}

	getfaculty, errors := GetFaculty(facultyName, schoolID)

	if errors != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + errors.Error() + `"}`))
		return
	}
	errors = json.NewEncoder(w).Encode(faculties)
	if errors != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + errors.Error() + `"}`))
		return
	}
	errors = json.NewEncoder(w).Encode(getfaculty)
	if errors != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + errors.Error() + `"}`))
		return
	}
}
