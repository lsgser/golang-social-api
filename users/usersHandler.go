package users

import(
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	CO "../config"
)

//ShowUsers displays all users
func ShowUsers(w http.ResponseWriter , req *http.Request , _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	users,err := GetAllUsers()

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))		
		return
	}

	err = json.NewEncoder(w).Encode(users)

	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte(`{"status":"Something went wrong"}`))
		return
	}	
}

//ShowUser displays a single user based on their username
func ShowUser(w http.ResponseWriter , req *http.Request , params httprouter.Params){
	CO.AddSafeHeaders(&w)
	username := params.ByName("u")
	
	if username == "" {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Username was not provided"}`))
		return
	}
	
	user,err := GetUser(username)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return	
	}

	if user.Username == ""{
		w.WriteHeader(404)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}			
}

//addUser is a function that will add a new user to the database
func AddUser(w http.ResponseWriter , req *http.Request , _ httprouter.Params){
	CO.AddSafeHeaders(&w)	
	body := req.Body
	user := NewUser()
	defer body.Close()
	err := json.NewDecoder(body).Decode(user)
	if err != nil{
		w.WriteHeader(400)
		/*
		w.Write([]byte(`{"status":"JSON decode error"}`))
		*/
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	if user.Username == "" || user.Name == "" || user.Surname == "" || user.Email == "" || user.Password == "" {
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"Fill in all fields"}`))
		return
	}

	err = user.SaveUser()
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}	

	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Success"}`))	 					
}

func RemoveUser(w http.ResponseWriter, req *http.Request, params httprouter.Params){
	CO.AddSafeHeaders(&w)
	username := params.ByName("u")

	if username == ""{
		w.WriteHeader(404)
		w.Write([]byte(`{"status":"Username not provided"}`))
		return
	}

	err := DeleteUser(username)

	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}
	
	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Success"}`))			
}

func UpdateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	CO.AddSafeHeaders(&w)
	body := req.Body
	user := NewEditForm()
	defer body.Close()
	err := json.NewDecoder(body).Decode(user)	
	
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"JSON decoding error"}`))
		return
	}
	
	if user.Username == "" && user.NewUsername == "" && user.Name == "" && user.Surname == "" && user.Email == ""{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"No field was filled for edits"}`))
		return
	}

	err = user.EditUser()
	
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte(`{"status":"`+err.Error()+`"}`))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(`{"status":"Success"}`))	
}
