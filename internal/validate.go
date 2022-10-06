package internal

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

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

	return keys
}

// Traverse openapi3.SchemaRefs
func traverseRefs(refs openapi3.SchemaRefs) {
	for _, r := range refs {
		if r.Value.AllOf != nil {
			traverseRefs(r.Value.AllOf)
		}

		if r.Value.Properties != nil {
			traverseProps(r.Value.Properties)
		}

		if r.Value.Items != nil {
			traverseRef(r.Value.Items)
		}
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
			// Check for keys that are ints and convert
			if _, err := strconv.Atoi(k); err == nil {
				k = fmt.Sprintf("Num%s", k)
			}

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

// JSONKeys will extract key values from a JSON object
func JSONKeys(data map[string]interface{}) []string {
	keys := []string{}

	for key, value := range data {
		// Check for json object
		if val, ok := value.(map[string]interface{}); ok {
			keys = append(keys, key)
			for _, subkey := range JSONKeys(val) {
				keys = append(keys, subkey)
			}
		} else if val, ok := value.([]interface{}); ok {
			// Check for json array
			for _, subkey := range val {
				// Check for json object
				if subObject, ok := subkey.(map[string]interface{}); ok {
					for _, subObjectKey := range JSONKeys(subObject) {
						keys = append(keys, subObjectKey)
					}
				}
			}
		} else {
			keys = append(keys, key)
		}
	}
	return keys
}

// Difference will find any string from a[] that are not present in b[]
func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
