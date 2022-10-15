package internal

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// DOB is a custom date of birth time format
type DOB time.Time

// CustomerProfileResponseItem contains the Customer's response to the Profile Requirement.
// Use `Response` to return a string value, or Num0/1/2 for an ordered list
type CustomerProfileResponseItem struct {
	Response string `json:"-,omitempty"`
	Num0     string `json:"0,omitempty"`
	Num1     string `json:"1,omitempty"`
	Num2     string `json:"2,omitempty"`
}

// Used for casting CustomerProfileResponseItem to prevent infinite loop when marshalling/unmarshalling
type responseItem CustomerProfileResponseItem

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

// MarshalJSON is a custom json marshaller for CustomerProfileResponseItem
func (c CustomerProfileResponseItem) MarshalJSON() ([]byte, error) {
	if c.Response != "" {
		return []byte(fmt.Sprintf("\"%s\"", c.Response)), nil
	}

	return json.Marshal((*responseItem)(&c))
}

// UnmarshalJSON is a custom json unmarshaller for CustomerProfileResponseItem
func (c *CustomerProfileResponseItem) UnmarshalJSON(b []byte) error {
	var resp string
	// Check for string response type
	if err := json.Unmarshal(b, &resp); err == nil {
		c.Response = resp
		return nil
	}
	return json.Unmarshal(b, (*responseItem)(c))
}
