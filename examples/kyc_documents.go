package examples

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"

	"github.com/rizefinance/rize-go-sdk"
)

// List documents
func ExampleKYCDocumentService_List(rc *rize.Client) {
	params := &rize.KYCDocumentListParams{
		EvaluationUID: "QSskNJkryskRXeYt",
	}
	resp, err := rc.KYCDocuments.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching documents\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Documents:", string(output))
}

// Upload Document
func ExampleKYCDocumentService_Upload(rc *rize.Client) {
	bytes, err := os.ReadFile("./file.png")
	if err != nil {
		log.Fatal("Error reading file\n", err)
	}
	base64Encoding := base64.StdEncoding.EncodeToString(bytes)
	params := &rize.KYCDocumentUploadParams{
		EvaluationUID: "sdfHFJnWvJxRJZxq",
		Filename:      "fred_smith_license.png",
		FileContent:   base64Encoding,
		Note:          "Uploaded via SDK",
		Type:          "license",
	}
	resp, err := rc.KYCDocuments.Upload(context.Background(), params)
	if err != nil {
		log.Fatal("Error uploading document\n", err)
	}

	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Upload Document:", string(output))
}

// Get Document
func ExampleKYCDocumentService_Get(rc *rize.Client) {
	resp, err := rc.KYCDocuments.Get(context.Background(), "u8EHFJnWvJxRJZxa")
	if err != nil {
		log.Fatal("Error fetching document\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Document:", string(output))
}

// View Document
func ExampleKYCDocumentService_View(rc *rize.Client) {
	resp, err := rc.KYCDocuments.View(context.Background(), "u8EHFJnWvJxRJZxa")
	if err != nil {
		log.Fatal("Error viewing document\n", err)
	}
	log.Println("View Document:", resp.Status)
}
