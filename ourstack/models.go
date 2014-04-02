package ourstack

import (
	"fmt"
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

type Tech struct {
	Name string
}

type Location struct {
	City    string
	Country string
}

func GetCompanies() []Company {
	companies := []Company{}
	err := db.Select(&companies, "SELECT * FROM companies")
	if err != nil {
		log.Printf("%v", err.Error())
	}

	return companies
}

// Gets a simple list of all tech we have in the db (only their name)
func GetTechList() []string {
	techs := []Tech{}

	err := db.Select(&techs, "SELECT DISTINCT name FROM techs")
	if err != nil {
		log.Printf("%v", err.Error())
	}

	techList := []string{}

	for _, tech := range techs {
		techList = append(techList, tech.Name)
	}

	return techList
}

func GetLocationList() []string {
	locations := []Location{}
	err := db.Select(&locations, "SELECT DISTINCT city, country FROM locations")
	if err != nil {
		log.Printf("%v", err.Error())
	}

	locationsList := []string{}

	for _, location := range locations {
		locationsList = append(locationsList, fmt.Sprintf("%s (%s)", location.City, location.Country))

	}

	return locationsList
}
