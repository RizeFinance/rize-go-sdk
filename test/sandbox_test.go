package rize_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/rizefinance/rize-go-sdk"
)

func TestCreateSandbox(t *testing.T) {
	params := &rize.SandboxCreateParams{
		TransactionType:  "atm_withdrawal",
		CustomerUID:      "uKxmLxUEiSj5h4M3",
		DebitCardUID:     "h9MzupcjtA3LPW2e",
		DenialReason:     "insufficient_funds",
		USDollarAmount:   21.89,
		Mcc:              "5200",
		MerchantLocation: "NEW YORK, NY",
		MerchantName:     "Widgets Incorporated",
		MerchantNumber:   "000067107015968",
		Description:      "test transaction",
	}
	resp, err := rc.Sandbox.Create(context.Background(), params)
	if err != nil {
		t.Fatal("Error creating Sandbox transactions\n", err)
	}

	if err := validateSchema(http.MethodPost, "/sandbox/mock_transactions", http.StatusCreated, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
