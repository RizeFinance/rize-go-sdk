package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
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

	// List documents
	kl, err := rc.KYCDocuments.List(context.Background(), "QSskNJkryskRXeYt")
	if err != nil {
		log.Fatal("Error fetching documents\n", err)
	}
	output, _ := json.MarshalIndent(kl, "", "\t")
	log.Println("List Documents:", string(output))

	// Upload Document
	bytes, err := ioutil.ReadFile("./file.png")
	if err != nil {
		log.Fatal("Error reading file\n", err)
	}
	base64Encoding := base64.StdEncoding.EncodeToString(bytes)
	kup := rize.KYCDocumentUploadParams{
		EvaluationUID: "sdfHFJnWvJxRJZxq",
		Filename:      "fred_smith_license.png",
		FileContent:   base64Encoding,
		Note:          "Uploaded via SDK",
		Type:          "license",
	}
	ku, err := rc.KYCDocuments.Upload(context.Background(), &kup)
	if err != nil {
		log.Fatal("Error uploading document\n", err)
	}
	output, _ = json.MarshalIndent(ku, "", "\t")
	log.Println("Upload Document:", string(output))

	// Get Document
	kg, err := rc.KYCDocuments.Get(context.Background(), "u8EHFJnWvJxRJZxa")
	if err != nil {
		log.Fatal("Error fetching document\n", err)
	}
	output, _ = json.MarshalIndent(kg, "", "\t")
	log.Println("Get Document:", string(output))

	// View Document
	kv, err := rc.KYCDocuments.View(context.Background(), "u8EHFJnWvJxRJZxa")
	if err != nil {
		log.Fatal("Error viewing document\n", err)
	}
	output, _ = json.MarshalIndent(kv, "", "\t")
	log.Println("View Document:", string(output))
}
