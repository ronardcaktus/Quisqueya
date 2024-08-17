package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

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

func getProvinceById(db *sql.DB, targetID string) (province, error) {
	stmt, err := db.Prepare("SELECT id, province, capital, region, department, area_km2, population_2021, density, established_year FROM provinces WHERE id = ?")
	if err != nil {
		return province{}, err
	}
	defer stmt.Close()
	targetProvince := province{}

	err = stmt.QueryRow(targetID).Scan(&targetProvince.id, &targetProvince.province, &targetProvince.capital, &targetProvince.region, &targetProvince.department, &targetProvince.area_km2, &targetProvince.population_2021, &targetProvince.density, &targetProvince.established_year)
	if err != nil {
		return province{}, err
	}

	return targetProvince, nil
}

func updateProvince(db *sql.DB, targetProvince province) int64 {

	stmt, err := db.Prepare("UPDATE provinces set province = ?, capital = ?, region = ?, department = ?, area_km2 = ?, population_2021 = ?, density = ?, established_year = ? where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(targetProvince.province, targetProvince.capital, targetProvince.region, targetProvince.department, targetProvince.area_km2, targetProvince.population_2021, targetProvince.density, targetProvince.established_year, targetProvince.id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	return affected
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// updateExistingField prompts the user for a field name while displaying
// the current one. It also cleans up inputed data: trim white space and
// ensures it isn't empty.
func updateExistingField(fieldName string, currentValue string, field *string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (Currently %s): ", fieldName, currentValue)
	input, _ := reader.ReadString('\n')
	if input != "\n" {
		*field = strings.TrimSuffix(input, "\n")
	}
}

// addNewField reads standard input from the user, prints instructions,
// and ensures that the input is cleaned prior to saving it to the DB.
func addNewField(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	if input != "\n" {
		input = strings.TrimSuffix(input, "\n")
	}
	return input
}
