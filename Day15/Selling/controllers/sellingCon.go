package controllers

import (
	"Dam/Day15/Selling/data"
	"encoding/json"
	"net/http"
)

func GetSelling(w http.ResponseWriter, r *http.Request) {
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
	repo := &data.SellingRepository{d}
	selling := data.GetAll(repo)
	// process eror
	// print data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(selling)
	w.Write(mdl)
}
