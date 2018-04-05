package models

type Selling struct {
	Invoice     string
	InvoiceDate string
	Item        int
	Total       float64
	Paid        float64
	Returned    float64
	OfficerCode string
}
