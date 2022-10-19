package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

func ExampleAuthService_GetToken(rc *rize.Client) {
	resp, err := rc.Auth.GetToken(context.Background())
	if err != nil {
		log.Fatal("Error fetching Auth token\n", err)
	}

	o, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Auth Token:", string(o))
}
