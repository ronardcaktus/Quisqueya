package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type province struct {
	id               int
	province         string
	capital          string
	region           string
	department       string
	area_km2         string
	population_2021  string
	density          string
	established_year string
}

func addProvince(db *sql.DB, newProvince province) {

	stmt, _ := db.Prepare("INSERT INTO provinces (id, province, capital, region, department, area_km2, population_2021, density, established_year ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	stmt.Exec(nil, newProvince.province, newProvince.capital, newProvince.region, newProvince.department, newProvince.area_km2, newProvince.population_2021, newProvince.density, newProvince.established_year)
	defer stmt.Close()

	fmt.Printf("The province of %v with an ID %v was just added. \n", newProvince.province, newProvince.id)
}
