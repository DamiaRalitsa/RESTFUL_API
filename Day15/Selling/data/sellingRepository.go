package data

import (
	m "Dam/Day15/Selling/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SellingRepository struct {
	DB *sql.DB
}

// for return value GetAll need structure from Item
// 2. a) create model Item
func GetAll(db *SellingRepository) []m.Selling {
	fmt.Println(db.DB)
	all := "Select 	Invoice, InvoiceDate, Item, Total, Paid, Returned, OfficerCode from tblSelling;"
	result, err := db.DB.Query(all)
	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	selling := []m.Selling{}
	for result.Next() {
		var p m.Selling
		if err := result.Scan(&p.Invoice, &p.InvoiceDate, &p.Item, &p.Total, &p.Paid, &p.Returned, &p.OfficerCode); err != nil {
			return nil
		}

		selling = append(selling, p)
	}
	return selling

}
