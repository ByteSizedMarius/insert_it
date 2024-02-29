Quick & Dirty implementation of insert-it.de Müllabfuhr API.

```go
package main

import (
	"fmt"
	"insert_it"
)

func main() {
	fmt.Println(insert_it.Regions) // supported regions

	insert_it.Region = insert_it.Regions["Mannheim"]

	// Returns ALL streets
	streets, err := insert_it.GetStreets()

	// Same as above, but with a local filter (startswith)
	streets, err := insert_it.GetStreetFilter("A")

	// Hausnummern for given street (data has a weird format)
	hn, err := insert_it.GetHouseNumbers("Aachener Straße")

	// The the next pickup dates (one date per waste type)
	dates, err := insert_it.GetNextEmptyings("Aachener Straße", hn[0])

	// Get the full calendar
	calendar, err := insert_it.GetCalendar("Aachener Straße", hn[0])

	// Service Points (Glascontainer usw)

	// Point types
	types, err := insert_it.GetServicePointTypes()
	/*
		[
			{1 Altkleidercontainer}
			{2 Papierkorb}
			{3 Altglascontainer}
			{4 Hundekottütenspender}
			{5 Recyclinghöfe}
		]
	*/

	// Get all points
	points, err := insert_it.GetServicePoints()
}

```