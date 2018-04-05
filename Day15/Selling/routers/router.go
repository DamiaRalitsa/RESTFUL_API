package routers

import "github.com/gorilla/mux"

func InitRouters() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)
	//part 1. a) set routing from officer
	// and then set routing from officer
	router = setSellingRouters(router)
	return router

}
