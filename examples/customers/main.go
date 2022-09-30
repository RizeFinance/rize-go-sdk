package main

import (
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
	"github.com/rizefinance/rize-go-sdk/internal"
	rize "github.com/rizefinance/rize-go-sdk/platform"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	config := rize.RizeConfig{
		ProgramUID:  internal.CheckEnvVariable("program_uid"),
		HMACKey:     internal.CheckEnvVariable("hmac_key"),
		Environment: internal.CheckEnvVariable("environment"),
		Debug:       true,
	}

	// Create new Rize client
	rc, err := rize.NewRizeClient(&config)
	if err != nil {
		log.Fatal("Error building RizeClient\n", err)
	}

	// List customers
	clp := rize.CustomerListParams{
		Limit: 10,
	}
	cl, err := rc.Customers.List(&clp)
	if err != nil {
		log.Fatal("Error fetching customers\n", err)
	}
	output, _ := json.Marshal(cl)
	log.Println("List Customers:", string(output))

	// Create new customer
	ccp := rize.CustomerCreateParams{
		CustomerType: "primary",
		Email:        "thomas@example.com",
	}
	cc, err := rc.Customers.Create(&ccp)
	if err != nil {
		log.Fatal("Error creating new customer\n", err)
	}
	output, _ = json.Marshal(cc)
	log.Println("New Customer:", string(output))

	// Get customer
	cg, err := rc.Customers.Get("EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching customer\n", err)
	}
	output, _ = json.Marshal(cg)
	log.Println("Get Customer:", string(output))

	// Update customer
	cup := rize.CustomerUpdateParams{
		Email: "olive.oyl@rizemoney.com",
		Details: rize.CustomerDetails{
			FirstName: "Olive",
			LastName:  "Oyl",
		},
	}
	cu, err := rc.Customers.Update("EhrQZJNjCd79LLYq", &cup)
	if err != nil {
		log.Fatal("Error updating customer\n", err)
	}
	output, _ = json.Marshal(cu)
	log.Println("Update Customer:", string(output))

	// Delete customer
	cdl, err := rc.Customers.Delete("EhrQZJNjCd79LLYq", "Archiving customer note")
	if err != nil {
		log.Fatal("Error archiving customer\n", err)
	}
	output, _ = json.Marshal(cdl)
	log.Println("Delete Customer:", string(output))

	// Confirm Identity
	ci, err := rc.Customers.ConfirmPIIData("EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error confirming identity\n", err)
	}
	output, _ = json.Marshal(ci)
	log.Println("Confirm customer identity:", string(output))

	// Lock customer
	clk, err := rc.Customers.Lock("EhrQZJNjCd79LLYq", "note", "reason")
	if err != nil {
		log.Fatal("Error locking customer\n", err)
	}
	output, _ = json.Marshal(clk)
	log.Println("Lock Customer:", string(output))

	// Unlock Customer
	cul, err := rc.Customers.Unlock("EhrQZJNjCd79LLYq", "note", "reason")
	if err != nil {
		log.Fatal("Error unlocking customer\n", err)
	}
	output, _ = json.Marshal(cul)
	log.Println("Unlock Customer:", string(output))

	// Update Profile Response
	cpp := []*rize.CustomerProfileResponseParams{
		{
			ProfileRequirementUID: "ptRLF7nQvy8VoqM1",
			ProfileResponse:       "",
		},
	}
	cpr, err := rc.Customers.UpdateProfileResponses("EhrQZJNjCd79LLYq", cpp)
	if err != nil {
		log.Fatal("Error updating profile response\n", err)
	}
	output, _ = json.Marshal(cpr)
	log.Println("Update Profile Response:", string(output))

	// Update Profile Response (ordered_list)
	cro := []*rize.CustomerProfileResponseOrderedListParams{{
		ProfileRequirementUID: "ptRLF7nQvy8VoqM1",
		ProfileResponse: &rize.CustomerProfileResponseItem{
			Num0: "string",
		},
	}}
	col, err := rc.Customers.UpdateProfileResponsesOrderedList("EhrQZJNjCd79LLYq", cro)
	if err != nil {
		log.Fatal("Error updating profile response (ordered_list)\n", err)
	}
	output, _ = json.Marshal(col)
	log.Println("Update Profile Response (ordered_list):", string(output))

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
	sc, err := rc.Customers.CreateSecondaryCustomer(&scp)
	if err != nil {
		log.Fatal("Error creating secondary customer\n", err)
	}
	output, _ = json.Marshal(sc)
	log.Println("Secondary Customer:", string(output))
}
