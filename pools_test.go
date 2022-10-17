package rize

import (
	"context"
	"net/http"
	"testing"
)

// Complete Pool{} response data
var pool = &Pool{
	UID:              "wTSMX1GubP21ev2h",
	Name:             "multi-byte success",
	OwnerCustomerUID: "uKxmLxUEiSj5h4M3",
	CustomerUIDs:     []string{"uKxmLxUEiSj5h4M3", "ivxTYrJtwrMC4f6w"},
}

func TestListPools(t *testing.T) {
	params := &PoolListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		ExternalUID: "client-generated-id",
		Limit:       100,
		Offset:      10,
	}
	resp, err := rc.Pools.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching pools\n", err)
	}

	if err := validateSchema(http.MethodGet, "/pools", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetPool(t *testing.T) {
	resp, err := rc.Pools.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching pool\n", err)
	}

	if err := validateSchema(http.MethodGet, "/pools/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
