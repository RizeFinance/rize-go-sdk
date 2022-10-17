package rize

import (
	"context"
	"net/http"
	"testing"
)

// Complete CustodialPartner{} response data
var custodialPartner = &CustodialPartner{
	UID:  "EhrQZJNjCd79LLYq",
	Name: "Big Apple Bank",
	Type: "Mutual savings bank",
}

func TestListCustodialPartners(t *testing.T) {
	resp, err := rc.CustodialPartners.List(context.Background())
	if err != nil {
		t.Fatal("Error fetching Custodial Partners\n", err)
	}

	if err := validateSchema(http.MethodGet, "/custodial_partners", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetCustodialPartner(t *testing.T) {
	resp, err := rc.CustodialPartners.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Custodial Partner\n", err)
	}

	if err := validateSchema(http.MethodGet, "/custodial_partners/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}