package platform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

// ComplianceWorkflowService handles all Compliance Workflow operations
type ComplianceWorkflowService service

// WorkflowQuery builds the parameters used in querying compliance workflows
type WorkflowQuery struct {
	CustomerUID string `url:"customer_uid,omitempty"`
	ProductUID  string `url:"product_uid,omitempty"`
	InProgress  bool   `url:"in_progress,omitempty"`
	Limit       int    `url:"limit,omitempty"`
	Offset      int    `url:"offset,omitempty"`
}

// CustomerWorkflowQuery builds parameters used in querying workflows for a customer
type CustomerWorkflowQuery struct {
	ProductCompliancePlanUID string `url:"product_compliance_plan_uid,omitempty"`
}

// Workflows is an API response containing a list of compliance workflows
type Workflows struct {
	TotalCount int           `json:"total_count"`
	Count      int           `json:"count"`
	Limit      int           `json:"limit"`
	Offset     int           `json:"offset"`
	Data       []interface{} `json:"data"`
}

// ListWorkflows retrieves a list of Compliance Workflows filtered by the given parameters
func (c *ComplianceWorkflowService) ListWorkflows(wq *WorkflowQuery) (*Workflows, error) {
	// Build WorkflowQuery into query string params
	v, err := query.Values(wq)
	if err != nil {
		return nil, err
	}

	res, err := c.rizeClient.doRequest(http.MethodGet, "compliance_workflows", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Workflows{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ComplianceWorkflowService) createWorkflow() {}

// ViewCustomerWorkflow to show the latest compliance workflow for a customer
func (c *ComplianceWorkflowService) ViewCustomerWorkflow(customerUID string, cwq *CustomerWorkflowQuery) (map[string]interface{}, error) {
	if customerUID == "" {
		return nil, fmt.Errorf("customerUID is required")
	}

	// Build query params
	v, err := query.Values(cwq)
	if err != nil {
		return nil, err
	}
	res, err := c.rizeClient.doRequest(http.MethodGet, fmt.Sprintf("compliance_workflows/latest/%s", customerUID), v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := make(map[string]interface{})
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ComplianceWorkflowService) acknowledgeDocument()  {}
func (c *ComplianceWorkflowService) acknowledgeDocuments() {}
