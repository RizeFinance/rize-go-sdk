package examples

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rizefinance/rize-go-sdk"
	"github.com/rizefinance/rize-go-sdk/internal"
)

// List customers
func (e Example) ExampleCustomerService_List(rc *rize.Client) {
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
		log.Fatal("Error fetching customers\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Customers:", string(output))
}

// Create new customer
func (e Example) ExampleCustomerService_Create(rc *rize.Client) {
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
		log.Fatal("Error creating customer\n", err)
	}

	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("New Customer:", string(output))
}

// Get customer
func (e Example) ExampleCustomerService_Get(rc *rize.Client) {
	resp, err := rc.Customers.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching customer\n", err)
	}

	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Customer:", string(output))
}

// Update customer
func (e Example) ExampleCustomerService_Update(rc *rize.Client) {
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
		log.Fatal("Error updating customer\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Update Customer:", string(output))
}

// Delete customer
func (e Example) ExampleCustomerService_Delete(rc *rize.Client) {
	params := &rize.CustomerDeleteParams{
		ArchiveNote: "Archiving customer note",
	}
	if _, err := rc.Customers.Delete(context.Background(), "EhrQZJNjCd79LLYq", params); err != nil {
		log.Fatal("Error archiving customer\n", err)
	}
	log.Println("Customer Deleted")
}

// Confirm Identity
func (e Example) ExampleCustomerService_ConfirmPIIData(rc *rize.Client) {
	resp, err := rc.Customers.ConfirmPIIData(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error confirming identity\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Confirm customer identity:", string(output))
}

// Lock customer
func (e Example) ExampleCustomerService_Lock(rc *rize.Client) {
	params := &rize.CustomerLockParams{
		LockNote:   "Fraud detected",
		LockReason: "Customer Reported Fraud",
	}
	resp, err := rc.Customers.Lock(context.Background(), "EhrQZJNjCd79LLYq", params)
	if err != nil {
		log.Fatal("Error locking customer\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Lock Customer:", string(output))
}

// Unlock Customer
func (e Example) ExampleCustomerService_Unlock(rc *rize.Client) {
	params := &rize.CustomerLockParams{
		LockNote:     "Fraud detected",
		UnlockReason: "Customer Reported Fraud",
	}
	// Unlock Customer
	resp, err := rc.Customers.Unlock(context.Background(), "EhrQZJNjCd79LLYq", params)
	if err != nil {
		log.Fatal("Error unlocking customer\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Unlock Customer:", string(output))
}

// Update Profile Response
func (e Example) ExampleCustomerService_UpdateProfileResponses(rc *rize.Client) {
	// Update Profile Response with string response
	params := &rize.CustomerProfileResponseParams{
		ProfileRequirementUID: "ptRLF7nQvy8VoqM1",
		ProfileResponse: &internal.CustomerProfileResponseItem{
			Response: "Response string",
		},
	}
	resp, err := rc.Customers.UpdateProfileResponses(context.Background(), "EhrQZJNjCd79LLYq", []*rize.CustomerProfileResponseParams{params})
	if err != nil {
		log.Fatal("Error updating profile response\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Update Profile Response:", string(output))

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
		log.Fatal("Error updating profile response (ordered_list)\n", err)
	}
	outputList, _ := json.MarshalIndent(res, "", "\t")
	log.Println("Update Profile Response:", string(outputList))
}
