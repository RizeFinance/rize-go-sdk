package rize

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rizefinance/rize-go-sdk/internal"
)

// Complete Customer{} response data
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
		ProfileResponse: &internal.CustomerProfileResponseItem{
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
		DOB:          internal.DOB(time.Now()),
		SSN:          "111-11-1111",
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

func TestListCustomers(t *testing.T) {
	params := &CustomerListParams{
		UID:              "uKxmLxUEiSj5h4M3",
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
		ExternalUID:      "client-generated-id",
		PoolUID:          "wTSMX1GubP21ev2h",
		Limit:            100,
		Offset:           0,
		Sort:             "first_name_asc",
	}

	resp, err := rc.Customers.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching customers\n", err)
	}

	if err := validateSchema(http.MethodGet, "/customers", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCreateCustomer(t *testing.T) {
	params := &CustomerCreateParams{
		ExternalUID:  "client-generated-id",
		CustomerType: "primary",
		Email:        "olive.oyl@popeyes.com",
	}

	resp, err := rc.Customers.Create(context.Background(), params)
	if err != nil {
		t.Fatal("Error creating customer\n", err)
	}

	if err := validateSchema(http.MethodPost, "/customers", http.StatusCreated, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetCustomer(t *testing.T) {
	resp, err := rc.Customers.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching customer\n", err)
	}

	if err := validateSchema(http.MethodGet, "/customers/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestUpdateCustomer(t *testing.T) {
	cup := &CustomerUpdateParams{
		Email: "olive.oyl@rizemoney.com",
		Details: CustomerDetails{
			FirstName:    "Olive",
			MiddleName:   "Olivia",
			LastName:     "Oyl",
			Suffix:       "Jr.",
			Phone:        "5555551212",
			BusinessName: "Oliver's Olive Emporium",
			DOB:          internal.DOB(time.Now()),
			SSN:          "111-22-3333",
			SSNLastFour:  "3333",
			Address: &CustomerAddress{
				Street1:    "123 Abc St.",
				Street2:    "Apt 2",
				City:       "Chicago",
				State:      "IL",
				PostalCode: "12345",
			},
		},
	}
	resp, err := rc.Customers.Update(context.Background(), "EhrQZJNjCd79LLYq", cup)
	if err != nil {
		t.Fatal("Error updating customer\n", err)
	}

	if err := validateSchema(http.MethodPut, "/customers/{uid}", http.StatusOK, nil, cup, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestDeleteCustomer(t *testing.T) {
	cd := &CustomerDeleteParams{
		ArchiveNote: "Archiving customer note",
	}

	// Delete customer
	if _, err := rc.Customers.Delete(context.Background(), "EhrQZJNjCd79LLYq", cd); err != nil {
		t.Fatal("Error archiving customer\n", err)
	}

	if err := validateSchema(http.MethodDelete, "/customers/{uid}", http.StatusNoContent, nil, cd, nil); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestConfirmPIIData(t *testing.T) {
	// Confirm Identity
	resp, err := rc.Customers.ConfirmPIIData(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error confirming identity\n", err)
	}

	if err := validateSchema(http.MethodPut, "/customers/{uid}/identity_confirmation", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestLockCustomer(t *testing.T) {
	cl := &CustomerLockParams{
		LockNote:   "Fraud detected",
		LockReason: "Customer Reported Fraud",
	}
	// Lock customer
	resp, err := rc.Customers.Lock(context.Background(), "EhrQZJNjCd79LLYq", cl)
	if err != nil {
		t.Fatal("Error locking customer\n", err)
	}

	if err := validateSchema(http.MethodPut, "/customers/{uid}/lock", http.StatusOK, nil, cl, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestUnlockCustomer(t *testing.T) {
	cl := &CustomerLockParams{
		LockNote:     "Fraud detected",
		UnlockReason: "Customer Reported Fraud",
	}
	// Unlock Customer
	resp, err := rc.Customers.Unlock(context.Background(), "EhrQZJNjCd79LLYq", cl)
	if err != nil {
		t.Fatal("Error unlocking customer\n", err)
	}

	if err := validateSchema(http.MethodPut, "/customers/{uid}/unlock", http.StatusOK, nil, cl, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestUpdateProfileResponses(t *testing.T) {
	// Update Profile Response with string response
	cpp := &CustomerProfileResponseParams{
		ProfileRequirementUID: "ptRLF7nQvy8VoqM1",
		ProfileResponse: &internal.CustomerProfileResponseItem{
			Response: "Response string",
		},
	}
	_, err := rc.Customers.UpdateProfileResponses(context.Background(), "EhrQZJNjCd79LLYq", []*CustomerProfileResponseParams{cpp})
	if err != nil {
		t.Fatal("Error updating profile response\n", err)
	}

	// TODO: Add string-type profileResponse parameter to the OpenAPI schema file
	// if err := validateSchema(http.MethodPut, "/customers/{uid}/update_profile_responses", http.StatusOK, nil, nil, resp); err != nil {
	// 	t.Fatalf(err.Error())
	// }

	// Update Profile Response with ordered list response
	cro := &CustomerProfileResponseParams{
		ProfileRequirementUID: "ptRLF7nQvy8VoqM1",
		ProfileResponse: &internal.CustomerProfileResponseItem{
			Num0: "string",
			Num1: "string",
			Num2: "string",
		},
	}
	res, err := rc.Customers.UpdateProfileResponses(context.Background(), "EhrQZJNjCd79LLYq", []*CustomerProfileResponseParams{cro})
	if err != nil {
		t.Fatal("Error updating profile response (ordered_list)\n", err)
	}

	if err := validateSchema(http.MethodPut, "/customers/{uid}/update_profile_responses", http.StatusOK, nil, cro, res); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCreateSecondaryCustomer(t *testing.T) {
	// Secondary Customers
	scp := &SecondaryCustomerParams{
		ExternalUID:        "7002440b-9b98-4a8b-82b9-4503fe8c6bf0",
		PrimaryCustomerUID: "kbF5TGrmwGizQuzZ",
		Email:              "tomas@example.com",
		Details: &CustomerDetails{
			FirstName:  "Olive",
			MiddleName: "Olivia",
			LastName:   "Oyl",
			Suffix:     "Jr.",
			DOB:        internal.DOB(time.Now()),
			Address: &CustomerAddress{
				Street1:    "123 Abc St.",
				Street2:    "Apt 2",
				City:       "Chicago",
				State:      "IL",
				PostalCode: "12345",
			},
		},
	}
	resp, err := rc.Customers.CreateSecondaryCustomer(context.Background(), scp)
	if err != nil {
		t.Fatal("Error creating secondary customer\n", err)
	}

	if err := validateSchema(http.MethodPost, "/customers/create_secondary", http.StatusCreated, nil, scp, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
