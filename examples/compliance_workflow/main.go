package main

import (
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
	wq := rize.WorkflowQuery{
		Limit: 10,
	}
	w, err := rc.ComplianceWorkflow.ListWorkflows(&wq)
	if err != nil {
		log.Fatal("Error fetching compliance workflows\n", err)
	}
	log.Printf("%+v", w)

	// Create workflow

	// View workflow for customer
	cwq := rize.CustomerWorkflowQuery{
		ProductCompliancePlanUID: "mheiDmW1K2LSMZQU",
	}
	cw, err := rc.ComplianceWorkflow.ViewCustomerWorkflow("SPbiwv93C6M5pSWu", &cwq)
	if err != nil {
		log.Fatal("Error fetching customer workflows\n", err)
	}
	log.Printf("%+v", cw)

	// Acknowledge document
	// Acknowledge multiple documents

}
