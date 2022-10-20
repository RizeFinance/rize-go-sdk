package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Customer Products
func (e Example) ExampleCustomerProductService_List(rc *rize.Client) {
	params := &rize.CustomerProductListParams{
		ProgramUID:  "pQtTCSXz57fuefzp",
		ProductUID:  "zbJbEa72eKMgbbBv",
		CustomerUID: "uKxmLxUEiSj5h4M3",
	}
	resp, err := rc.CustomerProducts.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Customer Products\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Customer Products:", string(output))

}

// Create Customer Product
func (e Example) ExampleCustomerProductService_Create(rc *rize.Client) {
	params := &rize.CustomerProductCreateParams{
		CustomerUID: "S62MaHx6WwsqG9vQ",
		ProductUID:  "pQtTCSXz57fuefzp",
	}
	resp, err := rc.CustomerProducts.Create(context.Background(), params)
	if err != nil {
		log.Fatal("Error creating Customer Product\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Create Customer Product:", string(output))
}

// Get Customer Product
func (e Example) ExampleCustomerProductService_Get(rc *rize.Client) {
	resp, err := rc.CustomerProducts.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Customer Product\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Customer Product:", string(output))
}
