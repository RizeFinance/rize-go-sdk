package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Pools
func (e Example) ExamplePoolService_List(rc *rize.Client) {
	params := &rize.PoolListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		ExternalUID: "client-generated-id",
		Limit:       100,
		Offset:      10,
	}
	resp, err := rc.Pools.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching pools\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Pools:", string(output))
}

// Get Pool
func (e Example) ExamplePoolService_Get(rc *rize.Client) {
	resp, err := rc.Pools.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching pool\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Pool:", string(output))
}
