package routes

import (
	F "../faculties"
	P "../profiles"
	U "../users"
	"github.com/julienschmidt/httprouter"
)

//NewRouter : new router returns all router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	//router.GET("/",Index)
	router.GET("/users", U.ShowUsers)
	router.POST("/adduser", U.AddUser)
	router.GET("/user/:u", U.ShowUser)
	router.PUT("/edituser", U.UpdateUser)
	router.DELETE("/removeuser/:u", U.RemoveUser)
	//Profiles
	router.GET("/profiles/", P.ShowProfiles)
	router.GET("/profile/:u", P.ShowProfile)
	router.POST("/addprofile", P.AddProfile)
	//Faculties
	router.GET("/faculties/:id", F.ShowFaculties)
	return router
}
