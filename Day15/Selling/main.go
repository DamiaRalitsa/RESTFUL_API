package main

import (
	r "Dam/Day15/Selling/routers"
	"log"
	"net/http"
)

func main() {
	//part 1. create routing
	//create routing function
	router := r.InitRouters()
	// create server configuration
	log.Fatal(http.ListenAndServe(":8887", router))
}
