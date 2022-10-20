package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Adjustments
func (e Example) ExampleAdjustmentService_List(rc *rize.Client) {
	params := &rize.AdjustmentListParams{
		CustomerUID:            "uKxmLxUEiSj5h4M3",
		AdjustmentTypeUID:      "2Ej2tsFbQT3S1HYd",
		ExternalUID:            "PT3sH7oxxQPwchrf",
		USDAdjustmentAmountMax: 10,
		USDAdjustmentAmountMin: 5,
		Sort:                   "adjustment_type_name_asc",
	}
	resp, err := rc.Adjustments.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Adjustments\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Adjustments:", string(output))
}

// Create Adjustment
func (e Example) ExampleAdjustmentService_Create(rc *rize.Client) {
	params := &rize.AdjustmentCreateParams{
		ExternalUID:         "partner-generated-id",
		CustomerUID:         "kaxHFJnWvJxRJZxq",
		USDAdjustmentAmount: "2.43",
		AdjustmentTypeUID:   "KM2eKbR98t4tdAyZ",
	}
	resp, err := rc.Adjustments.Create(context.Background(), params)
	if err != nil {
		log.Fatal("Error creating Adjustment\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Create Adjustment:", string(output))
}

// Get Adjustment
func (e Example) ExampleAdjustmentService_Get(rc *rize.Client) {
	resp, err := rc.Adjustments.Get(context.Background(), "exMDShw6yM3NHLYV")
	if err != nil {
		log.Fatal("Error fetching Adjustment\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Adjustment:", string(output))
}

// List Adjustment Types
func (e Example) ExampleAdjustmentService_ListAdjustmentTypes(rc *rize.Client) {
	params := &rize.AdjustmentTypeListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		ProgramUID:  "DbxJUHVuqt3C7hGK",
	}
	resp, err := rc.Adjustments.ListAdjustmentTypes(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Adjustment Types\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Adjustment Types:", string(output))
}

// Get Adjustment Type
func (e Example) ExampleAdjustmentService_GetAdjustmentType(rc *rize.Client) {
	resp, err := rc.Adjustments.GetAdjustmentType(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Adjustment Type\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Adjustment Type:", string(output))
}
