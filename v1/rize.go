package rize

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

func (e environment) String() string {
	envs := [...]string{"Sandbox", "Integration", "Production"}

	if e < sandbox || e > production {
		return "Invalid"
	}

	return envs[e]
}
