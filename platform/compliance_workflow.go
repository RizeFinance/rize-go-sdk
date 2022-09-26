package platform

// ComplianceWorkflowService handles all Compliance Workflow
type ComplianceWorkflowService service

// ListWorkflows retrieves a list of Compliance Workflows filtered by the given parameters
func (c *ComplianceWorkflowService) ListWorkflows()        {}
func (c *ComplianceWorkflowService) createWorkflow()       {}
func (c *ComplianceWorkflowService) viewCustomerWorkflow() {}
func (c *ComplianceWorkflowService) acknowledgeDocument()  {}
func (c *ComplianceWorkflowService) acknowledgeDocuments() {}
