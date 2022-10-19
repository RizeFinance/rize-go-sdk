package rize_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rizefinance/rize-go-sdk"
)

// Complete DebitCard{} response data
var debitCard = &rize.DebitCard{
	UID:                 "h9MzupcjtA3LPW2e",
	ExternalUID:         "9b463440-deff-4ff2-bfc9-1b71077a3b92",
	CustomerUID:         "uKxmLxUEiSj5h4M3",
	PoolUID:             "wTSMX1GubP21ev2h",
	SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
	CustodialAccountUID: "Gw4gr1T81YrvLT6M",
	CardLastFourDigits:  "9012",
	CardArtworkUID:      "EhrQZJNjCd79LLYq",
	IssuedOn:            "2019-10-21",
	Status:              "normal",
	Type:                "physical",
	ReadyToUse:          true,
	LockReason:          "lock reason",
	LockedAt:            time.Now(),
	ClosedAt:            time.Now(),
	LatestShippingAddress: &rize.DebitCardShippingAddress{
		Street1:    "123 Abc St.",
		Street2:    "Apt 2",
		City:       "Chicago",
		State:      "IL",
		PostalCode: "12345",
	},
}

var PINToken = &rize.DebitCardPINTokenResponse{
	PinChangeToken: "header.payload.signature",
}

var accessToken = &rize.DebitCardAccessToken{
	Token:    "VmU27goFku4DyxfsdyoH5G1mlztvwskBywKrskVN9jQOh50Yy7",
	ConfigID: "1",
}

func TestListDebitCards(t *testing.T) {
	params := &rize.DebitCardListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		ExternalUID: "client-generated-id",
		Limit:       100,
		Offset:      10,
		PoolUID:     "wTSMX1GubP21ev2h",
		Locked:      true,
		Status:      "queued",
	}

	resp, err := rc.DebitCards.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Debit Cards\n", err)
	}

	if err := validateSchema(http.MethodGet, "/debit_cards", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCreateDebitCard(t *testing.T) {
	params := &rize.DebitCardCreateParams{
		ExternalUID:    "partner-generated-id",
		CardArtworkUID: "EhrQZJNjCd79LLYq",
		CustomerUID:    "uKxmLxUEiSj5h4M3",
		PoolUID:        "wTSMX1GubP21ev2h",
		ShippingAddress: &rize.DebitCardShippingAddress{
			Street1:    "123 Abc St",
			Street2:    "Apt 2",
			City:       "Chicago",
			State:      "IL",
			PostalCode: "12345",
		},
	}
	resp, err := rc.DebitCards.Create(context.Background(), params)
	if err != nil {
		t.Fatal("Error creating Debit Card\n", err)
	}

	if err := validateSchema(http.MethodPost, "/debit_cards", http.StatusCreated, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetDebitCard(t *testing.T) {
	resp, err := rc.DebitCards.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Debit Card\n", err)
	}

	if err := validateSchema(http.MethodGet, "/debit_cards/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestActivateDebitCard(t *testing.T) {
	params := &rize.DebitCardActivateParams{
		CardLastFourDigits: "1234",
		CVV:                "012",
		ExpiryDate:         "2023-08",
	}
	resp, err := rc.DebitCards.Activate(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		t.Fatal("Error activating Debit Card\n", err)
	}

	if err := validateSchema(http.MethodPut, "/debit_cards/{uid}/activate", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestLockDebitCard(t *testing.T) {
	params := &rize.DebitCardLockParams{
		LockReason: "Fraud detected",
	}
	resp, err := rc.DebitCards.Lock(context.Background(), "Lt6qjTNnYLjFfEWL", params)
	if err != nil {
		t.Fatal("Error locking Debit Card\n", err)
	}

	if err := validateSchema(http.MethodPut, "/debit_cards/{uid}/lock", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestUnlockDebitCard(t *testing.T) {
	resp, err := rc.DebitCards.Unlock(context.Background(), "Lt6qjTNnYLjFfEWL")
	if err != nil {
		t.Fatal("Error unlocking Debit Card\n", err)
	}

	if err := validateSchema(http.MethodPut, "/debit_cards/{uid}/unlock", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestReissueDebitCard(t *testing.T) {
	params := &rize.DebitCardReissueParams{
		CardArtworkUID: "EhrQZJNjCd79LLYq",
		ReissueReason:  "damaged",
		ShippingAddress: &rize.DebitCardShippingAddress{
			Street1:    "123 Abc St",
			Street2:    "Apt 2",
			City:       "Chicago",
			State:      "IL",
			PostalCode: "12345",
		},
	}
	resp, err := rc.DebitCards.Reissue(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		t.Fatal("Error reissuing Debit Card\n", err)
	}

	if err := validateSchema(http.MethodPut, "/debit_cards/{uid}/reissue", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetPINToken(t *testing.T) {
	params := &rize.DebitCardGetPINTokenParams{
		ForceReset: true,
	}
	resp, err := rc.DebitCards.GetPINToken(context.Background(), "Lt6qjTNnYLjFfEWL", params)
	if err != nil {
		t.Fatal("Error fetching Debit Card PIN Token\n", err)
	}

	if err := validateSchema(http.MethodGet, "/debit_cards/{uid}/pin_change_token", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetAccessToken(t *testing.T) {
	resp, err := rc.DebitCards.GetAccessToken(context.Background(), "Lt6qjTNnYLjFfEWL")
	if err != nil {
		t.Fatal("Error fetching Debit Card Access Token\n", err)
	}

	if err := validateSchema(http.MethodGet, "/debit_cards/{uid}/access_token", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMigrateVirtualDebitCard(t *testing.T) {
	params := &rize.VirtualDebitCardMigrateParams{
		ExternalUID:    "partner-generated-id",
		CardArtworkUID: "EhrQZJNjCd79LLYq",
		ShippingAddress: &rize.DebitCardShippingAddress{
			Street1:    "123 Abc St",
			Street2:    "Apt 2",
			City:       "Chicago",
			State:      "IL",
			PostalCode: "12345",
		},
	}
	resp, err := rc.DebitCards.MigrateVirtualDebitCard(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		t.Fatal("Error migrating Debit Card\n", err)
	}
	if err := validateSchema(http.MethodPut, "/debit_cards/{uid}/migrate", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetVirtualDebitCardImage(t *testing.T) {
	params := &rize.DebitCardAccessToken{
		Token:    "VmU27goFku4DyxfsdyoH5G1mlztvwskBywKrskVN9jQOh50Yy7",
		ConfigID: "1",
	}
	_, err := rc.DebitCards.GetVirtualDebitCardImage(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Virtual Debit Card Image\n", err)
	}

	if err := validateSchema(http.MethodGet, "/assets/virtual_card_image", http.StatusOK, params, nil, nil); err != nil {
		t.Fatalf(err.Error())
	}
}
