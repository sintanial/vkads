package vkobj

import (
	"strconv"
	"strings"
	"time"
)

type Date time.Time

func DateRef(t time.Time) *Date {
	d := Date(t)
	return &d
}

func (d *Date) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*d = Date(t) //set result using the pointer
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format("2006-01-02") + `"`), nil
}

type DateTime time.Time

func (d *DateTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02 15:04:05", value) //parse time
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format("2006-01-02 15:04:05") + `"`), nil
}

type Float64 float64

func Float64Ref(f float64) *Float64 {
	fr := Float64(f)
	return &fr
}

func (d *Float64) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	*d = Float64(f)
	return nil
}

func (d Float64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(d), 'f', -1, 64)), nil
}

type Int int

func IntRef(f int) *Int {
	ir := Int(f)
	return &ir
}

func (d *Int) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	*d = Int(i)
	return nil
}

func (d Int) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(d))), nil
}
