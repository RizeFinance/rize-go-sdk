package rize

import (
	"context"
	"encoding/base64"
	"net/http"
	"testing"
	"time"
)

// Complete KYCDocument{} response data
var kycDocument = &KYCDocument{
	UID:       "u8EHFJnWvJxRJZxa",
	Type:      "contract",
	Filename:  "john_smith_passport",
	Note:      "Uploaded via Client App.",
	Extension: "png",
	CreatedAt: time.Now(),
}

func TestListKYCDocuments(t *testing.T) {
	params := &KYCDocumentListParams{
		EvaluationUID: "QSskNJkryskRXeYt",
	}
	resp, err := rc.KYCDocuments.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching documents\n", err)
	}

	if err := validateSchema(http.MethodGet, "/kyc_documents", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestUploadKYCDocument(t *testing.T) {
	base64Encoding := base64.StdEncoding.EncodeToString([]byte("File info"))
	params := &KYCDocumentUploadParams{
		EvaluationUID: "sdfHFJnWvJxRJZxq",
		Filename:      "fred_smith_license.png",
		FileContent:   base64Encoding,
		Note:          "Uploaded via SDK",
		Type:          "license",
	}
	resp, err := rc.KYCDocuments.Upload(context.Background(), params)
	if err != nil {
		t.Fatal("Error uploading document\n", err)
	}

	if err := validateSchema(http.MethodPost, "/kyc_documents", http.StatusCreated, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetKYCDocument(t *testing.T) {
	resp, err := rc.KYCDocuments.Get(context.Background(), "u8EHFJnWvJxRJZxa")
	if err != nil {
		t.Fatal("Error fetching document\n", err)
	}

	if err := validateSchema(http.MethodGet, "/kyc_documents/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestViewKYCDocument(t *testing.T) {
	_, err := rc.KYCDocuments.View(context.Background(), "u8EHFJnWvJxRJZxa")
	if err != nil {
		t.Fatal("Error viewing document\n", err)
	}

	if err := validateSchema(http.MethodGet, "/kyc_documents/{uid}/view", http.StatusOK, nil, nil, nil); err != nil {
		t.Fatalf(err.Error())
	}
}
