package rize_test

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/rizefinance/rize-go-sdk"
	"github.com/rizefinance/rize-go-sdk/internal"
)

// Complete Customer{} response data
var customer = &rize.Customer{
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
	ProfileResponses: []*rize.CustomerProfileResponse{{
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
	Details: &rize.CustomerDetails{
		FirstName:    "Olive",
		MiddleName:   "Olivia",
		LastName:     "Oyl",
		Suffix:       "Jr.",
		Phone:        "5555551212",
		BusinessName: "Oliver's Olive Emporium",
		DOB:          internal.DOB(time.Now()),
		SSNLastFour:  "3333",
		Address: &rize.CustomerAddress{
			Street1:    "123 Abc St.",
			Street2:    "Apt 2",
			City:       "Chicago",
			State:      "IL",
			PostalCode: "12345",
		},
	},
}

func TestCustomerService_List(t *testing.T) {
	params := &rize.CustomerListParams{
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

func TestCustomerService_Create(t *testing.T) {
	params := &rize.CustomerCreateParams{
		CustomerType:       "primary",
		PrimaryCustomerUID: "kbF5TGrmwGizQuzZ",
		ExternalUID:        "client-generated-id",
		Email:              "olive.oyl@popeyes.com",
		Details: &rize.CustomerDetails{
			FirstName:    "Olive",
			MiddleName:   "Olivia",
			LastName:     "Oyl",
			Suffix:       "Jr.",
			Phone:        "5555551212",
			BusinessName: "Oliver's Olive Emporium",
			SSN:          "111-22-3333",
			DOB:          internal.DOB(time.Now()),
			Address: &rize.CustomerAddress{
				Street1:    "123 Abc St.",
				Street2:    "Apt 2",
				City:       "Chicago",
				State:      "IL",
				PostalCode: "12345",
			},
		},
	}

	resp, err := rc.Customers.Create(context.Background(), params)
	if err != nil {
		t.Fatal("Error creating customer\n", err)
	}

	if err := validateSchema(http.MethodPost, "/customers", http.StatusCreated, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCustomerService_Create_SecondaryMissingPrimaryCustomerUID(t *testing.T) {
	params := &rize.CustomerCreateParams{
		CustomerType: "secondary",
	}
	resp, err := rc.Customers.Create(context.Background(), params)

	if err == nil || !strings.Contains(err.Error(), "primary_customer_uid is required") {
		t.Fatal("Expected to see an error regarding missing primary_customer_uid for secondary customer type")
	}
	if resp != nil {
		t.Fatal("Expected response to be nil due to error")
	}
}

func TestCustomerService_Get(t *testing.T) {
	resp, err := rc.Customers.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching customer\n", err)
	}

	if err := validateSchema(http.MethodGet, "/customers/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCustomerService_Update(t *testing.T) {
	params := &rize.CustomerUpdateParams{
		Email:       "olive.oyl@rizemoney.com",
		ExternalUID: "client-generated-id",
		Details: &rize.CustomerDetails{
			FirstName:    "Olive",
			MiddleName:   "Olivia",
			LastName:     "Oyl",
			Suffix:       "Jr.",
			Phone:        "5555551212",
			BusinessName: "Oliver's Olive Emporium",
			DOB:          internal.DOB(time.Now()),
			SSN:          "111-22-3333",
			SSNLastFour:  "3333",
			Address: &rize.CustomerAddress{
				Street1:    "123 Abc St.",
				Street2:    "Apt 2",
				City:       "Chicago",
				State:      "IL",
				PostalCode: "12345",
			},
		},
	}
	resp, err := rc.Customers.Update(context.Background(), "EhrQZJNjCd79LLYq", params)
	if err != nil {
		t.Fatal("Error updating customer\n", err)
	}

	if err := validateSchema(http.MethodPut, "/customers/{uid}", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCustomerService_Delete(t *testing.T) {
	params := &rize.CustomerDeleteParams{
		ArchiveNote: "Archiving customer note",
	}

	// Delete customer
	if _, err := rc.Customers.Delete(context.Background(), "EhrQZJNjCd79LLYq", params); err != nil {
		t.Fatal("Error archiving customer\n", err)
	}

	if err := validateSchema(http.MethodDelete, "/customers/{uid}", http.StatusNoContent, nil, params, nil); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCustomerService_ConfirmPIIData(t *testing.T) {
	// Confirm Identity
	resp, err := rc.Customers.ConfirmPIIData(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error confirming identity\n", err)
	}

	if err := validateSchema(http.MethodPut, "/customers/{uid}/identity_confirmation", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCustomerService_Lock(t *testing.T) {
	params := &rize.CustomerLockParams{
		LockNote:   "Fraud detected",
		LockReason: "Customer Reported Fraud",
	}
	// Lock customer
	resp, err := rc.Customers.Lock(context.Background(), "EhrQZJNjCd79LLYq", params)
	if err != nil {
		t.Fatal("Error locking customer\n", err)
	}

	if err := validateSchema(http.MethodPut, "/customers/{uid}/lock", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCustomerService_Unlock(t *testing.T) {
	params := &rize.CustomerLockParams{
		LockNote:     "Fraud detected",
		UnlockReason: "Customer Reported Fraud",
	}
	// Unlock Customer
	resp, err := rc.Customers.Unlock(context.Background(), "EhrQZJNjCd79LLYq", params)
	if err != nil {
		t.Fatal("Error unlocking customer\n", err)
	}

	if err := validateSchema(http.MethodPut, "/customers/{uid}/unlock", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCustomerService_UpdateProfileResponses(t *testing.T) {
	// Update Profile Response with string response
	params := &rize.CustomerProfileResponseParams{
		ProfileRequirementUID: "ptRLF7nQvy8VoqM1",
		ProfileResponse: &internal.CustomerProfileResponseItem{
			Response: "Response string",
		},
	}
	_, err := rc.Customers.UpdateProfileResponses(context.Background(), "EhrQZJNjCd79LLYq", []*rize.CustomerProfileResponseParams{params})
	if err != nil {
		t.Fatal("Error updating profile response\n", err)
	}

	// TODO: Add string-type profileResponse parameter to the OpenAPI schema file
	// if err := validateSchema(http.MethodPut, "/customers/{uid}/update_profile_responses", http.StatusOK, nil, nil, resp); err != nil {
	// 	t.Fatalf(err.Error())
	// }

	// Update Profile Response with ordered list response
	paramList := &rize.CustomerProfileResponseParams{
		ProfileRequirementUID: "ptRLF7nQvy8VoqM1",
		ProfileResponse: &internal.CustomerProfileResponseItem{
			Num0: "string",
			Num1: "string",
			Num2: "string",
		},
	}
	res, err := rc.Customers.UpdateProfileResponses(context.Background(), "EhrQZJNjCd79LLYq", []*rize.CustomerProfileResponseParams{paramList})
	if err != nil {
		t.Fatal("Error updating profile response (ordered_list)\n", err)
	}

	if err := validateSchema(http.MethodPut, "/customers/{uid}/update_profile_responses", http.StatusOK, nil, paramList, res); err != nil {
		t.Fatalf(err.Error())
	}
}
