package rize_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/rizefinance/rize-go-sdk"
)

// Complete CardArtwork{} response data
var artwork = &rize.CardArtwork{
	UID:        "EhrQZJNjCd79LLYq",
	IsDefault:  true,
	Name:       "Rize Default",
	ProgramUID: "kaxHFJnWvJxRJZxr",
	Staged:     true,
	StyleID:    "000",
}

func TestListCardArtwork(t *testing.T) {
	params := &rize.CardArtworkListParams{
		ProgramUID: "DbxJUHVuqt3C7hGK",
		Limit:      100,
		Offset:     10,
	}

	resp, err := rc.CardArtworks.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Card Artwork\n", err)
	}

	if err := validateSchema(http.MethodGet, "/card_artworks", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetCardArtwork(t *testing.T) {
	resp, err := rc.CardArtworks.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Card Artwork\n", err)
	}

	if err := validateSchema(http.MethodGet, "/card_artworks/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
