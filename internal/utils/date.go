package utils

import (
	"encoding/json"
	"time"
)

type Date struct {
	time.Time
}

func NewDate(t time.Time) Date {
	return Date{t}
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(data []byte) error {
	// Parse the date in "YYYY-MM-DD" format
	parsedTime, err := time.Parse("2006-01-02", string(data[1:len(data)-1]))
	if err != nil {
		return err
	}
	d.Time = parsedTime
	return nil
}
