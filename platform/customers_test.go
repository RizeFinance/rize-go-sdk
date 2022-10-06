package platform

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/rizefinance/rize-go-sdk/internal"
)

var customer = &Customer{
	UID:                "y9reyPMNEWuuYSC1",
	ExternalUID:        "partner-generated-id",
	ActivatedAt:        time.Now(),
	CreatedAt:          time.Now(),
	CustomerType:       "primary",
	Email:              "olive.oyl@rizemoney.com",
	KYCStatus:          "manual_review",
	KYCStatusReasons:   []string{"Approved"},
	LockReason:         "other",
	LockedAt:           time.Now(),
	PIIConfirmedAt:     time.Now(),
	PoolUIDs:           []string{"HiuQZJNjCd79LLYq", "NoPJB9g9ZQTh5qMv"},
	PrimaryCustomerUID: "EhrQZJNjCd79LLYq",
	ProfileResponses: []*CustomerProfileResponse{{
		ProfileRequirement: "Please provide your approximate annual income in USD.",
		ProfileResponse: &CustomerProfileResponseItem{
			Num0: "string",
			Num1: "string",
			Num2: "string",
		},
		ProfileRequirementUID: "ptRLF7nQvy8VoqM1",
	}},
	ProgramUID:            "kaxHFJnWvJxRJZxr",
	SecondaryCustomerUIDs: []string{"464QyebpxbBNrGkX"},
	Status:                "initiated",
	TotalBalance:          "12345.67",
	Details: &CustomerDetails{
		FirstName:    "Olive",
		MiddleName:   "Olivia",
		LastName:     "Oyl",
		Suffix:       "Jr.",
		Phone:        "5555551212",
		BusinessName: "Oliver's Olive Emporium",
		DOB:          time.Now(),
		SSN:          "111111111",
		Address: &CustomerAddress{
			Street1:    "123 Abc St.",
			Street2:    "Apt 2",
			City:       "Chicago",
			State:      "IL",
			PostalCode: "12345",
		},
		SSNLastFour: "3333",
	},
}

func TestList(t *testing.T) {
	params := CustomerListParams{
		Status:           "identity_verified",
		IncludeInitiated: true,
		KYCStatus:        "denied",
		CustomerType:     "primary",
		FirstName:        "Olive",
		LastName:         "Oyl",
		Email:            "olive.oyl@popeyes.com",
		Locked:           false,
		ProgramUID:       "pQtTCSXz57fuefzp",
		BusinessName:     "Business inc",
		PoolUID:          "wTSMX1GubP21ev2h",
		Limit:            100,
		Offset:           0,
		Sort:             "first_name_asc",
	}
	v, err := query.Values(params)
	if err != nil {
		t.Fatal(err)
	}

	// Validate request schema
	input, err := internal.ValidateRequest(http.MethodGet, "customers", v, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Validate response schema
	customers := append([]*Customer{}, customer)
	bytesMessage, err := json.Marshal(&CustomerResponse{Data: customers})
	if err != nil {
		t.Fatal(err)
	}
	if err := internal.ValidateResponse(200, bytesMessage, input); err != nil {
		t.Fatal(err)
	}

	// Generate list of keys from OpenAPI schema response
	keys := internal.RecurseResponseKeys(http.MethodGet, "/customers", 200)

	// Generate list of keys from Customers struct json
	data := make(map[string]interface{})
	if err := json.Unmarshal(bytesMessage, &data); err != nil {
		t.Fatal(err)
	}
	k := internal.JSONKeys(data)

	// Compare OpenAPI spec response keys with keys from SDK struct
	diff := internal.Difference(keys, k)
	if len(diff) > 0 {
		t.Fail()
		t.Log("Customer response is missing the following keys that are present in the OpenAPI schema:\n", diff)
	}
}
