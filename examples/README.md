<p align="center">
  <a href="https://developer.rizefs.com/" target="_blank" align="center">
    <img src="https://cdn.rizefs.com/web-content/logos/rize-github.png" width="200">
  </a>
  <br />
</p>

# SDK Examples

Implementation examples for each SDk service.

```sh
# Run an example Platform API method <SERVICE_NAME> <METHOD_NAME>
$ go run cmd/platform/main.go -s CustomerService -m List
```

# Example Customer Onboarding Flow

The steps below will walk through the process of onboarding a new customer. Be sure to check out the [Onboarding Documentation](https://developer.rizefs.com/docs/onboarding-customers-into-a-product) in our [Resource Center](https://developer.rizefs.com/docs) before getting started. 

### 1. Create a New Customer

The first call to start the onboarding flow will be `Customers.Create`. Make a request to create a Customer with an ExternalUID, Email and CustomerType. A new `Customer{}` will be returned.

```go
params := &rize.CustomerCreateParams{
	ExternalUID:  "client-generated-id",
	CustomerType: "primary",
	Email:        "olive.oyl@popeyes.com",
}

resp, err := rc.Customers.Create(context.Background(), params)
if err != nil {
	log.Fatal("Error creating customer\n", err)
}

customer := resp // Customer data
```

### 2. Get Available Products

Next, we need to retrieve the list of Products available for your Program, along with all of the requirements needed for the new Customer to complete. Later on, we'll use the `Product.ProfileRequirements` to submit Profile Responses for all required questions.

```go
params := &rize.ProductListParams{
	ProgramUID: customer.UID,
}
resp, err := rc.Products.List(context.Background(), params)
if err != nil {
	log.Fatal("Error fetching Products\n", err)
}

product := resp.Data[0] // First Product
```

### 3. Create a new Compliance Workflow

Creating a Compliance Workflow requires the new Customer UID and the ProductCompliancePlanUID from the last two steps.

```go
params := &rize.WorkflowCreateParams{
	CustomerUID:              customer.UID,
	ProductCompliancePlanUID: product.ProductCompliancePlanUID,
}
resp, err := rc.ComplianceWorkflows.Create(context.Background(), params)
if err != nil {
	log.Fatal("Error creating Compliance Workflow\n", err)
}

complianceWorkflowUID := resp.UID // Workflow UID
```

Sending a request to the Compliance Workflow initiates a Compliance Workflow for the onboarding Customer. The Compliance Workflow contains disclosure and Compliance Documents that the Customer must acknowledge before opening a Custodial Account.

### 4. Display Compliance Documents

Disclosure and Compliance Documents are defined by the Custodian(s) participating in your Program. The disclosure documents and acknowledgments keep your Program and the Custodian(s) operating in a manner compliant with the Custodian's regulations.

Rize stores all acknowledgements, documents, document versions, and timestamps to assist Custodians with regulatory inquiries. These records are available to you and your customers through the Compliance Workflows endpoint.

```go
params := &rize.WorkflowLatestParams{
	ProductCompliancePlanUID: product.ProductCompliancePlanUID,
}
resp, err := rc.ComplianceWorkflows.ViewLatest(context.Background(), customer.UID, params)
if err != nil {
	log.Fatal("Error fetching Compliance Workflow\n", err)
}
documents := resp.CurrentStepDocumentsPending
```

Your application can display or allow users to download the Compliance Document using the `ComplianceDocumentURL` value returned with each document in the Compliance Workflows response.

### 5. Acknowledge the Compliance Documents

Rize supplies Compliance Documents in steps. All Documents in the current step must be acknowledged before the Documents in the next step are supplied. Any Documents that are awaiting acknowledgement in the current step are available in the `CurrentStepDocumentsPending` array. Once Rize receives the acknowledgements for all Documents in the current step, the documents in the next step will be supplied.

The requirements for acknowledging a Document will differ depending on the Document content. If a Document requires an electronic signature, additional fields must be supplied in the acknowledgment. Your application can discern which Documents require an electronic signature through the `ElectronicSignatureRequired` field returned with each Document in the Compliance Workflows response.

```go
params := &rize.WorkflowBatchDocumentsParams{
	CustomerUID: customer.UID,
	Documents: []*rize.WorkflowDocumentParams{{
		Accept:      "yes",
		DocumentUID: documents[0].UID,
		IPAddress:   "107.56.230.156",
		UserName:    "gilbert chesterton",
	}, {
		Accept:      "yes",
		DocumentUID: documents[1].UID,
	}},
}
resp, err := rc.ComplianceWorkflows.BatchAcknowledgeDocuments(context.Background(), complianceWorkflowUID, params)
if err != nil {
	log.Fatal("Error acknowledging compliance document\n", err)
}
```

### 6. Submit Customer Personally Identifiable Information (PII)

After a Customer starts a Compliance Workflow, Rize expects you to submit their remaining PII.

```go
params := &rize.CustomerUpdateParams{
	Email: "olive.oyl@rizemoney.com",
	Details: &rize.CustomerDetails{
		FirstName:    "Olive",
		MiddleName:   "Olivia",
		LastName:     "Oyl",
		Suffix:       "Jr.",
		Phone:        "5555551212",
		DOB:          internal.DOB(time.Now()),
		SSN:          "111-22-3333",
		Address: &rize.CustomerAddress{
			Street1:    "123 Abc St.",
			Street2:    "Apt 2",
			City:       "Chicago",
			State:      "IL",
			PostalCode: "12345",
		},
	},
}
resp, err := rc.Customers.Update(context.Background(), customer.UID, params)
if err != nil {
	log.Fatal("Error updating customer\n", err)
}
```

### 7. Answer Profile Answers for the Plan Requirements (if needed)

To update profile answers:

```go
params := &rize.CustomerProfileResponseParams{
	ProfileRequirementUID: customer.ProfileResponses[0].ProfileRequirementUID,
	ProfileResponse: &internal.CustomerProfileResponseItem{
		Response: "yes",
	},
}
resp, err := rc.Customers.UpdateProfileResponses(context.Background(), customer.UID, []*rize.CustomerProfileResponseParams{params})
if err != nil {
	log.Fatal("Error updating profile response\n", err)
}
```

### 8. Create new Customer Product

The last step will trigger verification that all Product requirements have been met. After completion, Rize will automatically initiate the KYC process for the Customer's first Product. Upon successful creation, all required Custodial Accounts will be created and the customer will be active. This is a billable event and is intentionally isolated for you to confirm that the Customer record is complete.

To request for identity verification:

```go
params := &rize.CustomerProductCreateParams{
	CustomerUID: customer.UID,
	ProductUID:  productUID,
}
resp, err := rc.CustomerProducts.Create(context.Background(), params)
if err != nil {
	log.Fatal("Error creating Customer Product\n", err)
}
```

## Sandbox Onboarding Evaluations

Once Rize receives complete PII, complete Compliance Document acknowledgments, and the new Customer Product request within the duration requirements, your Customer record will automatically be sent to Rize's identity verification partner.

Because your Customers have assets held by financial institutions, Customers must undergo a KYC/AML evaluation before they can onboard to the Rize Platform.

In the Rize Sandbox, you can generate and test your receipt of the customer evaluation statuses using the last names supplied with the Customer record. The statuses in the Rize Sandbox and the last names that trigger these statuses are listed below.

| Customer Last Name                     | Customer KYC_Status Returned | Description                                                  |
| -------------------------------------- | ---------------------------- | ------------------------------------------------------------ |
| Any value, excluding Smith and Johnson | Approved                     | Approved status indicates that the Custodian will allow the Customer to access their Service Offering immediately. The Customer has completed onboarding to the Program. No documentation is required. |
| Johnson                                | Denied                       | The Customer is denied access to the Custodian's products. No recourse is available through the Custodian or Rize. |
| Smith                                  | Manual Review                | The Custodian is unable to onboard the Customer through an automated process, and additional information is required. |
