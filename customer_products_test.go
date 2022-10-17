package rize

import (
	"context"
	"net/http"
	"testing"
)

// Complete CustomerProduct{} response data
var customerProduct = &CustomerProduct{
	UID:           "Tegvs2E4TQgVYYMj",
	Status:        "active",
	CustomerUID:   "DuSzg6Ywr3cY9mw4",
	CustomerEmail: "olive.oyl@rizemoney.com",
	ProductUID:    "b8bemEKjAQhunbah",
	ProductName:   "Checking",
	ProgramUID:    "F1oFMKafpB2Zm6ng",
}

func TestListCustomerProducts(t *testing.T) {
	params := &CustomerProductListParams{
		ProgramUID:  "pQtTCSXz57fuefzp",
		ProductUID:  "zbJbEa72eKMgbbBv",
		CustomerUID: "uKxmLxUEiSj5h4M3",
	}
	resp, err := rc.CustomerProducts.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Customer Products\n", err)
	}

	if err := validateSchema(http.MethodGet, "/customer_products", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCreateCustomerProduct(t *testing.T) {
	params := &CustomerProductCreateParams{
		CustomerUID: "S62MaHx6WwsqG9vQ",
		ProductUID:  "pQtTCSXz57fuefzp",
	}
	resp, err := rc.CustomerProducts.Create(context.Background(), params)
	if err != nil {
		t.Fatal("Error creating Customer Product\n", err)
	}

	if err := validateSchema(http.MethodPost, "/customer_products", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetCustomerProduct(t *testing.T) {
	resp, err := rc.CustomerProducts.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Customer Product\n", err)
	}

	if err := validateSchema(http.MethodGet, "/customer_products/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
