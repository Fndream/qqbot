package tb

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
)

type StringSlice []string

// Scan implements the Scanner interface
func (ss *StringSlice) Scan(value interface{}) error {
	// check if the value is nil
	if value == nil {
		*ss = nil
		return nil
	}
	// check if the value is a byte slice
	if b, ok := value.([]byte); ok {
		// convert the byte slice to a string
		s := string(b)
		// split the string by space and assign it to ss
		*ss = strings.Split(s, " ")
		return nil
	}
	// return an error if the value is not a byte slice
	return errors.New("invalid value type")
}

// Value implements the Valuer interface
func (ss StringSlice) Value() (driver.Value, error) {
	// check if the slice is nil or empty
	if ss == nil || len(ss) == 0 {
		return nil, nil
	}
	// join the slice by space and return it as a driver.Value
	return strings.Join(ss, " "), nil
}

// MarshalJSON implements the json.Marshaler interface
func (ss StringSlice) MarshalJSON() ([]byte, error) {
	// use the default json.Marshal function to encode the slice
	return json.Marshal([]string(ss))
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (ss *StringSlice) UnmarshalJSON(data []byte) error {
	// use the default json.Unmarshal function to decode the data
	return json.Unmarshal(data, (*[]string)(ss))
}
