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
	db, err := sql.Open("sqlite3", "./quisqueya.db")
	checkErr(err)
	defer db.Close()

	menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })

	menu.Option("Add a new Province", 0, false, nil)
	menu.Option("Find a Province", 1, true, nil)
	menu.Option("Update a Province's information", 2, false, nil)
	menu.Option("Delete a Province", 3, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the province name: ")
		province_name, _ := reader.ReadString('\n')
		if province_name != "\n" {
			province_name = strings.TrimSuffix(province_name, "\n")
		}
		fmt.Print("Enter the name of the capital city: ")
		capital, _ := reader.ReadString('\n')
		if capital != "\n" {
			capital = strings.TrimSuffix(capital, "\n")
		}
		fmt.Print("Enter the region's name: ")
		region, _ := reader.ReadString('\n')
		if region != "\n" {
			region = strings.TrimSuffix(region, "\n")
		}
		fmt.Print("Enter the department's name: ")
		department, _ := reader.ReadString('\n')
		if department != "\n" {
			department = strings.TrimSuffix(department, "\n")
		}
		fmt.Print("Enter the region's name: ")
		area_km2, _ := reader.ReadString('\n')
		if area_km2 != "\n" {
			area_km2 = strings.TrimSuffix(area_km2, "\n")
		}
		fmt.Print("Enter province's population: ")
		population_2021, _ := reader.ReadString('\n')
		if population_2021 != "\n" {
			population_2021 = strings.TrimSuffix(population_2021, "\n")
		}
		fmt.Print("Enter the province's density: ")
		density, _ := reader.ReadString('\n')
		if density != "\n" {
			density = strings.TrimSuffix(density, "\n")
		}
		fmt.Print("Enter the year in which this province was established: ")
		established_year, _ := reader.ReadString('\n')
		if established_year != "\n" {
			established_year = strings.TrimSuffix(established_year, "\n")
		}

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
			fmt.Printf("\nID: %d\n------\nProvince: %s\nCapital: %s\nRegion: %s\nDepartment: %s\nArea(km2): %s\nPopulation(2021): %s\nDensity: %s\nDepartment: %s\n", foundProvince.id, foundProvince.province, foundProvince.capital, foundProvince.region, foundProvince.department, foundProvince.area_km2, foundProvince.population_2021, foundProvince.density, foundProvince.established_year)
		}
		break

	case 2:
		fmt.Println("Update a Province's information")
	case 3:
		fmt.Println("Deleting a Province by ID")
	case 4:
		fmt.Println("Quitting application")
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
