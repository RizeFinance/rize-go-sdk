package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Products
func (e Example) ExampleProductService_List(rc *rize.Client) {
	params := &rize.ProductListParams{
		ProgramUID: "pQtTCSXz57fuefzp",
	}
	resp, err := rc.Products.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Products\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Products:", string(output))
}

// Get Product
func (e Example) ExampleProductService_Get(rc *rize.Client) {
	resp, err := rc.Products.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Product\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Product:", string(output))
}
