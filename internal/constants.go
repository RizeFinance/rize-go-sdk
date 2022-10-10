package internal

import (
	"fmt"
	"strings"
	"time"
)

// DOB is a custom date of birth time format
type DOB time.Time

// Shared SDK Constants
const (
	// Platform SDK
	BasePath       = "api/v1"
	TimeoutSeconds = time.Second * 30
	TokenMaxAge    = int64((time.Hour * 23) / time.Millisecond)
	// Message Queue
	MQSendTimeout    = time.Millisecond * 5000
	MQReceiveTimeout = time.Millisecond * 5000

	SDKVersion = "0.0.1"
)

// Environments are Rize infrastructure tiers
var Environments = []string{"sandbox", "integration", "production"}

// MQServices are the available Message Queue topic services
var MQServices = []string{"adjustments", "customer", "debit_card", "synthetic_account", "transfer", "transaction"}

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
