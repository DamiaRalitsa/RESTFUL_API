package data

import (
	m "Dam/Day15/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type OfficerRepository struct {
	DB *sql.DB
}

// for return value GetAll need structure from Item
// 2. a) create model Item
func GetAll(db *OfficerRepository) []m.Officer {
	fmt.Println(db.DB)
	all := "Select OfficerCode, OfficerName, OfficerPassword, OfficerStatus from tblOfficer;"
	result, err := db.DB.Query(all)
	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	officer := []m.Officer{}
	for result.Next() {
		var p m.Officer
		if err := result.Scan(&p.OfficerCode, &p.OfficerName, &p.OfficerPassword, &p.OfficerStatus); err != nil {
			return nil
		}

		officer = append(officer, p)
	}
	return officer

}
