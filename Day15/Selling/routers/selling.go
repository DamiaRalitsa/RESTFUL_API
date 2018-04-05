package routers

import (
	c "Dam/Day15/Selling/controllers"

	"github.com/gorilla/mux"
)

func setSellingRouters(router *mux.Router) *mux.Router {
	// 1. b) create function in folder controller first!
	router.HandleFunc("/selling", c.GetSelling).Methods("GET")

	return router

}
