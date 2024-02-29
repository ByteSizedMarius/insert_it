package insert_it

import (
	"encoding/json"
	"fmt"
	"time"
)

// NextTrashDate is the struct used for GetNextEmptyings
type NextTrashDate struct {
	ID                 int
	Name               string
	AppCalendarIconUrl string
	BinSize            int
	ExecutionDate      time.Time
}

// TrashDate is the struct used for GetCalendar
type TrashDate struct {
	BmsWasteTypeId     int
	BmsWasteTypeName   string
	Deadline           time.Time
	AppCalendarIconUrl string
	EmptyingCountCycle int
	SizeName           string
	Size               int
	Cycle              int
	EmptyingCount      int
	CycleAsText        string
}

func (ntd *NextTrashDate) UnmarshalJSON(data []byte) error {
	type Alias NextTrashDate
	aux := &struct {
		ExecutionDate      string
		AppCalendarIconUrl string
		*Alias
	}{
		Alias: (*Alias)(ntd),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02T15:04:05", aux.ExecutionDate)
	if err != nil {
		return err
	}
	ntd.ExecutionDate = t
	ntd.AppCalendarIconUrl = imgUrl + aux.AppCalendarIconUrl

	return nil
}

func (td *TrashDate) UnmarshalJSON(data []byte) error {
	type Alias TrashDate
	aux := &struct {
		Deadline           string
		AppCalendarIconUrl string
		*Alias
	}{
		Alias: (*Alias)(td),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02T15:04:05", aux.Deadline)
	if err != nil {
		return err
	}
	td.Deadline = t
	td.AppCalendarIconUrl = imgUrl + aux.AppCalendarIconUrl

	return nil
}

//goland:noinspection GoMixedReceiverTypes
func (ntd NextTrashDate) String() string {
	return fmt.Sprintf("ID: %d, Name: %s,  BinSize: %d, ExecutionDate: %s", ntd.ID, ntd.Name, ntd.BinSize, ntd.ExecutionDate.Format("01.02.06"))
}

//goland:noinspection GoMixedReceiverTypes
func (td TrashDate) String() string {
	return fmt.Sprintf("BmsWasteTypeId: %d, BmsWasteTypeName: %s, Deadline: %s,  EmptyingCountCycle: %d, SizeName: %s, Size: %d, Cycle: %d, EmptyingCount: %d, CycleAsText: %s", td.BmsWasteTypeId, td.BmsWasteTypeName, td.Deadline.Format("01.02.06"), td.EmptyingCountCycle, td.SizeName, td.Size, td.Cycle, td.EmptyingCount, td.CycleAsText)
}

type HouseNumber struct {
	HouseNumberStart      int
	HouseNumberStartExtra string
	HouseNumberEnd        string
	HouseNumberEndExtra   string
}

func (hn HouseNumber) String() string {
	str := fmt.Sprintf("%d%s", hn.HouseNumberStart, hn.HouseNumberStartExtra)
	if hn.HouseNumberEnd != "" || hn.HouseNumberEndExtra != "" {
		str += fmt.Sprintf("-%s%s", hn.HouseNumberEnd, hn.HouseNumberEndExtra)
	}
	return str
}
