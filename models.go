package main

import (
	"database/sql"
	"fmt"
	"log"

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

	fmt.Printf("The province of %v was just added. \n", newProvince.province)
}

func searchProvince(db *sql.DB, searchString string) []province {
	rows, err := db.Query("SELECT id, province, capital, region, department, area_km2, population_2021, density, established_year FROM provinces WHERE province like '%" + searchString + "%' OR capital like '%" + searchString + "%'")
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	provinces := make([]province, 0)

	for rows.Next() {
		foundProvince := province{}
		err = rows.Scan(&foundProvince.id, &foundProvince.province, &foundProvince.capital, &foundProvince.region, &foundProvince.department, &foundProvince.area_km2, &foundProvince.population_2021, &foundProvince.density, &foundProvince.established_year)
		if err != nil {
			log.Fatal(err)
		}

		provinces = append(provinces, foundProvince)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return provinces
}
