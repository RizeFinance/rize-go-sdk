package rize_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rizefinance/rize-go-sdk"
)

// Complete PinwheelJob{} response data
var pinwheelJob = &rize.PinwheelJob{
	UID:                  "EhrQZJNjCd79LLYq",
	SyntheticAccountUID:  "Jy8degj6iv2QngLo",
	Status:               "initiated",
	CreatedAt:            time.Now(),
	StatusUpdatedAt:      time.Now(),
	CustomerUID:          "h9MzupcjtA3LPW2e",
	LinkToken:            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	ExpiresAt:            time.Now(),
	JobNames:             []string{"direct_deposit_switch"},
	Amount:               1000,
	DisablePartialSwitch: false,
	OrganizationName:     "Chipotle Mexican Grill, Inc.",
	SkipWelcomeScreen:    false,
}

func TestPinwheelJobService_List(t *testing.T) {
	params := &rize.PinwheelJobListParams{
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
		Limit:               100,
		Offset:              10,
	}
	resp, err := rc.PinwheelJobs.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Pinwheel Jobs\n", err)
	}

	if err := validateSchema(http.MethodGet, "/pinwheel_jobs", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestPinwheelJobService_Create(t *testing.T) {
	params := &rize.PinwheelJobCreateParams{
		JobNames:             []string{"direct_deposit_switch"},
		SyntheticAccountUID:  "4XkJnsfHsuqrxmeX",
		Amount:               1000,
		DisablePartialSwitch: true,
		OrganizationName:     "Chipotle Mexican Grill, Inc.",
		SkipWelcomeScreen:    true,
	}
	resp, err := rc.PinwheelJobs.Create(context.Background(), params)
	if err != nil {
		t.Fatal("Error creating Pinwheel Job\n", err)
	}

	if err := validateSchema(http.MethodPost, "/pinwheel_jobs", http.StatusCreated, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestPinwheelJobService_Get(t *testing.T) {
	resp, err := rc.PinwheelJobs.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Pinwheel Job\n", err)
	}

	if err := validateSchema(http.MethodGet, "/pinwheel_jobs/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
