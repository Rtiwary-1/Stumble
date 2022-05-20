package routes

import (
	"github.com/gorilla/mux"
	"github.com/users/LENOVO/Downloads/Stumble/pkg/controllers"
)

var RetrieveMatch = func(router *mux.Router){
	
	router.HandleFunc("/matches/",controllers.GetMatches).Methods("GET")
	router.HandleFunc("/local_distance/",controllers.GetAllWithinK).Methods("GET")
	router.HandleFunc("/users/{query}/",controllers.GetUsersByQuery).Methods("GET")
}
