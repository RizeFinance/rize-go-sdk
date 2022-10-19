package rize_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rizefinance/rize-go-sdk"
)

// Complete Workflow{} response data
var workflow = &rize.Workflow{
	UID: "SPbiwv93C6M5pSWu",
	Summary: &rize.WorkflowSummary{
		AcceptedQuantity: 1,
		BegunAt:          time.Now(),
		CompletedStep:    1,
		CurrentStep:      1,
		Status:           "in_progress",
	},
	Customer: &rize.WorkflowCustomer{
		Email:       "tomas@example.com",
		ExternalUID: "client-generated-42",
		UID:         "h9MzupcjtA3LPW2e",
	},
	ProductUID:               "ncGN746ya7JLzKu3",
	ProductCompliancePlanUID: "mheiDmW1K2LSMZQU",
	AcceptedDocuments: []*rize.WorkflowAcceptedDocument{{
		ElectronicSignatureRequired: "no",
		ExternalStorageName:         "usa_ptrt_0",
		ComplianceDocumentURL:       "https://document-bucket-example.s3.amazonaws.com/RW2HF123FyVMUxNw/8ABcG12MYgAsomZm",
		Name:                        "USA PATRIOT Act",
		Step:                        1,
		Version:                     1,
		UID:                         "vU7f8jrLCww7ev2H",
		AcceptedAt:                  time.Now(),
	}},
	CurrentStepDocumentsPending: []*rize.WorkflowPendingDocument{{
		ElectronicSignatureRequired: "no",
		ExternalStorageName:         "eft_auth_0",
		ComplianceDocumentURL:       "https://document-bucket-example.s3.amazonaws.com/RW2HF123FyVMUxNw/123cG1OMggAso4B2",
		Name:                        "EFT Authorization Agreement",
		Step:                        1,
		Version:                     1,
		UID:                         "dc6PApa2nn9K3jwL",
	}},
	AllDocuments: []*rize.WorkflowDocument{{
		ElectronicSignatureRequired: "no",
		ExternalStorageName:         "eft_auth_0",
		ComplianceDocumentURL:       "https://document-bucket-example.s3.amazonaws.com/RW2HF123FyVMUxNw/123cG1OMggAso4B2",
		Name:                        "EFT Authorization Agreement",
		Step:                        1,
		Version:                     1,
	}},
}

func TestListWorkflows(t *testing.T) {
	params := &rize.WorkflowListParams{
		CustomerUID: "S62MaHx6WwsqG9vQ",
		ProductUID:  "pQtTCSXz57fuefzp",
		InProgress:  true,
		Limit:       100,
		Offset:      10,
	}
	resp, err := rc.ComplianceWorkflows.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Compliance Workflows\n", err)
	}

	if err := validateSchema(http.MethodGet, "/compliance_workflows", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCreateWorkflow(t *testing.T) {
	params := &rize.WorkflowCreateParams{
		CustomerUID:              "h9MzupcjtA3LPW2e",
		ProductCompliancePlanUID: "25NQX3GGXpAtpUmP",
	}
	resp, err := rc.ComplianceWorkflows.Create(context.Background(), params)
	if err != nil {
		t.Fatal("Error creating Compliance Workflow\n", err)
	}

	if err := validateSchema(http.MethodPost, "/compliance_workflows", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestViewLatestWorkflow(t *testing.T) {
	params := &rize.WorkflowLatestParams{
		ProductCompliancePlanUID: "pQtTCSXz57fuefzp",
	}
	resp, err := rc.ComplianceWorkflows.ViewLatest(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		t.Fatal("Error fetching Compliance Workflow\n", err)
	}

	if err := validateSchema(http.MethodGet, "/compliance_workflows/latest/{customer_uid}", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAcknowledgeDocument(t *testing.T) {
	params := &rize.WorkflowDocumentParams{
		Accept:      "yes",
		CustomerUID: "h9MzupcjtA3LPW2e",
		DocumentUID: "Yqyjk5b2xgQ9FrxS",
		IPAddress:   "107.56.230.156",
		UserName:    "gilbert chesterton",
	}
	resp, err := rc.ComplianceWorkflows.AcknowledgeDocument(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		t.Fatal("Error acknowledging compliance document\n", err)
	}

	if err := validateSchema(http.MethodPut, "/compliance_workflows/{uid}/acknowledge_document", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestBatchAcknowledgeDocuments(t *testing.T) {
	params := &rize.WorkflowBatchDocumentsParams{
		CustomerUID: "h9MzupcjtA3LPW2e",
		Documents: []*rize.WorkflowDocumentParams{{
			Accept:      "yes",
			DocumentUID: "Yqyjk5b2xgQ9FrxS",
			IPAddress:   "107.56.230.156",
			UserName:    "gilbert chesterton",
			CustomerUID: "h9MzupcjtA3LPW2e",
		}, {
			Accept:      "yes",
			DocumentUID: "BgT64WeR0IxkgH6D",
			CustomerUID: "h9MzupcjtA3LPW2e",
		}},
	}
	resp, err := rc.ComplianceWorkflows.BatchAcknowledgeDocuments(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		t.Fatal("Error acknowledging compliance document\n", err)
	}

	if err := validateSchema(http.MethodPut, "/compliance_workflows/{uid}/batch_acknowledge_documents", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
