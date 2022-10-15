package rize

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// Handles all Compliance Workflow operations
type complianceWorkflowService service

// Workflow data type
type Workflow struct {
	UID                         string                      `json:"uid,omitempty"`
	Summary                     *WorkflowSummary            `json:"summary,omitempty"`
	Customer                    *WorkflowCustomer           `json:"customer,omitempty"`
	ProductUID                  string                      `json:"product_uid,omitempty"`
	ProductCompliancePlanUID    string                      `json:"product_compliance_plan_uid,omitempty"`
	AcceptedDocuments           []*WorkflowAcceptedDocument `json:"accepted_documents,omitempty"`
	CurrentStepDocumentsPending []*WorkflowPendingDocument  `json:"current_step_documents_pending,omitempty"`
	AllDocuments                []*WorkflowDocument         `json:"all_documents,omitempty"`
}

// WorkflowSummary contains a status summary of the Compliance Workflow
type WorkflowSummary struct {
	AcceptedQuantity int       `json:"accepted_quantity,omitempty"`
	BegunAt          time.Time `json:"begun_at,omitempty"`
	CompletedStep    int       `json:"completed_step,omitempty"`
	CurrentStep      int       `json:"current_step,omitempty"`
	Status           string    `json:"status,omitempty"`
}

// WorkflowCustomer contains Customer information related to this Compliance Workflow
type WorkflowCustomer struct {
	Email       string `json:"email,omitempty"`
	ExternalUID string `json:"external_uid,omitempty"`
	UID         string `json:"uid,omitempty"`
}

// WorkflowAcceptedDocument contains information about accepted Compliance Workflow documents
type WorkflowAcceptedDocument struct {
	ElectronicSignatureRequired string    `json:"electronic_signature_required,omitempty"`
	ExternalStorageName         string    `json:"external_storage_name,omitempty"`
	ComplianceDocumentURL       string    `json:"compliance_document_url,omitempty"`
	Name                        string    `json:"name,omitempty"`
	Step                        int       `json:"step,omitempty"`
	Version                     int       `json:"version,omitempty"`
	UID                         string    `json:"uid,omitempty"`
	AcceptedAt                  time.Time `json:"accepted_at,omitempty"`
}

// WorkflowPendingDocument contains information about pending Compliance Workflow documents
type WorkflowPendingDocument struct {
	ElectronicSignatureRequired string `json:"electronic_signature_required,omitempty"`
	ExternalStorageName         string `json:"external_storage_name,omitempty"`
	ComplianceDocumentURL       string `json:"compliance_document_url,omitempty"`
	Name                        string `json:"name,omitempty"`
	Step                        int    `json:"step,omitempty"`
	Version                     int    `json:"version,omitempty"`
	UID                         string `json:"uid,omitempty"`
}

// WorkflowDocument contains information about all Compliance Workflow documents
type WorkflowDocument struct {
	ElectronicSignatureRequired string `json:"electronic_signature_required,omitempty"`
	ExternalStorageName         string `json:"external_storage_name,omitempty"`
	ComplianceDocumentURL       string `json:"compliance_document_url,omitempty"`
	Name                        string `json:"name,omitempty"`
	Step                        int    `json:"step,omitempty"`
	Version                     int    `json:"version,omitempty"`
}

// WorkflowListParams builds the query parameters used in querying Compliance Workflows
type WorkflowListParams struct {
	CustomerUID string `url:"customer_uid,omitempty" json:"customer_uid,omitempty"`
	ProductUID  string `url:"product_uid,omitempty" json:"product_uid,omitempty"`
	InProgress  bool   `url:"in_progress,omitempty" json:"in_progress,omitempty"`
	Limit       int    `url:"limit,omitempty" json:"limit,omitempty"`
	Offset      int    `url:"offset,omitempty" json:"offset,omitempty"`
}

// WorkflowLatestParams builds the query parameters used in querying the latest Compliance Workflow for a customer
type WorkflowLatestParams struct {
	ProductCompliancePlanUID string `url:"product_compliance_plan_uid,omitempty" json:"product_compliance_plan_uid,omitempty"`
}

// WorkflowCreateParams are the body params used when creating a new Compliance Workflow
type WorkflowCreateParams struct {
	CustomerUID              string `json:"customer_uid"`
	ProductCompliancePlanUID string `json:"product_compliance_plan_uid"`
}

// WorkflowDocumentParams are the body params used when acknowledging a compliance document
type WorkflowDocumentParams struct {
	Accept      string `json:"accept"`
	DocumentUID string `json:"document_uid"`
	IPAddress   string `json:"ip_address,omitempty"`
	UserName    string `json:"user_name,omitempty"`
	// Required for AcknowledgeDocument but omitted for AcknowledgeDocuments
	CustomerUID string `json:"customer_uid,omitempty"`
}

// WorkflowDocumentsParams are the body params used when acknowledging multiple compliance documents
type WorkflowDocumentsParams struct {
	CustomerUID string                    `json:"customer_uid"`
	Documents   []*WorkflowDocumentParams `json:"documents"`
}

// WorkflowResponse is an API response containing a list of Compliance Workflows
type WorkflowResponse struct {
	BaseResponse
	Data []*Workflow `json:"data"`
}

// Retrieves a list of Compliance Workflows filtered by the given parameters
func (c *complianceWorkflowService) List(ctx context.Context, wlp *WorkflowListParams) (*WorkflowResponse, error) {
	// Build WorkflowListParams into query string params
	v, err := query.Values(wlp)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodGet, "compliance_workflows", v, nil)
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

// Associates a new Compliance Workflow and set of Compliance Documents (for acknowledgment) with a Customer
func (c *complianceWorkflowService) Create(ctx context.Context, wcp *WorkflowCreateParams) (*Workflow, error) {
	if wcp.CustomerUID == "" || wcp.ProductCompliancePlanUID == "" {
		return nil, fmt.Errorf("CustomerUID and ProductCompliancePlanUID values are required")
	}

	bytesMessage, err := json.Marshal(wcp)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodPost, "compliance_workflows", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Workflow{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// ViewLatest is a helper endpoint for retrieving the most recent Compliance Workflow for a Customer.
// A Customer UID must be supplied as the path parameter.
func (c *complianceWorkflowService) ViewLatest(ctx context.Context, customerUID string, wlp *WorkflowLatestParams) (*Workflow, error) {
	if customerUID == "" {
		return nil, fmt.Errorf("customerUID is required")
	}

	// Build query params
	v, err := query.Values(wlp)
	if err != nil {
		return nil, err
	}
	res, err := c.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("compliance_workflows/latest/%s", customerUID), v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Workflow{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// AcknowledgeDocument is used to indicate acceptance or rejection of a Compliance Document within a given Compliance Workflow
func (c *complianceWorkflowService) AcknowledgeDocument(ctx context.Context, uid string, wd *WorkflowDocumentParams) (*Workflow, error) {
	if uid == "" || wd.Accept == "" || wd.DocumentUID == "" || wd.CustomerUID == "" {
		return nil, fmt.Errorf("UID, Accept, DocumentUID and CustomerUID values are required")
	}

	bytesMessage, err := json.Marshal(wd)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("compliance_workflows/%s/acknowledge_document", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Workflow{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// AcknowledgeDocuments is used to indicate acceptance or rejection of multiple Compliance Documents within a given Compliance Workflow
func (c *complianceWorkflowService) AcknowledgeDocuments(ctx context.Context, uid string, wdp *WorkflowDocumentsParams) (*Workflow, error) {
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

	res, err := c.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("compliance_workflows/%s/batch_acknowledge_documents", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Workflow{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
