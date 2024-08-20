package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dixonwille/wmenu"
)

func main() {

	// Connect to database & create DB connection pool.
	db, err := sql.Open("sqlite3", "../../quisqueya.db")
	checkErr(err)
	defer db.Close()

	menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })

	menu.Option("Add a new Province", 0, false, nil)
	menu.Option("Find a Province", 1, true, nil)
	menu.Option("Update a Province's information", 2, false, nil)
	menu.Option("Delete a Province", 3, false, nil)
	menu.Option("Quit application", 4, false, nil)

	menuerr := menu.Run()
	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:
		reader := bufio.NewReader(os.Stdin)

		province_name := addNewField(reader, "Enter the province name: ")
		capital := addNewField(reader, "Enter the name of the capital city: ")
		region := addNewField(reader, "Enter the region's name: ")
		department := addNewField(reader, "Enter the department's name: ")
		area_km2 := addNewField(reader, "Enter the area's size in km2: ")
		population_2021 := addNewField(reader, "Enter province's population: ")
		density := addNewField(reader, "Enter the province's density: ")
		established_year := addNewField(reader, "Enter the year in which this province was established: ")

		newProvince := province{
			province:         province_name,
			capital:          capital,
			region:           region,
			department:       department,
			area_km2:         area_km2,
			population_2021:  population_2021,
			density:          density,
			established_year: established_year,
		}

		addProvince(db, newProvince)

		break

	case 1:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a province's name or capital to search by: ")
		searchString, _ := reader.ReadString('\n')
		searchString = strings.TrimSuffix(searchString, "\n")
		province := searchProvince(db, searchString)

		fmt.Printf("Found %v results\n", len(province))

		for _, foundProvince := range province {
			fmt.Printf("\nID: %d\n------\nProvince: %s\nCapital: %s\nRegion: %s\nDepartment: %s\nArea(km2): %s\nPopulation(2021): %s\nDensity: %s\nEstablished Year: %s\n", foundProvince.id, foundProvince.province, foundProvince.capital, foundProvince.region, foundProvince.department, foundProvince.area_km2, foundProvince.population_2021, foundProvince.density, foundProvince.established_year)
		}
		break

	case 2:

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the id of the Province you wish to update: ")
		provinceId, _ := reader.ReadString('\n')
		currentProvince, err := getProvinceById(db, provinceId)
		if err != nil {
			log.Fatal(err)
		}

		updateExistingField("Province's Name", currentProvince.province, &currentProvince.province)
		updateExistingField("Province's Capital", currentProvince.capital, &currentProvince.capital)
		updateExistingField("Province's Region", currentProvince.region, &currentProvince.region)
		updateExistingField("Province's Department", currentProvince.department, &currentProvince.department)
		updateExistingField("Province's Population", currentProvince.population_2021, &currentProvince.population_2021)
		updateExistingField("Province's Density", currentProvince.density, &currentProvince.density)
		updateExistingField("Province's Established Year", currentProvince.established_year, &currentProvince.established_year)

		affected := updateProvince(db, currentProvince)

		if affected == 1 {
			fmt.Println("One Province modified")
		}

		break
	case 3:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the ID of the province you want to delete: ")
		searchString, _ := reader.ReadString('\n')

		idToDelete := strings.TrimSuffix(searchString, "\n")

		affected := deleteProvince(db, idToDelete)

		if affected == 1 {
			fmt.Printf("Deleted Province with ID %s from database", idToDelete)
		}

		break
	case 4:
		fmt.Println("Goodbye!")
		os.Exit(0)
	}
}
