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
	doc       *openapi3.T
	router    routers.Router
	err       error
	keys      []string
	server    = "https://sandbox.rizefs.com/api/v1/"
	ctx       = context.Background()
	isRequest = false
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
	req.Header.Add("Authorization", "auth-header.payload.signature")
	req.URL.RawQuery = params.Encode()

	// Match request to an OpenAPI route
	route, pathParams, err := router.FindRoute(req)
	if err != nil {
		return nil, err
	}

	// Validate request
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:     req,
		PathParams:  pathParams,
		QueryParams: params,
		Route:       route,
		Options: &openapi3filter.Options{
			AuthenticationFunc: func(c context.Context, input *openapi3filter.AuthenticationInput) error {
				return nil
			},
		},
	}
	if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
		return nil, err
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

// BuildSchemaPathsList generates a list of all path and method combinations in the OpenAPI schema
func BuildSchemaPathsList() []string {
	s := []string{}
	for p, op := range doc.Paths {
		for m := range op.Operations() {
			s = append(s, p+"_"+m)
		}
	}
	return s
}

// FindRoute matches an HTTP request with the path/operation in the OpenAPI schema
func FindRoute(req *http.Request) (string, error) {
	route, _, err := router.FindRoute(req)
	if err != nil {
		return "", err
	}
	return route.Path + "_" + route.Method, nil
}

// GetRequestKeys returns any request keys (query string or body params) from the OpenAPI schema
func GetRequestKeys(method string, path string, status int) ([]string, error) {
	var output []string

	params := doc.Paths.Find(path).GetOperation(method).Parameters
	if len(params) == 0 {
		// Sometimes params exist outside of the Operation block
		params = doc.Paths.Find(path).Parameters
	}
	for _, s := range params {
		// Only check query parameters
		if s.Value.In == "path" || s.Value.In == "header" {
			continue
		}
		output = append(output, s.Value.Name)
	}

	reqBody := doc.Paths.Find(path).GetOperation(method).RequestBody
	if reqBody != nil {
		output, err = RecurseRequestKeys(method, path, status)
	}
	return output, nil
}

// RecurseRequestKeys recursively traverses the Request object for a given OpenAPI path and returns a list of all properties
func RecurseRequestKeys(method string, path string, status int) ([]string, error) {
	isRequest = true
	p := doc.Paths.Find(path)
	if p == nil {
		return nil, fmt.Errorf("path %s not found in request", path)
	}
	op := p.GetOperation(method)
	if op == nil {
		return nil, fmt.Errorf("method %s not found in request", method)
	}
	rb := op.RequestBody
	if rb == nil {
		return nil, fmt.Errorf("requestBody %d not found in request", status)
	}
	mime := rb.Value.Content.Get("application/json")
	if mime == nil {
		return nil, fmt.Errorf("mime application/json not found in request")
	}
	traverseSchema(mime.Schema.Value)

	// Clear slice between tests
	list := keys
	keys = nil
	isRequest = false

	return list, nil
}

// RecurseResponseKeys recursively traverses the Response object for a given OpenAPI path and returns a list of all properties
func RecurseResponseKeys(method string, path string, status int) ([]string, error) {
	p := doc.Paths.Find(path)
	if p == nil {
		return nil, fmt.Errorf("path %s not found in response", path)
	}
	op := p.GetOperation(method)
	if op == nil {
		return nil, fmt.Errorf("method %s not found in response", method)
	}
	st := op.Responses.Get(status)
	if st == nil {
		return nil, fmt.Errorf("status %d not found in response", status)
	}
	mime := st.Value.Content.Get("application/json")
	if mime == nil {
		return nil, fmt.Errorf("mime application/json not found in response")
	}
	traverseSchema(mime.Schema.Value)

	// Clear slice between tests
	list := keys
	keys = nil

	return list, nil
}

// Traverse openapi3.Schema
func traverseSchema(schema *openapi3.Schema) {
	if schema.AllOf != nil {
		traverseRefs(schema.AllOf)
	}

	if schema.Properties != nil {
		traverseProps(schema.Properties)
	}
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
	for k, v := range props {
		if k == "details" || k == "data" {
			traverseRef(v)
		}

		// Ignore keys on the Request object marked as read-only
		if isRequest && v.Value.ReadOnly {
			continue
		}

		// Ignore keys on the Response object marked as write-only
		if !isRequest && v.Value.WriteOnly {
			continue
		}

		// Note the key that was found
		if k != "details" && k != "data" {
			keys = append(keys, k)
		}

		// Check for oneOf
		if v.Value.OneOf != nil {
			traverseRefs(v.Value.OneOf)
		}

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
