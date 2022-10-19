package rize_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rizefinance/rize-go-sdk"
)

// Complete Evaluation{} response data
var evaluation = &rize.Evaluation{
	UID:       "EhrQZJNjCd79LLYq",
	Outcome:   "approved",
	CreatedAt: time.Now(),
	Flags: &rize.EvaluationFlag{
		DocumentQualityCheck: true,
		FraudCheck:           true,
		FinancialCheck:       true,
		WatchListCheck:       true,
	},
	PIIMatch: &rize.EvaluationPIIMatch{
		DOBMatch:     true,
		SSNMatch:     true,
		NameMatch:    true,
		EmailMatch:   true,
		PhoneMatch:   true,
		AddressMatch: true,
	},
}

func TestEvaluationService_List(t *testing.T) {
	params := &rize.EvaluationListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		Latest:      true,
	}
	resp, err := rc.Evaluations.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Evaluations\n", err)
	}

	if err := validateSchema(http.MethodGet, "/evaluations", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestEvaluationService_Get(t *testing.T) {
	resp, err := rc.Evaluations.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Evaluation\n", err)
	}

	if err := validateSchema(http.MethodGet, "/evaluations/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
