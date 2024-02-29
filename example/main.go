package main

import (
	"fmt"
	"insert_it"
)

func main() {

	fmt.Println(insert_it.GetServicePointTypes())
	return

	insert_it.Region = insert_it.Regions["Krefeld"]

	fmt.Println(insert_it.GetStreets())

	fmt.Println(insert_it.GetHouseNumbers("Wiesenstr."))

	fmt.Println(insert_it.GetCalendar("Wiesenstr.", insert_it.HouseNumber{HouseNumberStart: 3}))

}
