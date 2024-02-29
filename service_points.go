package insert_it

import (
	"encoding/json"
	"net/http"
)

type PointObjectType struct {
	ID             int
	AppDisplayName string
}

type PointObject struct {
	ID                   int
	Lat                  float64
	Lon                  float64
	Remark               *string
	BmsPointObjectTypeId int
}

func GetServicePointTypes() (types []PointObjectType, err error) {
	resp, err := http.Get(svcUrl + servicePointTypes)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&types); err != nil {
		return nil, err
	}

	return
}

func GetServicePoints() (points []PointObject, err error) {
	resp, err := http.Get(svcUrl + servicePoints)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&points); err != nil {
		return nil, err
	}

	return
}
