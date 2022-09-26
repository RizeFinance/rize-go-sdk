package internal

// Shared SDK Constants
const (
	DEFAULT_TOKEN_MAX_AGE = 82800000
	DEFAULT_BASE_PATH     = "/api/v1/"
	DEFAULT_TIMEOUT       = 80000
	SDK_VERSION           = "1.0.0"
)

// ENVIRONMENTS are Rize infrastructure tiers
var ENVIRONMENTS = []string{"sandbox", "integration", "production"}
