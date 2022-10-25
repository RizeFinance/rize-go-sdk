package rize

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Handles all Product operations
type productService service

// Product data type
type Product struct {
	UID                      string                `json:"uid,omitempty"`
	Name                     string                `json:"name,omitempty"`
	Description              string                `json:"description,omitempty"`
	ProductCompliancePlanUID string                `json:"product_compliance_plan_uid,omitempty"`
	CompliancePlanName       string                `json:"compliance_plan_name,omitempty"`
	CustomerTypes            []string              `json:"customer_types,omitempty"`
	PrerequisiteProductUIDs  []string              `json:"prerequisite_product_uids,omitempty"`
	ProgramUID               string                `json:"program_uid,omitempty"`
	ProfileRequirements      []*ProfileRequirement `json:"profile_requirements,omitempty"`
}

// ProfileRequirement is a list of Profile Requirements a Customer must provide Profile Responses to
type ProfileRequirement struct {
	ProfileRequirementUID string   `json:"profile_requirement_uid,omitempty"`
	ProfileRequirement    string   `json:"profile_requirement,omitempty"`
	Category              string   `json:"category,omitempty"`
	Required              bool     `json:"required,omitempty"`
	RequirementType       string   `json:"requirement_type,omitempty"`
	ResponseValues        []string `json:"response_values,omitempty"`
}

// ProductListParams builds the query parameters used in querying Products
type ProductListParams struct {
	ProgramUID string `url:"program_uid,omitempty" json:"program_uid,omitempty"`
}

// ProductResponse is an API response containing a list of Products
type ProductResponse struct {
	BaseResponse
	Data []*Product `json:"data"`
}

// List retrieves a list of Products filtered by the given parameters
func (p *productService) List(ctx context.Context, params *ProductListParams) (*ProductResponse, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	res, err := p.client.doRequest(ctx, http.MethodGet, "products", v, nil)
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

// Get returns a single Product
func (p *productService) Get(ctx context.Context, uid string) (*Product, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := p.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("products/%s", uid), nil, nil)
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
