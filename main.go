package main

import (
	"fmt"
	"log"

	"github.com/dixonwille/wmenu"
)

func main() {

	menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(opts); return nil })

	menu.Option("Add a new Province", 0, false, nil)
	menu.Option("Find a Province", 1, true, nil)
	menu.Option("Update a Province's information", 2, false, nil)
	menu.Option("Delete a Province", 3, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

func handleFunc(opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:
		fmt.Println("Adding a new Province")
	case 1:
		fmt.Println("Finding a Province")
	case 2:
		fmt.Println("Update a Province's information")
	case 3:
		fmt.Println("Deleting a Province by ID")
	case 4:
		fmt.Println("Quitting application")
	}
}
