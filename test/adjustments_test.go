package rize_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rizefinance/rize-go-sdk"
)

// Complete Adjustment{} response data
var adjustment = &rize.Adjustment{
	UID:                 "EhrQZJNjCd79LLYq",
	ExternalUID:         "PT3sH7oxxQPwchrf",
	CustomerUID:         "uKxmLxUEiSj5h4M3",
	USDAdjustmentAmount: "2.43",
	CreatedAt:           time.Now(),
	Status:              "initiated",
	AdjustmentType: &rize.AdjustmentType{
		UID:         "KM2eKbR98t4tdAyZ",
		Name:        "weekly_membership",
		Description: "Weekly membership fee",
		Fee:         true,
		Deprecated:  true,
	},
}

// Complete AdjustmentType{} response data
var adjustmentType = &rize.AdjustmentType{
	UID:         "EhrQZJNjCd79LLYq",
	Name:        "monthly_membership",
	Description: "Monthly membership fee",
	Fee:         true,
	ProgramUID:  "DbxJUHVuqt3C7hGK",
	Deprecated:  true,
}

func TestAdjustmentService_List(t *testing.T) {
	params := &rize.AdjustmentListParams{
		CustomerUID:            "uKxmLxUEiSj5h4M3",
		AdjustmentTypeUID:      "2Ej2tsFbQT3S1HYd",
		ExternalUID:            "PT3sH7oxxQPwchrf",
		USDAdjustmentAmountMax: 10,
		USDAdjustmentAmountMin: 5,
		Sort:                   "adjustment_type_name_asc",
	}

	resp, err := rc.Adjustments.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Adjustments\n", err)
	}

	if err := validateSchema(http.MethodGet, "/adjustments", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAdjustmentService_Create(t *testing.T) {
	params := &rize.AdjustmentCreateParams{
		ExternalUID:         "partner-generated-id",
		CustomerUID:         "kaxHFJnWvJxRJZxq",
		USDAdjustmentAmount: "2.43",
		AdjustmentTypeUID:   "KM2eKbR98t4tdAyZ",
	}

	resp, err := rc.Adjustments.Create(context.Background(), params)
	if err != nil {
		t.Fatal("Error creating Adjustment\n", err)
	}

	if err := validateSchema(http.MethodPost, "/adjustments", http.StatusCreated, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAdjustmentService_Get(t *testing.T) {
	resp, err := rc.Adjustments.Get(context.Background(), "exMDShw6yM3NHLYV")
	if err != nil {
		t.Fatal("Error fetching Adjustment\n", err)
	}

	if err := validateSchema(http.MethodGet, "/adjustments/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAdjustmentService_ListAdjustmentTypes(t *testing.T) {
	params := &rize.AdjustmentTypeListParams{
		CustomerUID:    "uKxmLxUEiSj5h4M3",
		ProgramUID:     "DbxJUHVuqt3C7hGK",
		ShowDeprecated: true,
	}

	resp, err := rc.Adjustments.ListAdjustmentTypes(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Adjustment Types\n", err)
	}

	if err := validateSchema(http.MethodGet, "/adjustment_types", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAdjustmentService_GetAdjustmentType(t *testing.T) {
	resp, err := rc.Adjustments.GetAdjustmentType(context.Background(), "exMDShw6yM3NHLYV")
	if err != nil {
		t.Fatal("Error fetching Adjustment Type\n", err)
	}

	if err := validateSchema(http.MethodGet, "/adjustment_types/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}

}
