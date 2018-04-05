package models

type Officer struct {
	OfficerCode     string `json:"OfficerCode"`
	OfficerName     string `json:"OfficerName"`
	OfficerPassword string
	OfficerStatus   string
}
