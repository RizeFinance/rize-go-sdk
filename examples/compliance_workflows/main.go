package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
	"github.com/rizefinance/rize-go-sdk"
	"github.com/rizefinance/rize-go-sdk/internal"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	config := rize.Config{
		ProgramUID:  internal.CheckEnvVariable("program_uid"),
		HMACKey:     internal.CheckEnvVariable("hmac_key"),
		Environment: internal.CheckEnvVariable("environment"),
		Debug:       true,
	}

	// Create new Rize client
	rc, err := rize.NewClient(&config)
	if err != nil {
		log.Fatal("Error building RizeClient\n", err)
	}

	// List workflows
	wlp := rize.WorkflowListParams{
		CustomerUID: "",
		ProductUID:  "",
		InProgress:  false,
		Limit:       10,
		Offset:      0,
	}
	l, err := rc.ComplianceWorkflows.List(context.Background(), &wlp)
	if err != nil {
		log.Fatal("Error fetching compliance workflows\n", err)
	}
	output, _ := json.MarshalIndent(l, "", "\t")
	log.Println("List Workflows", string(output))

	// Create workflow
	wcp := rize.WorkflowCreateParams{
		CustomerUID:              "h9MzupcjtA3LPW2e",
		ProductCompliancePlanUID: "25NQX3GGXpAtpUmP",
	}
	c, err := rc.ComplianceWorkflows.Create(context.Background(), &wcp)
	if err != nil {
		log.Fatal("Error creating new compliance workflow\n", err)
	}
	output, _ = json.MarshalIndent(c, "", "\t")
	log.Println("Create Workflow", string(output))

	// View latest workflow for a customer
	lp := rize.WorkflowLatestParams{
		ProductCompliancePlanUID: "pQtTCSXz57fuefzp",
	}
	cw, err := rc.ComplianceWorkflows.ViewLatest(context.Background(), "h9MzupcjtA3LPW2e", &lp)
	if err != nil {
		log.Fatal("Error fetching latest customer workflow\n", err)
	}
	output, _ = json.MarshalIndent(cw, "", "\t")
	log.Println("View Latest Workflow", string(output))

	// Acknowledge document
	wd := rize.WorkflowDocumentParams{
		Accept:      "yes",
		CustomerUID: "h9MzupcjtA3LPW2e",
		DocumentUID: "Yqyjk5b2xgQ9FrxS",
		IPAddress:   "107.56.230.156",
		UserName:    "gilbert chesterton",
	}
	ad, err := rc.ComplianceWorkflows.AcknowledgeDocument(context.Background(), "dolordo", &wd)
	if err != nil {
		log.Fatal("Error acknowledging compliance document\n", err)
	}
	output, _ = json.MarshalIndent(ad, "", "\t")
	log.Println("Acknowledge Document", string(output))

	// Acknowledge multiple documents
	wdp := rize.WorkflowDocumentsParams{
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
	ads, err := rc.ComplianceWorkflows.AcknowledgeDocuments(context.Background(), "dolordo", &wdp)
	if err != nil {
		log.Fatal("Error acknowledging compliance documents\n", err)
	}
	output, _ = json.MarshalIndent(ads, "", "\t")
	log.Println("Acknowledge Document", string(output))
}
