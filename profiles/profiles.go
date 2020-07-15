package profiles

import(
	CO "../config" 
	"log"
	"errors"
)

type Profile struct{
	ID int64 `json:"id,omitempty"`
	UserID int64 `json:"user_id,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	Gender int `json:"gender,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
	Residence string `json:"residence,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func NewProfile() *Profile{
	return new(Profile)
}

/*
	Display all the profiles that exist from the database
*/
func GetAllProfiles() ([]Profile,error){
	//Initialize a slice by using make
	profiles := make([]Profile,0)
	db,err := CO.GetDB()
	
	if err != nil {
		err = errors.New("DB connection error")
		return profiles,err 
	}

	rows,err := db.Query("SELECT * FROM profiles") 
	
	if err != nil{
		return profiles,err
	}
	
	defer rows.Close()
	
	for rows.Next(){
		profile := Profile{}
	rows.Scan(&profile.ID,&profile.UserID,&profile.ProfilePicture,&profile.Gender,&profile.BirthDate,&profile.Residence,&profile.CreatedAt,&profile.UpdatedAt)
		profiles = append(profiles,profile)
	}

	log.Println("Profiles :", profiles)
	return profiles,nil	
}

//Get a specific profile data based on a username input
func GetProfile(username string) (Profile,error){
	profile := Profile{}
	db,err := CO.GetDB()
	
	var userID int64

	if err != nil{
		err := errors.New("DB connection error")
		return profile,err		
	}
	
	idQuery,err := db.Prepare("SELECT id FROM users WHERE username=?")
	
	if err != nil{
		return profile,err
	}

	defer idQuery.Close()

	err = idQuery.QueryRow(username).Scan(&userID)

	if err != nil{
		return profile,err
	}

	profileQuery,err := db.Prepare("SELECT * FROM profiles WHERE user_id=?")

	if err != nil{
		return profile,err
	}

	defer profileQuery.Close()

	err = profileQuery.QueryRow(userID).Scan(&profile.ID,&profile.UserID,&profile.ProfilePicture,&profile.Gender,&profile.BirthDate,&profile.Residence,&profile.CreatedAt,&profile.UpdatedAt)
	if err != nil{
		return profile,err	
	}

	return profile,nil								
}

/*
	The SaveProfile method is used to store a profile
	in in the database
*/
func (p *Profile) SaveProfile() (err error){
	profile := Profile{}
	
	db,err := CO.GetDB()
	
	var username string
	

	if err != nil{
		return err
	}
	
	/*
		Check if the user profile user already exists
		in the profiles table to avoid inserting 
		the same user_id twice,also if the userID
		data is empty return an error 
	*/
	if p.UserID != 0 {
		userQuery,err := db.Prepare("SELECT username FROM users WHERE id=?")
		
		if err != nil{
			return err
		}

		defer userQuery.Close()
		
		err = userQuery.QueryRow(p.UserID).Scan(&username)	
		
		if err != nil{
			return err
		}
		
		profileQuery,err := db.Prepare("SELECT user_id FROM profiles WHERE user_id=?")
		
		if err != nil{
			return err
		}

		defer profileQuery.Close()
		
		err = profileQuery.QueryRow(p.UserID).Scan(&profile.UserID)	
		
		if err == nil{
			err = errors.New("User profile already inserted")
			return err
		}

	}else{
		err = errors.New("User ID was not provided")
		return err
	}

	/*
		Insert a new user profile with the given data in the profiles table
		but check if at least one column data is provided
	*/
	if p.ProfilePicture != "" || p.Gender != 0 || p.BirthDate != "" || p.Residence != ""{
		/*
			Passing an empty BirthDate string field gives us an error
			So we'll create two sql insert queries,one for when the birthdate variable
			exists and the other for when the p.BirthDate variable does not exist
		*/
		if p.BirthDate == ""{
			profileStmt,err := db.Prepare("INSERT INTO profiles (user_id,profile_picture,gender,residence,created_at,updated_at) VALUES (?,?,?,?,NOW(),NOW())")
			if err != nil{
				return err
			}
			
			_,err = profileStmt.Exec(p.UserID,p.ProfilePicture,p.Gender,p.Residence)
		
			if err != nil{
				return err
			}	
		}else{
			profileStmt,err := db.Prepare("INSERT INTO profiles (user_id,profile_picture,gender,birth_date,residence,created_at,updated_at) VALUES (?,?,?,?,?,NOW(),NOW())")
			if err != nil{
				return err
			}
			
			_,err = profileStmt.Exec(p.UserID,p.ProfilePicture,p.Gender,p.BirthDate,p.Residence)
		
			if err != nil{
				return err
			}
		}
	}else{
		err = errors.New("Please provide at least one field")
		return err
	}

	return err
}
