package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
	"github.com/rizefinance/rize-go-sdk"
	"github.com/rizefinance/rize-go-sdk/internal"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	config := rize.Config{
		ProgramUID:  internal.CheckEnvVariable("program_uid"),
		HMACKey:     internal.CheckEnvVariable("hmac_key"),
		Environment: internal.CheckEnvVariable("environment"),
		Debug:       true,
	}

	// Create new Rize client
	rc, err := rize.NewClient(&config)
	if err != nil {
		log.Fatal("Error building RizeClient\n", err)
	}

	// List customers
	clp := rize.CustomerListParams{
		Limit: 10,
	}
	cl, err := rc.Customers.List(context.Background(), &clp)
	if err != nil {
		log.Fatal("Error fetching customers\n", err)
	}
	output, _ := json.MarshalIndent(cl, "", "\t")
	log.Println("List Customers:", string(output))

	// Create new customer
	ccp := rize.CustomerCreateParams{
		CustomerType: "primary",
		Email:        "thomas@example.com",
	}
	cc, err := rc.Customers.Create(context.Background(), &ccp)
	if err != nil {
		log.Fatal("Error creating new customer\n", err)
	}
	output, _ = json.MarshalIndent(cc, "", "\t")
	log.Println("New Customer:", string(output))

	// Get customer
	cg, err := rc.Customers.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching customer\n", err)
	}
	output, _ = json.MarshalIndent(cg, "", "\t")
	log.Println("Get Customer:", string(output))

	// Update customer
	cup := rize.CustomerUpdateParams{
		Email: "olive.oyl@rizemoney.com",
		Details: &rize.CustomerDetails{
			FirstName: "Olive",
			LastName:  "Oyl",
		},
	}
	cu, err := rc.Customers.Update(context.Background(), "EhrQZJNjCd79LLYq", &cup)
	if err != nil {
		log.Fatal("Error updating customer\n", err)
	}
	output, _ = json.MarshalIndent(cu, "", "\t")
	log.Println("Update Customer:", string(output))

	// Delete customer
	cd := rize.CustomerDeleteParams{
		ArchiveNote: "Archiving customer note",
	}
	cdl, err := rc.Customers.Delete(context.Background(), "EhrQZJNjCd79LLYq", &cd)
	if err != nil {
		log.Fatal("Error archiving customer\n", err)
	}
	output, _ = json.MarshalIndent(cdl, "", "\t")
	log.Println("Delete Customer:", string(output))

	// Confirm Identity
	ci, err := rc.Customers.ConfirmPIIData(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error confirming identity\n", err)
	}
	output, _ = json.MarshalIndent(ci, "", "\t")
	log.Println("Confirm customer identity:", string(output))

	// Lock customer
	ck := rize.CustomerLockParams{
		LockNote:   "Fraud detected",
		LockReason: "Customer Reported Fraud",
	}
	clk, err := rc.Customers.Lock(context.Background(), "EhrQZJNjCd79LLYq", &ck)
	if err != nil {
		log.Fatal("Error locking customer\n", err)
	}
	output, _ = json.MarshalIndent(clk, "", "\t")
	log.Println("Lock Customer:", string(output))

	// Unlock Customer
	co := rize.CustomerLockParams{
		LockNote:     "Fraud detected",
		UnlockReason: "Customer Reported Fraud",
	}
	cul, err := rc.Customers.Unlock(context.Background(), "EhrQZJNjCd79LLYq", &co)
	if err != nil {
		log.Fatal("Error unlocking customer\n", err)
	}
	output, _ = json.MarshalIndent(cul, "", "\t")
	log.Println("Unlock Customer:", string(output))

	// Update Profile Response
	cpp := []*rize.CustomerProfileResponseParams{
		{
			ProfileRequirementUID: "ptRLF7nQvy8VoqM1",
			ProfileResponse: &internal.CustomerProfileResponseItem{
				Response: "string",
			},
		},
	}
	cpr, err := rc.Customers.UpdateProfileResponses(context.Background(), "EhrQZJNjCd79LLYq", cpp)
	if err != nil {
		log.Fatal("Error updating profile response\n", err)
	}
	output, _ = json.MarshalIndent(cpr, "", "\t")
	log.Println("Update Profile Response:", string(output))

	// Secondary Customers
	scp := rize.SecondaryCustomerParams{
		PrimaryCustomerUID: "kbF5TGrmwGizQuzZ",
		Details: &rize.CustomerDetails{
			FirstName: "Olive",
			LastName:  "Oyl",
			Address: &rize.CustomerAddress{
				PostalCode: "12345",
			},
		},
	}
	sc, err := rc.Customers.CreateSecondaryCustomer(context.Background(), &scp)
	if err != nil {
		log.Fatal("Error creating secondary customer\n", err)
	}
	output, _ = json.MarshalIndent(sc, "", "\t")
	log.Println("Secondary Customer:", string(output))
}
