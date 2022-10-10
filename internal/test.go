package internal

import (
	"context"
	"encoding/json"
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
	keys   []string
	server = "https://sandbox.rizefs.com/api/v1/"
	ctx    = context.Background()
)

// ValidateRequest is used to validate the given input according to the loaded OpenAPIv3 spec
func ValidateRequest(method string, path string, params url.Values, body io.Reader) (*openapi3filter.RequestValidationInput, error) {
	// Load openapi specification file
	if doc == nil {
		u, err := url.Parse("https://cdn.rizefs.com/web-content/openapi/rize_external.yaml")
		if err != nil {
			return nil, err
		}
		doc, err = openapi3.NewLoader().LoadFromURI(u)
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
	req, _ := http.NewRequest(method, fmt.Sprintf("%s%s", server, path), body)
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

// GetRequestKeys returns any request keys (query string or body params) from the OpenAPI schema
func GetRequestKeys(method string, path string) ([]string, []string, error) {
	var paramKeys, bodyKeys []string

	params := doc.Paths.Find(path).GetOperation(method).Parameters
	for _, s := range params {
		paramKeys = append(paramKeys, s.Value.Name)
	}

	reqBody := doc.Paths.Find(path).GetOperation(method).RequestBody
	if reqBody != nil {
		body, err := doc.Paths.Find(path).GetOperation(method).RequestBody.Value.Content.Get("application/json").Schema.MarshalJSON()
		if err != nil {
			return nil, nil, err
		}
		data := make(map[string]interface{})
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, nil, err
		}
		bodyKeys = JSONKeys(data)

	}

	return paramKeys, bodyKeys, nil
}

// RecurseResponseKeys recursively traverse the response object for a given OpenAPI path and returns a list of all properties
func RecurseResponseKeys(method string, path string, status int) []string {
	schema := doc.Paths.Find(path).GetOperation(method).Responses.Get(status).Value.Content.Get("application/json").Schema.Value

	// TODO: Add support for schema.OneOf and schema.AnyOf

	if schema.AllOf != nil {
		traverseRefs(schema.AllOf)
	}

	if schema.Properties != nil {
		traverseProps(schema.Properties)
	}

	// Clear slice between tests
	list := keys
	keys = nil

	return list
}

// Traverse openapi3.SchemaRefs
func traverseRefs(refs openapi3.SchemaRefs) {
	for _, r := range refs {
		traverseRef(r)
	}
}

// Traverse openapi3.SchemaRef
func traverseRef(ref *openapi3.SchemaRef) {
	if ref.Value.AllOf != nil {
		traverseRefs(ref.Value.AllOf)
	}

	if ref.Value.Properties != nil {
		traverseProps(ref.Value.Properties)
	}

	if ref.Value.Items != nil {
		traverseRef(ref.Value.Items)
	}
}

// Traverse openapi3.Schema.Properties
func traverseProps(props openapi3.Schemas) {
	if data, ok := props["data"]; ok {
		traverseRef(data)
	} else if details, ok := props["details"]; ok {
		traverseRef(details)
	} else {
		for k, v := range props {
			// Note the key that was found
			keys = append(keys, k)

			// Check for nested items
			if v.Value.Items != nil {
				traverseRef(v.Value.Items)
			}

			// Check for nested properties
			if v.Value.Properties != nil {
				traverseProps(v.Value.Properties)
			}
		}
	}
}
