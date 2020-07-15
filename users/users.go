package users

import(
	"github.com/badoux/checkmail"
	"errors"
	CO "../config"
	"log"
	"strings"
)

//User struct
type User struct{
	ID int64 `json:"id,omitempty"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:"username"`
	Email string `json:"email"`
	EmailVerified string `json:"-"`
	Password string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`	
}

//EditForm struct will be used to edit the users data
type EditForm struct{
	Username string `json:"username"`
	NewUsername string `json:"newusername"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"` 	
}
 
//NewUser returns a pointer to a User
func NewUser() *User{
	//new() makes it easier to create a pointer to non composite types
	return new(User)
}

func NewEditForm() *EditForm{
	return new(EditForm)
}
//GetAllUsers returns slice of User that contains all uses in the database
func GetAllUsers() ([]User,error){
	//Initialize a slice by using make
	users := make([]User,0)
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return users,err
	}
	
	rows,err := db.Query("SELECT name,surname,username,email,created_at FROM users")

	if err != nil{
		err = errors.New("DB query error")
		return users,err
	}
	defer rows.Close()
	for rows.Next(){
		user := User{}
		rows.Scan(&user.Name,&user.Surname,&user.Username,&user.Email,&user.CreatedAt)
		users = append(users,user) 
	}

	log.Println("Users: ",users)
	return users,nil
}

//Gets one user based on the given username
func GetUser(username string) (User,error){
	user := User{}
	db,err := CO.GetDB()
	if err != nil{
		err = errors.New("DB connection error")
		return user,err
	}
	
	stmt,err := db.Prepare("SELECT name,surname,username,email,created_at FROM users WHERE username=?")
	
	if err != nil{
		return user,err	
	}
	
	defer stmt.Close()	 	
	
	err = stmt.QueryRow(username).Scan(&user.Name,&user.Surname,&user.Username,&user.Email,&user.CreatedAt)
	if err != nil{
		return user,err
	}

	return user,nil	
}

//Store a new user 
func (u *User) SaveUser() (err error){
	db,err := CO.GetDB()
	user := User{}
	if err != nil{
		return err
	}

	/*
		Check if the email is valid
	*/
	if checkmail.ValidateFormat(u.Email) != nil{
		err = errors.New("Invalid email format")
		return err
	}

	hashPass,err := CO.HashPassword(u.Password)
	
	if err != nil{
		err = errors.New("Password hash error")
		return
	}
	
	query,err := db.Prepare("SELECT username,email FROM users WHERE username=? OR email=?")
	if err != nil{
		return err	
	}
	defer query.Close()

	_ = query.QueryRow(u.Username,u.Email).Scan(&user.Username,&user.Email)
	
	if strings.ToLower(strings.TrimSpace(user.Username)) == strings.ToLower(strings.TrimSpace(u.Username)){
		err = errors.New("Username already exists")
		return err
	}

	if strings.ToLower(strings.TrimSpace(user.Email)) == strings.ToLower(strings.TrimSpace(u.Email)){
		err = errors.New("Email already exists")
		return err
	}

	stmt,err := db.Prepare("INSERT INTO users (name,surname,username,email,password,created_at,updated_at) VALUES (?,?,?,?,?,NOW(),NOW())")

	if err != nil{
		err = errors.New("Database insert error")
		return err
	}
	
	_,err = stmt.Exec(u.Name,u.Surname,u.Username,u.Email,string(hashPass))

	if err != nil{
		//err = errors.New("Database execution error")
		return err
	}
	
	return nil 	
}

//Delete a user via their username
func DeleteUser(username string) (err error){
	db,err := CO.GetDB()
	user := User{}
	if err != nil{
		return err
	}
	/*
		Check if the username exists in 
		the database
	*/
	stmt,err := db.Prepare("SELECT username FROM users WHERE username = ?")
	
	if err != nil{
		return err
	}
	defer stmt.Close()
	
	/*
		If the query we ran returned an empty row
		it means that the username does no exist
		hence an error will be returned
	*/
	err = stmt.QueryRow(username).Scan(&user.Username)
	
	if err != nil{
		err = errors.New("Username does not exist")
		return err
	}
		
	delUser,err := db.Prepare("DELETE FROM users WHERE username = ?")
	
	if err != nil{
		return err
	}

	_,err = delUser.Exec(username)
	
	if err != nil{
		return err	
	}

	return err
}

//EditUser edits the username,email,name,etc of the user
func (u *EditForm) EditUser() (err error){
	db,err := CO.GetDB()
	e := EditForm{}

	if err != nil{
		return err
	}
	/*
		Check if the username that is provided exists
		in the database
	*/
	log.Println(u)
	log.Println(u.Name)
	if u.Username != ""{
		usernameQ,err := db.Prepare("SELECT username FROM users WHERE username = ?")
		
		if err != nil{
			return err
		}
		
		defer usernameQ.Close()
		err = usernameQ.QueryRow(u.Username).Scan(&e.Username)
	
		if err != nil{
			err = errors.New("Username does not exist")
			return err
		}		
	}else{
		err = errors.New("Current username is required")
		return err
	}

	/*
		Check if the new username that is provided does not
		exist in the database
	*/
	if u.NewUsername != ""{
		newUsernameQ,err := db.Prepare("SELECT username FROM users WHERE username = ?")
		
		if err != nil{
			return err
		}
		
		defer newUsernameQ.Close()
		err = newUsernameQ.QueryRow(u.NewUsername).Scan(&e.NewUsername)
	
		if err == nil{
			err = errors.New("The Username you picked already exists")
			return err
		}
	}
	
	if u.Name != ""{
		updateName,err := db.Prepare("UPDATE users SET name=?,updated_at=NOW() WHERE username=?")
		if err != nil{
			return err
		}
	
		_,err = updateName.Exec(u.Name,u.Username)

		if err != nil{
			return err
		}
	}

	if u.Surname != ""{
		updateSurname,err := db.Prepare("UPDATE users SET surname=?,updated_at=NOW() WHERE username=?")
		if err != nil{
			return err
		}
		_,err = updateSurname.Exec(u.Surname,u.Username)
		
		if err != nil{
			return err
		}
	}

	if u.Email != ""{
		if checkmail.ValidateFormat(u.Email) != nil{
			err = errors.New("Invalid email format")
			return err
		}

		updateEmail,err := db.Prepare("UPDATE users SET email=?,updated_at=NOW() WHERE username=?")
		if err != nil{
			return err
		}
		_,err = updateEmail.Exec(u.Email,u.Username)
		
		if err != nil{
			return err
		}
	}

	/*
		Check if the original username is set and that it does not match
		with the new username
	*/
	if ((strings.ToLower(strings.TrimSpace(u.Username)) != "") && (strings.ToLower(strings.TrimSpace(u.NewUsername)) != "")) &&(strings.ToLower(strings.TrimSpace(u.NewUsername)) != strings.ToLower(strings.TrimSpace(u.Username))){

		updateUname,err := db.Prepare("UPDATE users SET username=?,updated_at=NOW() WHERE username=?")
		if err != nil{
			return err
		}
		_,err = updateUname.Exec(u.NewUsername,u.Username)
		if err != nil{
			return err
		}																	
	}

	return err
}

/*
	Checks if a user exists via their user id
*/
func UserExists(u int64) (bool,error){
	user := User{}
	db,err := CO.GetDB()
	if err != nil{
		err = errors.New("DB connection error")
		return false,err
	}
	
	stmt,err := db.Prepare("SELECT name,surname,username,email,created_at FROM users WHERE id=?")
	
	if err != nil{
		return false,err	
	}
	
	defer stmt.Close()	 	
	
	err = stmt.QueryRow(u).Scan(&user.Name,&user.Surname,&user.Username,&user.Email,&user.CreatedAt)
	if err != nil{
		return false,err
	}

	return true,nil
}
