package profiles

import(
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	CO "../config"
	UP "../uploads"
	U "../users"
	"strconv"
	"strings"
)

//We set a megabyte to a constant
const(
	MB = 1 << 20
)

func ShowProfiles(w http.ResponseWriter,req *http.Request , _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	
	profiles,err := GetAllProfiles()
	
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(profiles)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
}

func ShowProfile(w http.ResponseWriter, r *http.Request,p httprouter.Params){
	CO.AddSafeHeaders(&w)
	username := p.ByName("u")
	
	if username == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Username was not provided"}`))
		return
	}

	profile,err := GetProfile(username)
	
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(profile)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
}

func AddProfile(w http.ResponseWriter, r *http.Request,_ httprouter.Params){
	CO.AddSafeHeaders(&w)
	profile := NewProfile()
	err := r.ParseMultipartForm(5 * MB)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return 
	}
	
	//Limit upload size to 5MB
	r.Body = http.MaxBytesReader(w,r.Body,5 * MB)
	/*
		For files formFile 
	*/
	file,handler,err := r.FormFile("picture")
	
	if err != nil && err.Error() != "http: no such file"{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return 
	}
	
	if file != nil{
		defer file.Close()
	}
	
	//Check if the userid is valid
	if strings.TrimSpace(r.FormValue("u")) != ""{
		profile.UserID,err = strconv.ParseInt(r.FormValue("u"),10,64)
		if err != nil{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"`+err.Error()+`"}`))
			return
		}
	}	
	
	exists,err := U.UserExists(profile.UserID)
	
	if err != nil && !exists{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
	
	if handler != nil{
		if handler.Filename != ""{
			fileName,err := UP.SingleFileUpload(file,handler)
			if err != nil{
				w.WriteHeader(400)
				w.Write([]byte(`{"status":"`+err.Error()+`"}`))
				return 
			}
			profile.ProfilePicture = fileName
		}
	}
	
	if strings.TrimSpace(r.FormValue("gender")) != ""{
		profile.Gender,err = strconv.Atoi(r.FormValue("gender"))
		if err != nil{
			w.WriteHeader(400)
			w.Write([]byte(`{"status":"`+err.Error()+`"}`))
			return
		}
	}	
	
	if strings.TrimSpace(r.FormValue("birth")) != ""{
		profile.BirthDate = r.FormValue("birth")
	}

	if strings.TrimSpace(r.FormValue("from")) != ""{
		profile.Residence = r.FormValue("from")
	}
	
	err = profile.SaveProfile()
	
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
	
	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Success"}`))		
}
