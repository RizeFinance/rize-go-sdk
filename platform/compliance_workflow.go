package platform

import (
	"encoding/json"
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

func (c *ComplianceWorkflowService) createWorkflow()       {}
func (c *ComplianceWorkflowService) viewCustomerWorkflow() {}
func (c *ComplianceWorkflowService) acknowledgeDocument()  {}
func (c *ComplianceWorkflowService) acknowledgeDocuments() {}
