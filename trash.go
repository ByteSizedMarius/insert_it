package insert_it

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetNextEmptyings returned das jeweils nächste Datum für alle Müll-Typen
func GetNextEmptyings(street string, hn HouseNumber) ([]TrashDate, error) {
	resp, err := http.Get(svcUrl + getNextEmptyings + fmt.Sprintf(searchParams, street, hn.HouseNumberStart, hn.HouseNumberStartExtra, hn.HouseNumberEnd, hn.HouseNumberEndExtra))
	if err != nil {
		return nil, err
	}

	return emptyingsFromBody(resp.Body)
}

// GetCalendar returned alle Müll-Termine für das Jahr
func GetCalendar(street string, hn HouseNumber) ([]TrashDate, error) {
	resp, err := http.Get(svcUrl + getEmptyings + fmt.Sprintf(searchParams, street, hn.HouseNumberStart, hn.HouseNumberStartExtra, hn.HouseNumberEnd, hn.HouseNumberEndExtra))
	if err != nil {
		return nil, err
	}

	return emptyingsFromBody(resp.Body)
}

func emptyingsFromBody(body io.ReadCloser) ([]TrashDate, error) {
	defer body.Close()
	var trashDates []TrashDate
	dec := json.NewDecoder(body)
	if err := dec.Decode(&trashDates); err != nil {
		return nil, err
	}

	return trashDates, nil
}
