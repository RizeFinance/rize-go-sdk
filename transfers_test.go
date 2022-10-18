package rize

import (
	"context"
	"net/http"
	"testing"
	"time"
)

// Complete Transfer{} response data
var transfer = &Transfer{
	UID:                            "EhrQZJNjCd79LLYq",
	ExternalUID:                    "partner-generated-id",
	SourceSyntheticAccountUID:      "4XkJnsfHsuqrxmeX",
	DestinationSyntheticAccountUID: "exMDShw6yM3NHLYV",
	InitiatingCustomerUID:          "iDtmSA52zRhgN4iy",
	Status:                         "pending",
	CreatedAt:                      time.Now(),
	USDTransferAmount:              "34.12",
	USDRequestedAmount:             "12.34",
}

func TestListTransfers(t *testing.T) {
	params := &TransferListParams{
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		ExternalUID:         "client-generated-id",
		PoolUID:             "wTSMX1GubP21ev2h",
		SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
		Limit:               100,
		Offset:              10,
	}
	resp, err := rc.Transfers.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Transfers\n", err)
	}

	if err := validateSchema(http.MethodGet, "/transfers", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCreateTransfer(t *testing.T) {
	params := &TransferCreateParams{
		ExternalUID:                    "partner-generated-id",
		SourceSyntheticAccountUID:      "4XkJnsfHsuqrxmeX",
		DestinationSyntheticAccountUID: "exMDShw6yM3NHLYV",
		InitiatingCustomerUID:          "iDtmSA52zRhgN4iy",
		USDTransferAmount:              "12.34",
	}
	resp, err := rc.Transfers.Create(context.Background(), params)
	if err != nil {
		t.Fatal("Error creating Transfer\n", err)
	}

	if err := validateSchema(http.MethodPost, "/transfers", http.StatusCreated, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetTransfer(t *testing.T) {
	resp, err := rc.Transfers.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Transfer\n", err)
	}

	if err := validateSchema(http.MethodGet, "/transfers/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
