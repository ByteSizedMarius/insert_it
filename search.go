package insert_it

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetStreets() ([]string, error) {
	resp, err := http.Get(svcUrl + getAllStreets)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []map[string]interface{}
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&data); err != nil {
		return nil, err
	}

	streets := make([]string, 0, len(data))
	for _, item := range data {
		if name, ok := item["Name"].(string); ok {
			streets = append(streets, name)
		}
	}

	return streets, nil
}

func GetStreetFilter(prefix string) ([]string, error) {
	streets, err := GetStreets()
	if err != nil {
		return nil, err
	}

	lowercaseStreets := make([]string, len(streets))
	prefix = strings.ToLower(prefix)
	for i, street := range streets {
		lowercaseStreets[i] = strings.ToLower(street)
	}

	var result []string
	for i, street := range lowercaseStreets {
		if strings.HasPrefix(street, prefix) {
			result = append(result, streets[i])
		}
	}

	return result, nil
}

func GetHouseNumbers(street string) ([]HouseNumber, error) {
	resp, err := http.Get(fmt.Sprintf(svcUrl+getHourseNumbersForStreet, street))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var houseNumbers []HouseNumber
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&houseNumbers); err != nil {
		return nil, err
	}

	if len(houseNumbers) == 0 {
		return nil, fmt.Errorf("no house numbers found for street %s", street)
	}

	return houseNumbers, nil
}
