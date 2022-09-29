package platform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Handles all Compliance Workflow operations
type complianceWorkflowService service

// WorkflowListParams builds the query parameters used in querying compliance workflows
type WorkflowListParams struct {
	CustomerUID string `url:"customer_uid,omitempty"`
	ProductUID  string `url:"product_uid,omitempty"`
	InProgress  bool   `url:"in_progress,omitempty"`
	Limit       int    `url:"limit,omitempty"`
	Offset      int    `url:"offset,omitempty"`
}

// WorkflowLatestParams builds the query parameters used in querying the latest workflow for a customer
type WorkflowLatestParams struct {
	ProductCompliancePlanUID string `url:"product_compliance_plan_uid,omitempty"`
}

// WorkflowCreateParams are the body params used when creating a new compliance workflow
type WorkflowCreateParams struct {
	CustomerUID              string `json:"customer_uid"`
	ProductCompliancePlanUID string `json:"product_compliance_plan_uid"`
}

// WorkflowDocument are the body params used when acknowledging a compliance document
type WorkflowDocument struct {
	Accept      string `json:"accept"`
	DocumentUID string `json:"document_uid"`
	IPAddress   string `json:"ip_address,omitempty"`
	UserName    string `json:"user_name,omitempty"`
	// Required for AcknowledgeDocument but omitted for AcknowledgeDocuments
	CustomerUID string `json:"customer_uid,omitempty"`
}

// WorkflowDocumentsParams are the body params used when acknowledging multiple compliance documents
type WorkflowDocumentsParams struct {
	CustomerUID string             `json:"customer_uid"`
	Documents   []WorkflowDocument `json:"documents"`
}

// WorkflowResponse is an API response containing a list of compliance workflows
type WorkflowResponse struct {
	BaseResponse
	Data []interface{} `json:"data"`
}

// ListWorkflows retrieves a list of Compliance Workflows filtered by the given parameters
func (c *complianceWorkflowService) List(wlp *WorkflowListParams) (*WorkflowResponse, error) {
	// Build WorkflowListParams into query string params
	v, err := query.Values(wlp)
	if err != nil {
		return nil, err
	}

	res, err := c.rizeClient.doRequest(http.MethodGet, "compliance_workflows", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &WorkflowResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateWorkflow associates a new Compliance Workflow and set of Compliance Documents (for acknowledgment) with a Customer
func (c *complianceWorkflowService) Create(wcp *WorkflowCreateParams) (*http.Response, error) {
	if wcp.CustomerUID == "" || wcp.ProductCompliancePlanUID == "" {
		return nil, fmt.Errorf("CustomerUID and ProductCompliancePlanUID values are required")
	}

	bytesMessage, err := json.Marshal(wcp)
	if err != nil {
		return nil, err
	}

	res, err := c.rizeClient.doRequest(http.MethodPost, "compliance_workflows", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil

}

// ViewLatest is a helper endpoint for retrieving the most recent Compliance Workflow for a Customer.
// A Customer UID must be supplied as the path parameter.
func (c *complianceWorkflowService) ViewLatest(customerUID string, wlp *WorkflowLatestParams) (*http.Response, error) {
	if customerUID == "" {
		return nil, fmt.Errorf("customerUID is required")
	}

	// Build query params
	v, err := query.Values(wlp)
	if err != nil {
		return nil, err
	}
	res, err := c.rizeClient.doRequest(http.MethodGet, fmt.Sprintf("compliance_workflows/latest/%s", customerUID), v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}

// AcknowledgeDocument is used to indicate acceptance or rejection of a Compliance Document within a given Compliance Workflow
func (c *complianceWorkflowService) AcknowledgeDocument(uid string, wd *WorkflowDocument) (*http.Response, error) {
	if uid == "" || wd.Accept == "" || wd.DocumentUID == "" || wd.CustomerUID == "" {
		return nil, fmt.Errorf("UID, Accept, DocumentUID and CustomerUID values are required")
	}

	bytesMessage, err := json.Marshal(wd)
	if err != nil {
		return nil, err
	}

	res, err := c.rizeClient.doRequest(http.MethodPut, fmt.Sprintf("compliance_workflows/%s/acknowledge_document", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}

// AcknowledgeDocuments is used to indicate acceptance or rejection of multiple Compliance Documents within a given Compliance Workflow
func (c *complianceWorkflowService) AcknowledgeDocuments(uid string, wdp *WorkflowDocumentsParams) (*http.Response, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	for _, d := range wdp.Documents {
		if d.Accept == "" || d.DocumentUID == "" {
			return nil, fmt.Errorf("Accept and DocumentUID values are required")
		}
		if d.CustomerUID != "" {
			d.CustomerUID = "" // Clear unsupported property
		}
	}

	bytesMessage, err := json.Marshal(wdp)
	if err != nil {
		return nil, err
	}

	res, err := c.rizeClient.doRequest(http.MethodPut, fmt.Sprintf("compliance_workflows/%s/batch_acknowledge_documents", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}
