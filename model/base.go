package model

import (
	"github.com/go-sql-driver/mysql"
	"encoding/json"
	"log"
	"time"
	"strings"
)


type JsonNullDate struct {
	mysql.NullTime
}

type JsonNullTime struct {
	mysql.NullTime
}

func (v *JsonNullDate) MarshalJSON() ([]byte, error) {
	if v.Valid {
		log.Println("MarshalJSON() v.Valid")
		return json.Marshal(v.Time)
	}
	log.Println("MarshalJSON() Invalid")
	return json.Marshal(nil)
}

func (v *JsonNullDate) UnmarshalJSON(data []byte) error {
	const LAYOUT = "2006-01-02"
	var err error
	if data != nil {
		v.Time, err = time.Parse(LAYOUT, strings.Trim(string(data), `"`))
		if err != nil {
			log.Println("Error in time.Parse()", err.Error())
			return err
		}
		v.Valid = true
		log.Println("data = ", string(data), "v.Time = ", v.Time)
	} else {
		v.Valid = false
	}
	return nil
}