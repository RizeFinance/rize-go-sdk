package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Evaluations
func (e Example) ExampleEvaluationService_List(rc *rize.Client) {
	params := &rize.EvaluationListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		Latest:      true,
	}
	resp, err := rc.Evaluations.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Evaluations\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Evaluations:", string(output))
}

// Get Evaluation
func (e Example) ExampleEvaluationService_Get(rc *rize.Client) {
	resp, err := rc.Evaluations.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Evaluation\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Evaluation:", string(output))
}
