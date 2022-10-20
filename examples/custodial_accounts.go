package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Custodial Accounts
func (e Example) ExampleCustodialAccountService_List(rc *rize.Client) {
	params := &rize.CustodialAccountListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		ExternalUID: "client-generated-id",
		Limit:       100,
		Offset:      10,
		Liability:   true,
		Type:        "dda",
	}
	resp, err := rc.CustodialAccounts.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Custodial Accounts\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Custodial Accounts:", string(output))
}

// Get Custodial Account
func (e Example) ExampleCustodialAccountService_Get(rc *rize.Client) {
	resp, err := rc.CustodialAccounts.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Custodial Account\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Custodial Account:", string(output))
}
