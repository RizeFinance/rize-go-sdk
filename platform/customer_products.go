package platform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Handles all Customer Product operations
type customerProductService service

// CustomerProduct data type
type CustomerProduct struct {
	UID           string `json:"uid,omitempty"`
	Status        string `json:"status,omitempty"`
	CustomerUID   string `json:"customer_uid,omitempty"`
	CustomerEmail string `json:"customer_email,omitempty"`
	ProductUID    string `json:"product_uid,omitempty"`
	ProductName   string `json:"product_name,omitempty"`
	ProgramUID    string `json:"program_uid,omitempty"`
}

// CustomerProductListParams builds the query parameters used in querying Customer Products
type CustomerProductListParams struct {
	ProgramUID  string `url:"program_uid,omitempty"`
	ProductUID  string `url:"product_uid,omitempty"`
	CustomerUID string `url:"customer_uid,omitempty"`
}

// CustomerProductCreateParams are the body params used when creating a new Customer Product
type CustomerProductCreateParams struct {
	CustomerUID string `json:"customer_uid"`
	ProductUID  string `json:"product_uid"`
}

// CustomerProductResponse is an API response containing a list of Customer Products
type CustomerProductResponse struct {
	BaseResponse
	Data []*CustomerProduct `json:"data"`
}

// List Customers and the Products they have onboarded onto, filtered by the given parameters
func (cp *customerProductService) List(cpp *CustomerProductListParams) (*CustomerProductResponse, error) {
	v, err := query.Values(cpp)
	if err != nil {
		return nil, err
	}

	res, err := cp.client.doRequest(http.MethodGet, "customer_products", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CustomerProductResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Create will submit a request to onboard a Customer onto a new product
func (cp *customerProductService) Create(ccp *CustomerProductCreateParams) (*CustomerProduct, error) {
	if ccp.CustomerUID == "" || ccp.ProductUID == "" {
		return nil, fmt.Errorf("CustomerUID and ProductUID are required")
	}

	bytesMessage, err := json.Marshal(ccp)
	if err != nil {
		return nil, err
	}

	res, err := cp.client.doRequest(http.MethodPost, "customer_products", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CustomerProduct{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get a single Customer Product
func (cp *customerProductService) Get(uid string) (*CustomerProduct, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := cp.client.doRequest(http.MethodGet, fmt.Sprintf("customer_products/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CustomerProduct{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
