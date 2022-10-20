package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Synthetic Accounts
func (e Example) ExampleSyntheticAccountService_List(rc *rize.Client) {
	params := &rize.SyntheticAccountListParams{
		CustomerUID:              "uKxmLxUEiSj5h4M3",
		ExternalUID:              "client-generated-id",
		PoolUID:                  "wTSMX1GubP21ev2h",
		Limit:                    100,
		Offset:                   10,
		SyntheticAccountTypeUID:  "q4mdMxMtjXfdbrjn",
		SyntheticAccountCategory: "general",
		Liability:                true,
		Status:                   "active",
		Sort:                     "name_asc",
	}
	resp, err := rc.SyntheticAccounts.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Synthetic Accounts\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Synthetic Accounts:", string(output))
}

// Create Synthetic Account
func (e Example) ExampleSyntheticAccountService_Create(rc *rize.Client) {
	params := &rize.SyntheticAccountCreateParams{
		ExternalUID:             "partner-generated-id",
		Name:                    "New Resource Name",
		PoolUID:                 "kaxHFJnWvJxRJZxq",
		SyntheticAccountTypeUID: "fRMwt6H14ovFUz1s",
		AccountNumber:           "123456789012",
		RoutingNumber:           "123456789",
		PlaidProcessorToken:     "processor-sandbox-96d86f35-ef58-4e4a-826f-4870b5d677f2",
		ExternalProcessorToken:  "processor-sandbox-96d86f35-ef58-4e4a-826f-4870b5d677f2",
	}
	resp, err := rc.SyntheticAccounts.Create(context.Background(), params)
	if err != nil {
		log.Fatal("Error creating Synthetic Account\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Create Synthetic Account:", string(output))
}

// Get Synthetic Account
func (e Example) ExampleSyntheticAccountService_Get(rc *rize.Client) {
	resp, err := rc.SyntheticAccounts.Get(context.Background(), "exMDShw6yM3NHLYV")
	if err != nil {
		log.Fatal("Error fetching Synthetic Account\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Synthetic Account:", string(output))
}

// Update Synthetic Account
func (e Example) ExampleSyntheticAccountService_Update(rc *rize.Client) {
	params := &rize.SyntheticAccountUpdateParams{
		Name: "New Resource Name",
		Note: "note",
	}
	resp, err := rc.SyntheticAccounts.Update(context.Background(), "EhrQZJNjCd79LLYq", params)
	if err != nil {
		log.Fatal("Error updating Synthetic Account\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Update Synthetic Account:", string(output))
}

// Delete Synthetic Account
func (e Example) ExampleSyntheticAccountService_Delete(rc *rize.Client) {
	resp, err := rc.SyntheticAccounts.Delete(context.Background(), "exMDShw6yM3NHLYV")
	if err != nil {
		log.Fatal("Error deleting Synthetic Account\n", err)
	}
	log.Println("Delete Synthetic Account:", resp.Status)
}

// List Synthetic Account Types
func (e Example) ExampleSyntheticAccountService_ListAccountTypes(rc *rize.Client) {
	params := &rize.SyntheticAccountTypeListParams{
		ProgramUID: "EhrQZJNjCd79LLYq",
		Limit:      100,
		Offset:     10,
	}
	resp, err := rc.SyntheticAccounts.ListAccountTypes(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Synthetic Account Types\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Synthetic Account Types:", string(output))
}

// Get Synthetic Account Type
func (e Example) ExampleSyntheticAccountService_GetAccountType(rc *rize.Client) {
	resp, err := rc.SyntheticAccounts.GetAccountType(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Synthetic Account Type\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Synthetic Account Type:", string(output))
}
