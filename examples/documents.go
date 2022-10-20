package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Documents
func (e Example) ExampleDocumentService_List(rc *rize.Client) {
	params := &rize.DocumentListParams{
		DocumentType:        "monthly_statement",
		Month:               1,
		Year:                2020,
		CustodialAccountUID: "yqyYk5b1xgXFFrXs",
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
		Limit:               100,
		Offset:              10,
	}
	resp, err := rc.Documents.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Documents\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Documents:", string(output))
}

// Get Document
func (e Example) ExampleDocumentService_Get(rc *rize.Client) {
	resp, err := rc.Documents.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Document\n", err)
	}

	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Document:", string(output))
}

// View Document
func (e Example) ExampleDocumentService_View(rc *rize.Client) {
	resp, err := rc.Documents.View(context.Background(), "u8EHFJnWvJxRJZxa")
	if err != nil {
		log.Fatal("Error viewing document\n", err)
	}

	log.Println("View Document:", resp.Status)
}
