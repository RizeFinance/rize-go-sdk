package rize_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rizefinance/rize-go-sdk"
)

// Complete Document{} response data
var document = &rize.Document{
	UID:                  "EhrQZJNjCd79LLYq",
	DocumentType:         "monthly_statement",
	ScopeType:            "customer",
	Name:                 "uKxmLxUEiSj5h4M3-November-2021",
	PeriodStartedAt:      time.Now(),
	PeriodEndedAt:        time.Now(),
	CreatedAt:            time.Now(),
	CustomerUIDs:         []string{"uKxmLxUEiSj5h4M3, EhrQZJNjCd79LLYq"},
	CustodialAccountUIDs: []string{"oib3bfMxwyuYLDKE, vwWQfAz1aY9VDHg4"},
	SyntheticAccountUIDs: []string{"Qtjp9FdPPpTiFeeo, NYeFXLPzNXbJQ2WR"},
}

func TestDocumentService_List(t *testing.T) {
	params := &rize.DocumentListParams{
		DocumentType:        "monthly_statement",
		Month:               1,
		Year:                2020,
		CustodialAccountUID: "yqyYk5b1xgXFFrXs",
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
		Limit:               100,
		Offset:              10,
	}
	resp, err := rc.Documents.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Documents\n", err)
	}

	if err := validateSchema(http.MethodGet, "/documents", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestDocumentService_Get(t *testing.T) {
	resp, err := rc.Documents.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Document\n", err)
	}

	if err := validateSchema(http.MethodGet, "/documents/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestDocumentService_View(t *testing.T) {
	_, err := rc.Documents.View(context.Background(), "u8EHFJnWvJxRJZxa")
	if err != nil {
		t.Fatal("Error viewing document\n", err)
	}

	if err := validateSchema(http.MethodGet, "/documents/{uid}/view", http.StatusOK, nil, nil, nil); err != nil {
		t.Fatalf(err.Error())
	}
}
