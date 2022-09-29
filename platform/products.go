package platform

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Handles all Product operations
type productService service

// Product data type
type Product struct {
	UID                      string   `json:"uid"`
	Name                     string   `json:"name"`
	Description              string   `json:"description"`
	ProductCompliancePlanUID string   `json:"product_compliance_plan_uid"`
	CompliancePlanName       string   `json:"compliance_plan_name"`
	PrerequisiteProductUids  []string `json:"prerequisite_product_uids"`
	ProgramUID               string   `json:"program_uid"`
	ProfileRequirements      []struct {
		ProfileRequirementUID string   `json:"profile_requirement_uid"`
		ProfileRequirement    string   `json:"profile_requirement"`
		Category              string   `json:"category"`
		Required              bool     `json:"required"`
		RequirementType       string   `json:"requirement_type"`
		ResponseValues        []string `json:"response_values"`
	} `json:"profile_requirements"`
}

// ProductResponse is an API response containing a list of Products
type ProductResponse struct {
	BaseResponse
	Data []*Product `json:"data"`
}

// List retrieves a list of Products filtered by the given parameters
func (p *productService) List(programUID string) (*ProductResponse, error) {
	v := url.Values{}
	v.Set("program_uid", programUID)

	res, err := p.rizeClient.doRequest(http.MethodGet, "products", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &ProductResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single product
func (p *productService) Get(uid string) (*Product, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := p.rizeClient.doRequest(http.MethodGet, fmt.Sprintf("products/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Product{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
