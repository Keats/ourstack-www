package ourstack

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

// Global connection, set in main and used throughout the functions
var db *sqlx.DB

func GetDBConnection(url string) {
	db = sqlx.MustConnect("postgres", url)
}

type Company struct {
	Id          int
	Name        string
	Website     string
	Description string
	Size        string
	Remote      bool
}

func GetCompanies() []Company {
	companies := []Company{}
	err := db.Select(&companies, "SELECT * FROM companies")
	if err != nil {
		log.Printf("%v", err.Error())
	}

	return companies
}
