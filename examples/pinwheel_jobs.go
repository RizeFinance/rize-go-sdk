package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Pinwheel Jobs
func ExamplePinwheelJobService_List(rc *rize.Client) {
	params := &rize.PinwheelJobListParams{
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
		Limit:               100,
		Offset:              10,
	}
	resp, err := rc.PinwheelJobs.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Pinwheel Jobs\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Pinwheel Jobs:", string(output))
}

// Create Pinwheel Job
func ExamplePinwheelJobService_Create(rc *rize.Client) {
	params := &rize.PinwheelJobCreateParams{
		JobNames:             []string{"direct_deposit_switch"},
		SyntheticAccountUID:  "4XkJnsfHsuqrxmeX",
		Amount:               1000,
		DisablePartialSwitch: true,
		OrganizationName:     "Chipotle Mexican Grill, Inc.",
		SkipWelcomeScreen:    true,
	}
	resp, err := rc.PinwheelJobs.Create(context.Background(), params)
	if err != nil {
		log.Fatal("Error creating Pinwheel Job\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Create Pinwheel Job:", string(output))
}

// Get PinwheelJob
func ExamplePinwheelJobService_Get(rc *rize.Client) {
	resp, err := rc.PinwheelJobs.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Pinwheel Job\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Pinwheel Job:", string(output))
}
