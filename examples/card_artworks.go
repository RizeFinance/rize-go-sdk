package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Card Artwork
func ExampleCardArtworkService_List(rc *rize.Client) {
	params := &rize.CardArtworkListParams{
		ProgramUID: "DbxJUHVuqt3C7hGK",
		Limit:      100,
		Offset:     10,
	}

	resp, err := rc.CardArtworks.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Card Artwork\n", err)
	}

	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Card Artwork:", string(output))
}

// Get Card Artwork
func ExampleCardArtworkService_Get(rc *rize.Client) {
	resp, err := rc.CardArtworks.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Card Artwork\n", err)
	}

	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get CardArtwork:", string(output))
}
