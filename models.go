package main

import (
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
