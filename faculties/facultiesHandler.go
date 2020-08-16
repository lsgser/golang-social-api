package faculties

import (
	"encoding/json"
	"net/http"

	"strconv"

	CO "../config"
	"github.com/julienschmidt/httprouter"
)

//ShowFaculties : Show all the Faculties for a specific id
func ShowFaculties(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	CO.AddSafeHeaders(&w)

	school, err := strconv.Atoi(params.ByName("s"))

	if school == 0 {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"ID was not provided"}`))
		return
	}

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	faculties, err := GetFaculties(int64(school))

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}
	err = json.NewEncoder(w).Encode(faculties)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}
}

//ShowFaculty : displays a single faculty based on their faculty id
func ShowFaculty(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	CO.AddSafeHeaders(&w)
	school, err := strconv.Atoi(params.ByName("f")) // change this back to an integer with strconv ??
	//schoolID = strconv.ParseInt(req.FormValue("s"),10,64) ??

	if school == 0 { // change the empty string to a value ??
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"school id was not provided"}`))
		return
	}
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	faculties, err := GetFaculty(int64(school))

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	err = json.NewEncoder(w).Encode(faculties)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

}
