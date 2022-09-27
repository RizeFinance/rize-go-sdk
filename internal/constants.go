package internal

import "time"

// Shared SDK Constants
const (
	APITokenMaxAge    = int64((time.Hour * 23) / time.Millisecond)
	APIBasePath       = "api/v1"
	APITimeoutSeconds = time.Second * 30
	SDKVersion        = "1.0.0"
)

// Environments are Rize infrastructure tiers
var Environments = []string{"sandbox", "staging", "integration", "production"}
