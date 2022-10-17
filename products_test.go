package rize

import (
	"context"
	"net/http"
	"testing"
)

// Complete Product{} response data
var product = &Product{
	UID:                      "f9VncZny4ejhcPF4",
	Name:                     "checking",
	Description:              "Supports checking accounts",
	ProductCompliancePlanUID: "Vxf1pY1KHbiJYo1e",
	CompliancePlanName:       "Compliance requirements for checking",
	CustomerTypes:            []string{"primary", "sole_proprietor"},
	PrerequisiteProductUIDs:  []string{"DaB7Mjj73Nz2JpHF", "x2z691J9HPCWAugv"},
	ProgramUID:               "W74Jrkxk8bVtvNNj",
	ProfileRequirements: []*ProfileRequirement{{
		ProfileRequirementUID: "ptRLF7nQvy8VoqM1",
		ProfileRequirement:    "Please provide your approximate annual income in USD.",
		Category:              "default",
		Required:              true,
		RequirementType:       "fixed_list",
		ResponseValues:        []string{"yes", "no"},
	}},
}

func TestListProducts(t *testing.T) {
	params := &ProductListParams{
		ProgramUID: "pQtTCSXz57fuefzp",
	}
	resp, err := rc.Products.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Products\n", err)
	}

	if err := validateSchema(http.MethodGet, "/products", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetProduct(t *testing.T) {
	resp, err := rc.Products.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Product\n", err)
	}

	if err := validateSchema(http.MethodGet, "/products/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
