package main

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var dbName string
	var csvName string
	flag.StringVar(&dbName, "db", "", "SQLite database to import to")
	flag.StringVar(&csvName, "csv", "", "CSV file to import from")
	flag.Parse()

	if dbName == "" || csvName == "" {
		flag.PrintDefaults()
		return
	}

	// Open DB
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("ping failed: %s", err)
	}
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS provinces (id INTEGER PRIMARY KEY AUTOINCREMENT, province TEXT, capital TEXT, region TEXT, department TEXT, area_km2 TEXT, population_2021 TEXT, density TEXT, established_year TEXT)")
	if err != nil {
		log.Fatalf("prepare failed: %s", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("exec failed: %s", err)
	}
	fmt.Println("DB created or existing DB found.")
	// Open CSV file
	f, err := os.Open(csvName)
	if err != nil {
		log.Fatalf("open failed: %s", err)
	}
	r := csv.NewReader(f)
	fmt.Println("CSV file opened.")
	// Read the header row.
	_, err = r.Read()
	if err != nil {
		log.Fatalf("missing header row(?): %s", err)
	}
	for {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		province := record[0]
		capital := record[1]
		region := record[2]
		department := record[3]
		area_km2 := record[4]
		population_2021 := record[5]
		density := record[6]
		established_year := record[7]

		stmt, err = db.Prepare("insert into provinces(province, capital, region, department, area_km2, population_2021, density, established_year) values(?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatalf("Insert prepare failed: %s", err)
		}

		_, err = stmt.Exec(province, capital, region, department, area_km2, population_2021, density, established_year)
		if err != nil {
			log.Fatalf("insert failed(%s): %s", province, err)
		}
	}
	fmt.Println("DB populated from CSV file.")

	// Move file up to top-level directory.
	oldPath := "quisqueya.db"
	topDirPath := "../../quisqueya.db"

	movingErr := os.Rename(oldPath, topDirPath)
	if movingErr != nil {
		fmt.Println("Error moving the file:", movingErr)
		return
	}
	fmt.Println("Database moved to top-level directory.")
}
