package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List workflows
func (e Example) ExampleComplianceWorkflowService_List(rc *rize.Client) {
	params := &rize.WorkflowListParams{
		CustomerUID: "S62MaHx6WwsqG9vQ",
		ProductUID:  "pQtTCSXz57fuefzp",
		InProgress:  true,
		Limit:       100,
		Offset:      10,
	}
	resp, err := rc.ComplianceWorkflows.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Compliance Workflows\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Workflows", string(output))
}

// Create workflow
func (e Example) ExampleComplianceWorkflowService_Create(rc *rize.Client) {
	params := &rize.WorkflowCreateParams{
		CustomerUID:              "h9MzupcjtA3LPW2e",
		ProductCompliancePlanUID: "25NQX3GGXpAtpUmP",
	}
	resp, err := rc.ComplianceWorkflows.Create(context.Background(), params)
	if err != nil {
		log.Fatal("Error creating Compliance Workflow\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Create Workflow", string(output))
}

// View latest workflow for a customer
func (e Example) ExampleComplianceWorkflowService_ViewLatest(rc *rize.Client) {
	params := &rize.WorkflowLatestParams{
		ProductCompliancePlanUID: "pQtTCSXz57fuefzp",
	}
	resp, err := rc.ComplianceWorkflows.ViewLatest(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		log.Fatal("Error fetching Compliance Workflow\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("View Latest Workflow", string(output))
}

// Acknowledge document
func (e Example) ExampleComplianceWorkflowService_AcknowledgeDocument(rc *rize.Client) {
	params := &rize.WorkflowDocumentParams{
		Accept:      "yes",
		CustomerUID: "h9MzupcjtA3LPW2e",
		DocumentUID: "Yqyjk5b2xgQ9FrxS",
		IPAddress:   "107.56.230.156",
		UserName:    "gilbert chesterton",
	}
	resp, err := rc.ComplianceWorkflows.AcknowledgeDocument(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		log.Fatal("Error acknowledging compliance document\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Acknowledge Document", string(output))
}

// Acknowledge multiple documents
func (e Example) ExampleComplianceWorkflowService_BatchAcknowledgeDocuments(rc *rize.Client) {
	params := &rize.WorkflowBatchDocumentsParams{
		CustomerUID: "h9MzupcjtA3LPW2e",
		Documents: []*rize.WorkflowDocumentParams{{
			Accept:      "yes",
			DocumentUID: "Yqyjk5b2xgQ9FrxS",
			IPAddress:   "107.56.230.156",
			UserName:    "gilbert chesterton",
		}, {
			Accept:      "yes",
			DocumentUID: "BgT64WeR0IxkgH6D",
		}},
	}
	resp, err := rc.ComplianceWorkflows.BatchAcknowledgeDocuments(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		log.Fatal("Error acknowledging compliance document\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Acknowledge Document", string(output))
}
