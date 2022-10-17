package rize

import (
	"context"
	"net/http"
	"testing"
	"time"
)

// Complete Evaluation{} response data
var evaluation = &Evaluation{
	UID:       "EhrQZJNjCd79LLYq",
	Outcome:   "approved",
	CreatedAt: time.Now(),
	Flags: &EvaluationFlag{
		DocumentQualityCheck: true,
		FraudCheck:           true,
		FinancialCheck:       true,
		WatchListCheck:       true,
	},
	PIIMatch: &EvaluationPIIMatch{
		DOBMatch:     true,
		SSNMatch:     true,
		NameMatch:    true,
		EmailMatch:   true,
		PhoneMatch:   true,
		AddressMatch: true,
	},
}

func TestListEvaluations(t *testing.T) {
	params := &EvaluationListParams{
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

func TestGetEvaluation(t *testing.T) {
	resp, err := rc.Evaluations.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Evaluation\n", err)
	}

	if err := validateSchema(http.MethodGet, "/evaluations/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
