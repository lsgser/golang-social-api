package faculties

import (
	"encoding/json"
	"net/http"

	CO "../config"
	"github.com/julienschmidt/httprouter"
)

//Show all the Faculties for a specific id
func ShowFaculties(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	CO.AddSafeHeaders(&w)
	var schoolID int64
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
