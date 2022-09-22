package rize

// Config handles Rize configuration values
type Config struct {
	Environment string
	HMACKey     string
	ProgramUID  string
}

type environment int

const (
	sandbox environment = iota
	integration
	production
)

const DEFAULT_TOKEN_MAX_AGE = 82800000
const DEFAULT_HOST = {
    staging: 'https://staging.rizefs.com',
    release: 'https://release.rizefs.com',
    sandbox: 'https://sandbox.rizefs.com',
    integration: 'https://integration.rizefs.com',
    production: 'https://production.rizefs.com'
};
const DEFAULT_BASE_PATH = '/api/v1/';
const DEFAULT_TIMEOUT = 80000;
const SDK_VERSION = 1.0

func (e environment) String() string {
	envs := [...]string{"Sandbox", "Integration", "Production"}

	if e < sandbox || e > production {
		return "Invalid"
	}

	return envs[e]
}
