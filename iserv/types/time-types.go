package types

import (
	"fmt"
	"strings"
	"time"
)

type IServTime1 time.Time

func (c *IServTime1) UnmarshalJSON(b []byte) error {
	value := strings.Split(strings.Trim(string(b), `"`), "+")[0]
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02T15:04:05", value) //parse time
	if err != nil {
		return err
	}
	*c = IServTime1(t) //set result using the pointer
	return nil
}

func (c IServTime1) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("2006-01-02T15:04:05+01:00") + `"`), nil
}

func (c IServTime1) Format(f fmt.State, r rune) {
	f.Write([]byte(time.Time(c).Format("2006-01-02T15:04:05+01:00")))
}

type IServTime2 time.Time

func (c *IServTime2) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	value = strings.Split(value, "+")[0]

	t, err := time.Parse("02.01.2006 15:04", value) //parse time
	if err != nil {
		return err
	}
	*c = IServTime2(t) //set result using the pointer
	return nil
}

func (c IServTime2) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("02.01.2006 15:04") + `"`), nil
}

func (c IServTime2) Format(f fmt.State, r rune) {
	f.Write([]byte(time.Time(c).Format("02.01.2006 15:04")))
}
