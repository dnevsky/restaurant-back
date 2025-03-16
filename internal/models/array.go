package models

import (
	"database/sql/driver"
	"encoding/json"
)

type ArrayBooking []Booking
type ArrayFood []Food

func (array ArrayBooking) Value() (driver.Value, error) {
	return anyValue(array)
}

func (array *ArrayBooking) Scan(src interface{}) error {
	source, err := sourceToBytes(src)
	if err != nil {
		return err
	}

	var tmpArray ArrayBooking
	err = json.Unmarshal(source, &tmpArray)
	if err != nil {
		return err
	}

	*array = tmpArray

	return nil
}

func (array ArrayFood) Value() (driver.Value, error) {
	return anyValue(array)
}

func (array *ArrayFood) Scan(src interface{}) error {
	source, err := sourceToBytes(src)
	if err != nil {
		return err
	}

	var tmpArray ArrayFood
	err = json.Unmarshal(source, &tmpArray)
	if err != nil {
		return err
	}

	*array = tmpArray

	return nil
}

func anyValue(value any) (string, error) {
	jsonValue, err := json.Marshal(value)
	return string(jsonValue), err
}

func sourceToBytes(src interface{}) ([]byte, error) {
	var source []byte
	switch v := src.(type) {
	case []byte:
		source = v
	case string:
		source = []byte(v)
	default:
		return nil, ErrTypeAssertionFailed
	}
	return source, nil
}

func anyScan(src interface{}) (interface{}, error) {
	source, err := sourceToBytes(src)
	if err != nil {
		return nil, err
	}

	var i interface{}
	err = json.Unmarshal(source, &i)

	return i, err
}
