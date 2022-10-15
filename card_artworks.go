package rize

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Handles all CardArtwork operations
type cardArtworkService service

// CardArtwork data type
type CardArtwork struct {
	UID        string `json:"uid,omitempty"`
	IsDefault  bool   `json:"is_default,omitempty"`
	Name       string `json:"name,omitempty"`
	ProgramUID string `json:"program_uid,omitempty"`
	Staged     bool   `json:"staged,omitempty"`
	StyleID    string `json:"style_id,omitempty"`
}

// CardArtworkListParams builds the query parameters used in querying Card Artwork
type CardArtworkListParams struct {
	ProgramUID string `url:"program_uid,omitempty" json:"program_uid,omitempty"`
	Limit      int    `url:"limit,omitempty" json:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty" json:"offset,omitempty"`
}

// CardArtworkResponse is an API response containing a list of Card Artwork
type CardArtworkResponse struct {
	BaseResponse
	Data []*CardArtwork `json:"data"`
}

// List retrieves a list of Card Artworks, optionally filtering by program
func (c *cardArtworkService) List(ctx context.Context, clp *CardArtworkListParams) (*CardArtworkResponse, error) {
	// Build CardArtworkListParams into query string params
	v, err := query.Values(clp)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodGet, "card_artworks", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CardArtworkResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single Card Artwork resource
func (c *cardArtworkService) Get(ctx context.Context, uid string) (*CardArtwork, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := c.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("card_artworks/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CardArtwork{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
