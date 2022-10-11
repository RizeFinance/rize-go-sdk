package internal

import (
	"fmt"
	"strings"
	"time"
)

// DOB is a custom date of birth time format
type DOB time.Time

// MarshalJSON is a custom json marshaller for formatting DOB strings
func (d DOB) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(d).Format("2006-01-02"))), nil
}

// UnmarshalJSON is a custom json unmarshaller for formatting DOB strings
func (d *DOB) UnmarshalJSON(b []byte) error {
	date := strings.ReplaceAll(string(b), "\"", "") // clean up extra ""
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}
	*d = DOB(t)
	return nil
}
