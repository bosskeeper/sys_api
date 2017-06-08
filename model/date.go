package model

import (
	"time"
	"log"
	"strings"
	"database/sql/driver"
	"errors"
)

type MyTime struct{
	*time.Time
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Format("\"2006-01-02 15:04:05\"")), nil
}


// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *MyTime) UnmarshalJSON(data []byte) (err error) {
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse("\"2006-01-02 15:04:05\"", string(data))
	*t = MyTime{&tt}
	return
}


const DateFormat = "2006-01-02" // yyyy-mm-dd

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	log.Println("json.Unmashaller == Overide UnmarshalJSON()", string(data))
	var err error
	d.Time, err = time.Parse(DateFormat, strings.Trim(string(data), `"`)) // << ตรงนี้ต้องทำการ Trim(") ออก
	if err != nil {
		return err
	}
	return nil
}

func (d *Date) Value() (driver.Value, error) {
	return d.Time, nil
}

func (d *Date) Scan(src interface{}) error {
	if date, ok := src.(time.Time); ok {
		d.Time = date
		return nil
	}
	//d.Time = Date(src.(time.Time))
	return errors.New("wrong type it's not time.Time")
}
