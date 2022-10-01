package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/rizefinance/rize-go-sdk/internal"
	rize "github.com/rizefinance/rize-go-sdk/platform"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	config := rize.RizeConfig{
		ProgramUID:  internal.CheckEnvVariable("program_uid"),
		HMACKey:     internal.CheckEnvVariable("hmac_key"),
		Environment: internal.CheckEnvVariable("environment"),
		Debug:       true,
	}

	// Create new Rize client
	rc, err := rize.NewRizeClient(&config)
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
	log.Printf("%+v", l)

	// Create workflow
	wcp := rize.WorkflowCreateParams{
		CustomerUID:              "h9MzupcjtA3LPW2e",
		ProductCompliancePlanUID: "25NQX3GGXpAtpUmP",
	}
	c, err := rc.ComplianceWorkflows.Create(context.Background(), &wcp)
	if err != nil {
		log.Fatal("Error creating new compliance workflow\n", err)
	}
	log.Printf("%+v", c)

	// View latest workflow for a customer
	lp := rize.WorkflowLatestParams{
		ProductCompliancePlanUID: "pQtTCSXz57fuefzp",
	}
	cw, err := rc.ComplianceWorkflows.ViewLatest(context.Background(), "h9MzupcjtA3LPW2e", &lp)
	if err != nil {
		log.Fatal("Error fetching latest customer workflow\n", err)
	}
	log.Printf("%+v", cw)

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
	log.Printf("%+v", ad)

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
	log.Printf("%+v", ads)
}
