package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Custodial Partner
func (e Example) ExampleCustodialPartnerService_List(rc *rize.Client) {
	resp, err := rc.CustodialPartners.List(context.Background())
	if err != nil {
		log.Fatal("Error fetching Custodial Partners\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Custodial Partners:", string(output))
}

// Get Custodial Partner
func (e Example) ExampleCustodialPartnerService_Get(rc *rize.Client) {
	resp, err := rc.CustodialPartners.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Custodial Partner\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Custodial Partner:", string(output))
}
