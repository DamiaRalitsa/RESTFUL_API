package controllers

import (
	"Dam/Day15/data"
	"encoding/json"
	"net/http"
)

func GetOfficer(w http.ResponseWriter, r *http.Request) {
	// get data
	// to get the data, need connection
	// 1. c) create connection
	context := Context{}
	//defer context.Close()
	d := DBInitial(context.DB)
	defer d.Close()
	// get data from repository
	// make command in repository
	// 1. d) create repo Item
	repo := &data.OfficerRepository{d}
	officer := data.GetAll(repo)
	// process eror
	// print data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(officer)
	w.Write(mdl)
}
