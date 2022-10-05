package internal

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

var (
	doc    *openapi3.T
	router routers.Router
	err    error
	server = "https://sandbox.rizefs.com/api/v1/"
	ctx    = context.Background()
)

// ValidateRequest is used to validate the given input according to the loaded OpenAPIv3 spec
func ValidateRequest(method string, path string, params url.Values, body io.Reader) (*openapi3filter.RequestValidationInput, error) {
	// Load openapi specification file
	if doc == nil {
		doc, err = openapi3.NewLoader().LoadFromFile("../testdata/rize_external.yaml")
		if err != nil {
			return nil, err
		}
	}

	// Create new router for testing API routes
	if router == nil {
		router, err = gorillamux.NewRouter(doc)
		if err != nil {
			return nil, err
		}
	}

	// Load the server URL
	if len(doc.Servers) > 0 {
		server = doc.Servers[0].URL
	}

	// Create a new request
	req, _ := http.NewRequest(method, fmt.Sprintf("%s/%s", server, path), body)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.URL.RawQuery = params.Encode()

	// Match request to an OpenAPI route
	route, pathParams, err := router.FindRoute(req)
	if err != nil {
		return nil, err
	}

	// Validate request
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
		Options: &openapi3filter.Options{
			AuthenticationFunc: func(c context.Context, input *openapi3filter.AuthenticationInput) error {
				return nil
			},
		},
	}
	if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
		return nil, err
	}

	// Validate request body
	if body != nil {
		if err = openapi3filter.ValidateRequestBody(ctx, requestValidationInput, route.Operation.RequestBody.Value); err != nil {
			return nil, err
		}
	}

	return requestValidationInput, nil
}

// ValidateResponse is used to validate the given input according to the loaded OpenAPIv3 spec
func ValidateResponse(status int, body []byte, requestValidationInput *openapi3filter.RequestValidationInput) error {
	// Validate response
	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 status,
		Header: http.Header{
			"Content-Type": []string{
				"application/json",
				"image/png",
				"image/jpeg",
				"application/pdf",
				"application/octet-stream",
			},
		},
	}

	if body != nil {
		responseValidationInput.SetBodyBytes(body)
	}

	if err = openapi3filter.ValidateResponse(ctx, responseValidationInput); err != nil {
		return err
	}

	return nil
}
