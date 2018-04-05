package routers

import (
	c "Dam/Day15/controllers"

	"github.com/gorilla/mux"
)

func setOfficerRouters(router *mux.Router) *mux.Router {
	// 1. b) create function in folder controller first!
	router.HandleFunc("/officer", c.GetOfficer).Methods("GET")

	return router

}
