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

// Handles all KYC Document operations
type kycDocumentService service

// KYCDocument data type
type KYCDocument struct {
	UID       string    `json:"uid,omitempty"`
	Type      string    `json:"type,omitempty"`
	Filename  string    `json:"filename,omitempty"`
	Note      string    `json:"note,omitempty"`
	Extension string    `json:"extension,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// KYCDocumentListParams builds the query parameters used in querying KYCDocuments
type KYCDocumentListParams struct {
	EvaluationUID string `url:"evaluation_uid,omitempty" json:"evaluation_uid,omitempty"`
}

// KYCDocumentUploadParams are the body params used when uploading a new KYC Document
type KYCDocumentUploadParams struct {
	EvaluationUID string `json:"evaluation_uid"`
	Filename      string `json:"filename"`
	FileContent   string `json:"file_content"`
	Note          string `json:"note"`
	Type          string `json:"type"`
}

// List retrieves a list of KYC Documents for a given evaluation
func (k *kycDocumentService) List(ctx context.Context, params *KYCDocumentListParams) (*ListResponse, error) {
	if params.EvaluationUID == "" {
		return nil, fmt.Errorf("EvaluationUID is required")
	}

	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	res, err := k.client.doRequest(ctx, http.MethodGet, "kyc_documents", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &ListResponse{Data: []*KYCDocument{}}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Upload a KYC Document for review
func (k *kycDocumentService) Upload(ctx context.Context, params *KYCDocumentUploadParams) (*KYCDocument, error) {
	if params.EvaluationUID == "" ||
		params.Filename == "" ||
		params.FileContent == "" ||
		params.Note == "" ||
		params.Type == "" {
		return nil, fmt.Errorf("all KYCDocumentUploadParams are required")
	}

	bytesMessage, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := k.client.doRequest(ctx, http.MethodPost, "kyc_documents", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &KYCDocument{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get is used to retrieve metadata for a KYC Document previously uploaded
func (k *kycDocumentService) Get(ctx context.Context, uid string) (*KYCDocument, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := k.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("kyc_documents/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &KYCDocument{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// View is used to retrieve a KYC Document (image, PDF, etc) previously uploaded
func (k *kycDocumentService) View(ctx context.Context, uid string) (*http.Response, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	// TODO: Does this require a different Accept header type (image/png)?
	res, err := k.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("kyc_documents/%s/view", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}
